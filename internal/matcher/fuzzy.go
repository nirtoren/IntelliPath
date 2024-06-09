package matcher

import (
	"errors"
	"intellipath/internal/record"
	
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type MinFunc func([]PathDist) []PathDist

type PathDist struct {
	Path        string
	LevDistance int
}

func FuzzyFind(path string, records []*record.PathRecord) ([]PathDist, error) {

	convertor := record.NewPathRecConvertor()
	dbPaths := convertor.RecordsToPaths(records)
	
	pathMap := make(map[string]string)
	formatter := record.NewPathFormatter()
	for _, fullPath := range dbPaths {
		base := formatter.ToBase(fullPath)
		pathMap[base] = fullPath
	}

	basePaths := make([]string, 0, len(pathMap))
	for k := range pathMap {
		basePaths = append(basePaths, k)
	}

	var foundPaths []string = fuzzy.Find(path, basePaths) //Should get a list of optional paths

	if len(foundPaths) == 0 {
		// Should exit the program with a message to the user
		return []PathDist{}, errors.New("could not find suitable path") // In case could not find any match

	} else if len(foundPaths) == 1 {
		singlePath := PathDist{Path: pathMap[foundPaths[0]], LevDistance: 0}
		return []PathDist{singlePath}, nil // In case only one match was found, no need to procced to Levenshtein OR Score filters

	} else { // In case more then one match was found
		levenshteinPaths := FindLevenshteinDistance(path, foundPaths, findTwoMin, pathMap) // Should get a list of two elements for future score filtering
		return levenshteinPaths, nil
	}
}
