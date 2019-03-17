package service

import (
	"../domain"
	"math/rand"
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

	for i := 0; i < configs.NumberAmount; i++ {
		result = append(result, getRandomValue(numbers))
	}

	for i := 0; i < configs.SpecialCharsAmount; i++ {
		result = append(result, getRandomValue(specialChars))
	}

	for i := 0; i < configs.MinLength+rand.Intn(20); i++ {
		result = append(result, getRandomValue(letters))
	}

	for i := 0; i < len(result); i++ {
		if _, ok := vowels[result[i]]; ok {
			result[i] = getRandomValue(numbers)
		}
	}

	return string(shuffleSlice(result))
}

func getRandomValue(runeSlice []rune) rune {
	return runeSlice[rand.Intn(len(runeSlice))]
}

func shuffleSlice(runeSlice []rune) []rune {
	res := make([]rune, len(runeSlice))
	perm := rand.Perm(len(runeSlice))
	for i, randIndex := range perm {
		res[i] = runeSlice[randIndex]
	}
	return res
}
