package interfaces

import (
	"fmt"
)

type Validate interface{
	ValidateInput(string) error
	ValidateFlags(string) error
}

type validator struct{
	Validate
	userInput []string
}

func NewValidator(userInput []string) (*validator, error) {
	if len(userInput) < 1 {
		panic("Incorrect user input")
	}

	return &validator{
		userInput: userInput,
	}, nil
}

func (validator *validator) ValidateInput(userInput string) error {
	fmt.Println(userInput)
	return nil
}

func (validator *validator) ValidateFlags(flag string) error {
	fmt.Println(flag)
	return nil
}
