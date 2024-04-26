package algo

import (
	"errors"
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type MinFunc func([]PathDistRecord) []PathDistRecord

type PathDistRecord struct{
	Path string
	LevDistance int
}

func FuzzyFind(path string, dbPaths []string) ([]PathDistRecord, error) {

	var foundPaths []string = fuzzy.Find(path, dbPaths) //Should get a list of optional paths 
	// var nullPathDistRec []PathDistRecord
	if len(foundPaths) == 0 {

		return []PathDistRecord{}, errors.New("could not find suitable path") // In case could not find any match

	} else if len(foundPaths) == 1{
		singlePath := PathDistRecord{Path: foundPaths[0], LevDistance: 0}
		return []PathDistRecord{singlePath}, nil // In case only one match was found, no need to procced to Levenshtein OR Score filters
	
		} else { // In case more then one match was found
	
			levenshteinPaths := findMinLevenshteinDistance(path, foundPaths, findTwoMin) // Should get a list of two elements for future score filtering
			return levenshteinPaths, nil
	}
}

func findMinLevenshteinDistance(userInput string, foundPaths []string, minFunction MinFunc) []PathDistRecord {
	var pathDist []PathDistRecord

	for _, path := range foundPaths{
		distance := fuzzy.LevenshteinDistance(userInput, path)
		pathDist = append(pathDist, PathDistRecord{Path: path, LevDistance: distance})
	}

	minimalLevenshtainDistance := minFunction(pathDist)

	return minimalLevenshtainDistance

}

func findTwoMin(records []PathDistRecord) []PathDistRecord {
	
	var sortedRecs []PathDistRecord

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