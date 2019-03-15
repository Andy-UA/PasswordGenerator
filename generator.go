package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

var SpecialChars = []rune("[{}\"/$\\_`~&+,:;=?@#|'<>.-^*()%!]")
var Numbers = []rune("1234567890")
var Letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var Vowels = []rune("AEIOUYaeiouy")

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
		http.Error(w, err.Error(), 500)
		return
	}
	if err := json.Unmarshal(body, &passwordConfigs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{
		generatePassword(passwordConfigs),
	}

	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func generatePassword(configs PasswordConfig) string {
	var result = []rune("")

	for i := 0; i < configs.NumberAmount; i++ {
		result = append(result, getRandomValue(Numbers))
	}

	for i := 0; i < configs.SpecialCharsAmount; i++ {
		result = append(result, getRandomValue(SpecialChars))
	}

	loopRange := configs.MinLength + rand.Intn(20)
	for i := 0; i < loopRange; i++ {
		result = append(result, getRandomValue(Letters))
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(Vowels); j++ {
			if result[i] == Vowels[j] {
				result[i] = getRandomValue(Numbers)
			}
		}
	}

	return string(shuffleSlice(result))
}

func getRandomValue(runeSlice []rune) rune{
	return runeSlice[rand.Intn(len(runeSlice))]
}

func shuffleSlice(runeSlice []rune) []rune {
	for i := 0; i < len(runeSlice); i++ {
		tmp := runeSlice[rand.Intn(len(runeSlice))]
		tmp, runeSlice[i] = runeSlice[i], tmp
	}
	return runeSlice
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/generate", getPasswordConfigs).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
