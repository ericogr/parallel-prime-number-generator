package main

import "fmt"

func ValidatePositiveNumber(number int) error {
	if number < 0 {
		return fmt.Errorf("validation: invalid number %d", number)
	}

	return nil
}
