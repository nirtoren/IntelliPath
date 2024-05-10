package flow

import (
	"intellipath/internal/record"
)

type Direct struct {
	pathsdb      *record.Database
	absolutePath string
}

func InitDirectFlow(pathDB *record.Database, absolutePath string) *Direct {
	if pathDB == nil {
		panic("could not initialize Light flow due to DB issue")
	}

	return &Direct{
		pathsdb:      pathDB,
		absolutePath: absolutePath,
	}
}

func (light *Direct) Act() (string, error) { // This should later on return a record
	var outPath string

	rec, err := light.pathsdb.PathSearch(light.absolutePath) // This should return a record if it exists
	if err != nil {
		return "", err
	}

	switch rec.GetPath() {
	case "": // In case no record was found
		record, err := record.NewRecord(light.absolutePath, 0)
		if err != nil {
			return "", err
		}

		if _, err = light.pathsdb.InsertRecord(record); err != nil {
			return "", err
		}
		outPath = light.absolutePath

	case light.absolutePath: // In case a matching record was found
		if err := light.pathsdb.UpdateScore(rec); err != nil {
			return "", err
		}
		outPath = light.absolutePath

	}

	return outPath, nil
}
