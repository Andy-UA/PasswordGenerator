package service

import (
	"PasswordGenerator/domain"
	"math/rand"
	"time"
)

var (
	specialChars = []rune("[{}\"/$\\_`~&+,:;=?@#|'<>.-^*()%!]")
	numbers      = []rune("1234567890")
	letters      = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	vowels       = map[rune]struct{}{'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}, 'Y': {},
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'y': {}}
)

func GeneratePassword(configs domain.PasswordConfig) string {
	var result []rune
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < configs.NumberAmount; i++ {
		result = append(result, numbers[rand.Intn(len(numbers))])
	}

	for i := 0; i < configs.SpecialCharsAmount; i++ {
		result = append(result, specialChars[rand.Intn(len(specialChars))])
	}

	for i := 0; i < configs.MinLength+rand.Intn(20); i++ {
		result = append(result, letters[rand.Intn(len(letters))])
	}

	for i := 0; i < len(result); i++ {
		if _, ok := vowels[result[i]]; ok {
			result[i] = numbers[rand.Intn(len(numbers))]
		}
	}

	return string(shuffleSlice(result))
}

func shuffleSlice(runeSlice []rune) []rune {
	rand.Shuffle(len(runeSlice), func(i, j int) {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	})
	return runeSlice
}
