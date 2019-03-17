package main

import (
	"fmt"
	"os"
	"strconv"

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

	var convertedArgs []int
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

	fmt.Printf("Generated password - %v", service.GeneratePassword(passwordConfigs))
}
