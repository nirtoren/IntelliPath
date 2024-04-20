package flow

import (
	"intellipath/internal/db"
	"intellipath/internal/algorithms"
)

type Light struct{
	pathsdb *db.Database
	score algo.Score
	relativePath string
	absolutePath string
}

func (l *Light) Act() error{

	return nil
}

func (l *Light) isExistInDB() (int8, error){

	return 0, nil
}

