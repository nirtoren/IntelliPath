// Un-used for now 

package utils

import (
	"fmt"
	"os"
	"errors"
	"intellipath/internal/constants"

)

type Validate interface{
	ValidateInput(string) error
	ValidateFlags(string) error
}

type validator struct{
	Validate
}

func NewValidator() *validator {
	return &validator{}
}

func (validator *validator) ValidateInput(userInput string) error {
	fmt.Println(userInput)
	return nil
}

func (validator *validator) ValidateFlags(flag string) error {
	fmt.Println(flag)
	return nil
}

func (validator *validator) ValidateENV() error {
	_, exists := os.LookupEnv(constants.INTELLIPATH_DIR) 
	if !exists{
		return errors.New("_INTELLIPATH_DIR not found within environmental variables")
	}

	_, exists = os.LookupEnv(constants.INTELLIPATH_DB_DTIMER) 
	if !exists{
		return errors.New("_INTELLIPATH_DB_DTIMER not found within environmental variables")
	}

	return nil
}