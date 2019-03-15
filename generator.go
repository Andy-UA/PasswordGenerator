package main

// make packages: api (handler, middleware), domain (structs), service (PasswordGenerator), cmd (cli for access to service)
// structure cmd(server(main.go), cli(main.go), api(), domain(), service())
// optional: Package cobra
// add DockerFile to build image for server and cli
// add MakeFile, which would call docker and run container
// test coverage (unit tests)
// use goModule (create VendorFolder)

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

var SpecialChars = []rune("[{}\"/$\\_`~&+,:;=?@#|'<>.-^*()%!]")
var Numbers = []rune("1234567890")
var Letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var Vowels = []rune("AEIOUYaeiouy") //should be as map of empy structs

type PasswordConfig struct {
	MinLength          int
	SpecialCharsAmount int
	NumberAmount       int
}

type Response struct {
	GeneratedPassword string
}

func getPasswordConfigs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var passwordConfigs PasswordConfig
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() //check on error
	if err := json.Unmarshal(body, &passwordConfigs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := generatePassword(passwordConfigs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseInstance := Response{
		response,
	}
	resp, err := json.Marshal(responseInstance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generatePassword(configs PasswordConfig) (string, error) {
	var result []rune

	if configs.MinLength < 0 || configs.NumberAmount < 0 || configs.SpecialCharsAmount < 0 {
		return "", fmt.Errorf("parameter's value can't be less than 0")
	}

	for i := 0; i < configs.NumberAmount; i++ {
		result = append(result, getRandomValue(Numbers))
	}

	for i := 0; i < configs.SpecialCharsAmount; i++ {
		result = append(result, getRandomValue(SpecialChars))
	}

	for i := 0; i < configs.MinLength + rand.Intn(20); i++ {
		result = append(result, getRandomValue(Letters))
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(Vowels); j++ {
			if result[i] == Vowels[j] {
				result[i] = getRandomValue(Numbers)
			}
		}
	}

	return string(shuffleSlice(result)), nil
}

func getRandomValue(runeSlice []rune) rune{
	return runeSlice[rand.Intn(len(runeSlice))]
}

//rewrite shuffle
func shuffleSlice(runeSlice []rune) []rune {
	//res := make([]rune, len(runeSlice))
	for i := 0; i < len(runeSlice); i++ {
		tmp := runeSlice[rand.Intn(len(runeSlice))]
		//res[i] =
		runeSlice[i], tmp = tmp, runeSlice[i]
	}
	return runeSlice
	//return res
}

func main() {
	router := mux.NewRouter()
	//add middleware to validate request and use context for passing valid data
	router.HandleFunc("/generate", getPasswordConfigs).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
