package algo

import (
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)


func FindLevenshteinDistance(userInput string, foundPaths []string, minFunction MinFunc, pathMap map[string]string) []PathDist {
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