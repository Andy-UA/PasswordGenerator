package main

import (
	"fmt"
	"os"
	"strconv"

	"../../api"
	"../../domain"
	"../../service"
)

func main() {
	argsWithoutName := os.Args[1:]

	if len(argsWithoutName) < 3 || len(argsWithoutName) > 3 {
		fmt.Println("Wrong amount of arguments!\n\tArguments template:\n\t" +
			"[minimal length] [special characters amount] [amount of numbers]")
		return
	}

	convertedArgs := make([]int, 3)
	for _, v := range argsWithoutName {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v is not a number!", v)
			return
		}
		convertedArgs = append(convertedArgs, num)
	}

	passwordConfigs := domain.PasswordConfig{
		MinLength: convertedArgs[0],
		SpecialCharsAmount: convertedArgs[1],
		NumberAmount: convertedArgs[2],
	}

	if err := api.CheckBodyContent(passwordConfigs); err != nil {
		return
	}

	fmt.Printf("Generated password - %v", service.GeneratePassword(passwordConfigs))
}
