package service

import (
	"PasswordGenerator/domain"
	"fmt"
)

func CheckBodyContent(configs domain.PasswordConfig) error {
	if configs.MinLength <= 0 || configs.NumberAmount <= 0 || configs.SpecialCharsAmount <= 0 {
		return fmt.Errorf("parameter's value can't be less or equal to 0")
	}
	return nil
}
