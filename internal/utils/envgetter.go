package utils

import "intellipath/internal/constants"

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
	pathFormatter := NewPathFormatter()
	dir, _ := env.validator.validateIntellipathDirENV()

	dbFile = dir + constants.DBpath

	DBabsolutePath := pathFormatter.ToAbs(dbFile)
	isDBExists := pathFormatter.IsExists(DBabsolutePath)

	if !isDBExists {
		panic("Could not find the database")
	}

	return dbFile
}
