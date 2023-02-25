package main

import "fmt"

func ValidateGreaterThanZero(number int) error {
	if number <= 0 {
		return fmt.Errorf("invalid number %d (must be greater than 0)", number)
	}

	return nil
}
