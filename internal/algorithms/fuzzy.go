package algo

import (
	"errors"
	"sort"
	"intellipath/internal/utils"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type MinFunc func([]PathDist) []PathDist

type PathDist struct{
	Path string
	LevDistance int
}

func FuzzyFind(path string, dbPaths []string) ([]PathDist, error) {

	pathMap := make(map[string]string)
	formatter := utils.NewPathFormatter()
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

	} else if len(foundPaths) == 1{
		singlePath := PathDist{Path: pathMap[foundPaths[0]], LevDistance: 0}
		return []PathDist{singlePath}, nil // In case only one match was found, no need to procced to Levenshtein OR Score filters
	
	} else { // In case more then one match was found
		levenshteinPaths := findMinLevenshteinDistance(path, foundPaths, findTwoMin, pathMap) // Should get a list of two elements for future score filtering
		return levenshteinPaths, nil
	}
}

func findMinLevenshteinDistance(userInput string, foundPaths []string, minFunction MinFunc, pathMap map[string]string) []PathDist {
	var pathDist []PathDist

	for _, path := range foundPaths{
		distance := fuzzy.LevenshteinDistance(userInput, path)
		pathDist = append(pathDist, PathDist{Path: pathMap[path], LevDistance: distance})
	}

	minimalLevenshtainDistance := minFunction(pathDist)

	return minimalLevenshtainDistance

}

func findTwoMin(records []PathDist) []PathDist {
	
	var sortedRecs []PathDist

	sort.Slice(records, func (i, j int) bool {
		return records[i].LevDistance < records[j].LevDistance
	})

	if len(records) > 0 {
		sortedRecs = append(sortedRecs,records[0])
		if len(records) > 1 {
			sortedRecs = append(sortedRecs, records[1])
		}
	}

	return sortedRecs
}
