package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func main() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "peppo" {
		return NotFoundError
	}

	return nil
}
