package env

import (
	"intellipath/internal/constants"
	"os"
	"path/filepath"
)

type ENVGetter struct {
	validator *validator
}

func NewENVGetter(validator *validator) *ENVGetter {
	return &ENVGetter{validator: validator}
}

func (env *ENVGetter) GetIntellipathDir() string {
	dir, err := env.validator.validateIntellipathDirENV()
	if err != nil {
		panic(err)
	}
	return dir
}

func (env *ENVGetter) GetDBPath() string {
	var dbFile string
	dir, _ := env.validator.validateIntellipathDirENV()
	dbFile = dir + constants.DBpath

	DBabsolutePath, _ := filepath.Abs(dbFile)
	_, err := os.Stat(DBabsolutePath)
	isDBExists := !os.IsNotExist(err)
	if !isDBExists {
		panic("Could not find the database")
	}

	return DBabsolutePath
}
