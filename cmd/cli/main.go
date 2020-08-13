package main

import (
	"fmt"
	"os"
	"strconv"

	"PasswordGenerator/domain"
	"PasswordGenerator/service"
)

func main() {
	argsWithoutName := os.Args[1:]

	if len(argsWithoutName) != 3 {
		fmt.Println("Wrong amount of arguments!\n\tArguments template:\n\t" +
			"[minimal length] [special characters amount] [amount of numbers]")
		return
	}

	convertedArgs := make([]int, len(argsWithoutName))
	for i, v := range argsWithoutName {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v is not a number!", v)
			return
		}
		convertedArgs[i] = num
	}

	passwordConfigs := domain.PasswordConfig{
		MinLength:          convertedArgs[0],
		SpecialCharsAmount: convertedArgs[1],
		NumberAmount:       convertedArgs[2],
	}

	if err := service.CheckBodyContent(passwordConfigs); err != nil {
		return
	}

	fmt.Printf("Generated password - %v", service.GeneratePassword(passwordConfigs))
}
