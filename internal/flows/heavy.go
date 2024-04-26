package flow

import (
	"errors"
	"fmt"
	"intellipath/internal/algorithms"
	"intellipath/internal/db"
)

type Heavy struct{
	pathsdb *db.Database
	score algo.Score
	basePath string
}

func InitHeavyFlow(pathDB *db.Database, basePath string) *Heavy{
	if pathDB == nil{
		fmt.Errorf("could not initialize Heavy flow")
		return nil
	}

	return &Heavy{
		pathsdb: pathDB,
		score: algo.Score{},
		basePath: basePath,
	}
}

func (h *Heavy) Act() (string, error){
	// Check in DB + fuzzy + levinshtein
	// if exists -> get result -> try cd -> delete path if fails / Score up & Act.
	// if does not exists -> fail the process As 'cd' would fail
	paths, err := h.pathsdb.GetAllPaths()
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}
	fuzzyResPaths, err := algo.FuzzyFind(h.basePath, paths) //fuzzy + levinshtein getting a PathDistRecord struct
	fmt.Println(fuzzyResPaths, err)
	// records, err := l.pathsdb.GetRecordsByName(fuzzyResPaths)
	// Score filtering
	// for path, score := range records {
	// 	// get the path with the heighest score
	// }

	return "", nil
}
