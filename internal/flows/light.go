package flow

import (
	"intellipath/internal/db"
)

type Light struct{
	pathsdb *db.Database
	absolutePath string
}

func InitLightFlow(pathDB *db.Database, absolutePath string) *Light{
	if pathDB == nil{
		panic("could not initialize Light flow due to DB issue")
	}

	return &Light{
		pathsdb: pathDB,
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
		
		if _ ,err = light.pathsdb.InsertPath(record); err != nil {
			return "", err
		}
		outPath = light.absolutePath

	case light.absolutePath: // In case a matching record was found
		if err := light.pathsdb.UpdateScore(light.absolutePath, score); err != nil {
			return "", err
		}
		outPath = light.absolutePath

	}
	
	return outPath, nil
}
