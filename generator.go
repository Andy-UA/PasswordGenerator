package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"sort"
)

type PasswordConfig struct {
	minLength int `json:"min_length"`
	specialCharsAmount int `json:"special_chars_amount"`
	numberAmount int `json:"number_amount"`
}

type ResponseMessage struct {
	originalPassword string
	transformPassword string
}

func getPasswordConfigs(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var passwordConfigs PasswordConfig
	_ = json.NewDecoder(r.Body).Decode(&passwordConfigs)
	response := generatePassword(passwordConfigs)
	json.NewEncoder(w).Encode(response)
}

func generatePassword(configs PasswordConfig) ResponseMessage {
	var originalPassword string
	numberAmount := configs.numberAmount
	specialCharsAmount :=  configs.specialCharsAmount
	loopRange := configs.minLength + configs.numberAmount + configs.specialCharsAmount + rand.Int()

	for i := 0; i < loopRange; i++ {
		checkNumAmount(&numberAmount, &originalPassword)
		checkSpecialChars(&specialCharsAmount, &originalPassword)
		addLetters(&originalPassword)
	}

	var response ResponseMessage
	response.originalPassword = originalPassword
	vowels := []int{65, 69, 73, 79, 85, 89, 97, 101, 105, 111, 117, 121}
	var transformPassword string

	for _, v := range originalPassword {
		if v == int32(sort.SearchInts(vowels, int(v))) && rand.Intn(2) != 0 {
			transformPassword += fmt.Sprint(getValueFromRange(0, 9))
		} else {
			transformPassword += fmt.Sprintf("%c", v)
		}
	}

	response.transformPassword = transformPassword

	return response
}

func checkNumAmount(numAmount *int, password *string) {
	if *numAmount > 0 && rand.Intn(2) != 0 {
		*password += fmt.Sprintf("%c", getValueFromRange(48, 57))
		*numAmount--
	}
}

func checkSpecialChars(charsAmount *int, password *string)  {
	if *charsAmount > 0 && rand.Intn(2) != 0 {
		switch rand.Intn(4) {
		case 0: *password += fmt.Sprintf("%c", getValueFromRange(33, 47))
		case 1: *password += fmt.Sprintf("%c", getValueFromRange(58, 64))
		case 2: *password += fmt.Sprintf("%c", getValueFromRange(91, 96))
		case 3: *password += fmt.Sprintf("%c", getValueFromRange(123, 126))
		}
		*charsAmount--
	}
}

func addLetters(password *string)  {
	switch rand.Intn(2) {
	case 0: *password += fmt.Sprintf("%c", getValueFromRange(65, 90))
	case 1: *password += fmt.Sprintf("%c", getValueFromRange(97, 122))
	}
}

func getValueFromRange(min int, max int) int {
	return rand.Intn(max - min) + min
}

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/generate", getPasswordConfigs).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
