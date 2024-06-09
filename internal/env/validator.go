// Un-used for now 

package env

import (
	"os"
	"errors"
	"intellipath/internal/constants"

)

type validator struct{}

func NewValidator() *validator {
	return &validator{}
}

func (validator *validator) ValidateENVs() {
	var err error
	_, err = validator.validateIntellipathDirENV()
	if err != nil {
		panic(err)
	}
	_, err = validator.validateIntellipathDTimerENV()
	if err != nil {
		panic(err)
	}
}

func (validator *validator) validateIntellipathDirENV() (string, error) {
	dir := os.Getenv(constants.INTELLIPATH_DIR)
	// if !exists{
	// 	return "", errors.New("_INTELLIPATH_DIR not found within environmental variables")
	// }

	return dir, nil
}

func (validator *validator) validateIntellipathDTimerENV() (string, error) {
	days, exists := os.LookupEnv(constants.INTELLIPATH_DB_DTIMER) 
	if !exists{
		return days, errors.New("_INTELLIPATH_DB_DTIMER not found within environmental variables")
	}

	return days, nil
}
