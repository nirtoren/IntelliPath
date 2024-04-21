package flow

import (
	"database/sql"
	"errors"
	"fmt"
	"intellipath/internal/algorithms"
	"intellipath/internal/db"
	"path/filepath"
	"sort"
)

type Light struct{
	pathsdb *db.Database
	score algo.Score
	relativePath string
	absolutePath string
}

func InitLightFlow(pathDB *db.Database, relativePath string) *Light{
	if pathDB == nil{
		fmt.Errorf("could not initialize Light flow")
		return nil
	}

	absolutePath, err := filepath.Abs(relativePath)
	if err != nil{
		fmt.Errorf("could not convert to absolute path")
		return nil
	}

	return &Light{
		pathsdb: pathDB,
		score: algo.Score{},
		relativePath: relativePath,
		absolutePath: absolutePath,
	}
}

func (l *Light) Act() error{
	// Check absolute path in db
	// if in DB -> Score up & Act
	path, err := l.pathsdb.PathSearch(l.absolutePath)
	if err != nil {
		return errors.New("could not get paths from DB")
	} else if path == "" && err == sql.ErrNoRows { // In case path does not exists in DB
		record, err := db.NewRecord(l.absolutePath, 0)
		if err != nil {
			return err
		}
		_ ,err = l.pathsdb.InsertPath(record)
		if err != nil {
			return err
		}
	} else if path != "" && err == nil{ // In case path DOES exists in DB
		// get the db row
		// update the score
		// return the absolute path as stdout
		
	}
	return nil
}

func (l *Light) isExistInDB() (bool, error){
	
	return false, errors.New("could not get paths from DB")
}

