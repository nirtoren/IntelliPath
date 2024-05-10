package flow

import (
	"errors"
	algo "intellipath/internal/algorithms"
	"intellipath/internal/record"
	"intellipath/internal/utils"

)

type FuzzyFlow struct {
	pathsdb  *record.Database
	basePath string
}

func InitFuzzyFlow(pathDB *record.Database, basePath string) *FuzzyFlow {
	if pathDB == nil {
		panic("could not initialize Heavy flow due to DB issue")
	}

	return &FuzzyFlow{
		pathsdb:  pathDB,
		basePath: basePath,
	}
}

func (h *FuzzyFlow) Act() (string, error) {
	// Check in DB + fuzzy + levinshtein
	pathFormatter := utils.NewPathFormatter()

	paths, err := h.pathsdb.GetAllPaths()
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}

	fuzzyResPaths, err := algo.FuzzyFind(h.basePath, paths) //fuzzy + levinshtein getting a PathDistRecord struct
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}

	var fuzzyPaths []string
	for _, pathRes := range fuzzyResPaths {
		fuzzyPaths = append(fuzzyPaths, pathRes.Path)
	}

	records, err := h.pathsdb.GetRecordsByName(fuzzyPaths)
	if err != nil {
		return "", errors.New("could not get paths from DB")
	}

	// filter by score function
	rec, err := h.filterByScore(records)
	if err != nil {
		return "", errors.New("failed to filter records by score")
	}

	// Check if found path exists
	if !pathFormatter.IsExists(rec.GetPath()) {
		h.pathsdb.DeletePath(rec)
		panic("Path does not exists")
	} else {
		_ = h.pathsdb.UpdateScore(rec)
		return rec.GetPath(), nil
	}
}

func (h *FuzzyFlow) filterByScore(records []*record.PathRecord) (*record.PathRecord, error) {

	if len(records) < 1 {
		return nil, errors.New("could not find any paths")
	} else if len(records) == 1 {
		return records[0], nil
	} else {
		if records[0].GetScore() > records[1].GetScore() {
			return records[0], nil
		} else {
			return records[1], nil
		}
	}
}
