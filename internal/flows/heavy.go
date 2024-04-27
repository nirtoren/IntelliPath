package flow

import (
	"errors"
	"intellipath/internal/algorithms"
	"intellipath/internal/db"
	"intellipath/internal/interfaces"
)

type Heavy struct{
	pathsdb *db.Database
	basePath string
}

func InitHeavyFlow(pathDB *db.Database, basePath string) *Heavy{
	if pathDB == nil{
		panic("could not initialize Heavy flow due to DB issue")
	}

	return &Heavy{
		pathsdb: pathDB,
		basePath: basePath,
	}
}

func (h *Heavy) Act() (string, error){
	// Check in DB + fuzzy + levinshtein
	pathFormatter := interfaces.NewPathFormatter()

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
	outPath, err := h.filterByScore(records)
	if err != nil {
		return "", errors.New("failed to filter records by score")
	}

	// Check if found path exists
	if !pathFormatter.IsExists(outPath) {
		h.pathsdb.DeletePath(outPath)
		panic("Path does not exists")
	}
	return outPath, nil
}

func (h *Heavy) filterByScore(records []db.PathRecord) (string, error) {

	if len(records) < 1 {
		return "", errors.New("could not find any paths")
	} else if len(records) == 1{
		return records[0].Path, nil
	} else {
		if records[0].Score > records[1].Score {
			return records[0].Path, nil
		} else {
			return records[1].Path, nil
		}
	}
}