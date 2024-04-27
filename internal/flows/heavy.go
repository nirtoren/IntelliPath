package flow

import (
	"errors"
	"fmt"
	"intellipath/internal/algorithms"
	"intellipath/internal/db"
)

type Heavy struct{
	pathsdb *db.Database
	basePath string
}

func InitHeavyFlow(pathDB *db.Database, basePath string) *Heavy{
	if pathDB == nil{
		fmt.Errorf("could not initialize Heavy flow")
		return nil
	}

	return &Heavy{
		pathsdb: pathDB,
		basePath: basePath,
	}
}

func (h *Heavy) Act() (string, error){
	// Check in DB + fuzzy + levinshtein
	var outPath string
	var fuzzyPaths []string

	paths, err := h.pathsdb.GetAllPaths()
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}

	fuzzyResPaths, err := algo.FuzzyFind(h.basePath, paths) //fuzzy + levinshtein getting a PathDistRecord struct
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}
	for _, pathRes := range fuzzyResPaths {
		fuzzyPaths = append(fuzzyPaths, pathRes.Path)
	}

	records, err := h.pathsdb.GetRecordsByName(fuzzyPaths)
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}
	
	// filter by score function

	if len(records) < 1 {
		outPath = ""
	} else if len(records) == 1{
		outPath = records[0].Path
	} else {
		if records[0].Score > records[1].Score {
			outPath = records[0].Path
		} else {
			outPath = records[1].Path
		}
	}

	return outPath, nil
}

