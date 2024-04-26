package flow

import (
	"errors"
	"fmt"
	"intellipath/internal/algorithms"
	"intellipath/internal/db"
	"os"
)

type Light struct{
	pathsdb *db.Database
	score algo.Score
	absolutePath string
}

func InitLightFlow(pathDB *db.Database, absolutePath string) *Light{
	if pathDB == nil{
		fmt.Errorf("could not initialize Light flow")
		return nil
	}

	return &Light{
		pathsdb: pathDB,
		score: algo.Score{},
		absolutePath: absolutePath,
	}
}

func (light *Light) Act() (string, error){ // This should later on return a record
	var outPath string

	path, score, err := light.pathsdb.PathSearch(light.absolutePath) // This should return a record if it exists
	if err != nil {
		return "", err
	}
	switch path {
	case "": // In case no record was found
		record, err := db.NewRecord(light.absolutePath, 0)
		if err != nil {
			return "", err
		}

		_ ,err = light.pathsdb.InsertPath(record)
		if err == nil {
			outPath = light.absolutePath
		}
	case light.absolutePath: // In case a matching record was found
		err = light.pathsdb.UpdateScore(light.absolutePath, score) // This should later on return an updated record 
		if err == nil {
			outPath = light.absolutePath
		}
	}
	
	return outPath, nil
}
