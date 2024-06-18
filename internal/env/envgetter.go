package env

import (
	"intellipath/internal/constants"
	"path/filepath"
	// "os"
	// "path/filepath"
)

type ENVGetter struct {
	validator *validator
}

func NewENVGetter(validator *validator) *ENVGetter {
	return &ENVGetter{validator: validator}
}

func (env *ENVGetter) GetIntellipathDir() (string, error) {
	dir, err := env.validator.validateIntellipathDirENV()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func (env *ENVGetter) GetDBPath() (string, error) {
	dir, err := env.validator.validateIntellipathDirENV()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, constants.DBLOCATION), nil
}
