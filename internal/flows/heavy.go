package flow

import (
	"intellipath/internal/db"
	"intellipath/internal/algorithms"
)

type Heavy struct{
	pathsdb *db.Database
	score algo.Score
	relativePath string
	absolutePath string
}

func (l *Heavy) Act() error{

	return nil
}

func (l *Heavy) isExistInDB() (int8, error){

	return 0, nil
}

