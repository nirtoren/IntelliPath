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

func (l *Light) Act() (string, error){
	path, score, err := l.pathsdb.PathSearch(l.absolutePath)
	if err != nil {
		return "", errors.New("could not get paths from DB")
	} else if path == "" && err == nil { // In case path does not exists in DB
		record, err := db.NewRecord(l.absolutePath, 0)
		if err != nil {
			return "", err
		}
		_ ,err = l.pathsdb.InsertPath(record)
		if err != nil {
			return "", err
		} else {
			os.Stdout.WriteString(l.absolutePath)
		}
	} else{ // In case path DOES exists in DB
		l.pathsdb.UpdateScore(l.absolutePath, score)
		os.Stdout.WriteString(l.absolutePath)
	}
	return l.absolutePath, nil
}
