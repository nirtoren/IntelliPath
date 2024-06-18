package matcher

import (
	"errors"
	"intellipath/internal/record"
	"sort"

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

	// var foundPaths []string = fuzzy.Find(path, basePaths) //Should get a list of optional paths

	var rankedPaths fuzzy.Ranks = fuzzy.RankFind(path, basePaths)

	// fmt.Println(rankedPaths)

	if rankedPaths.Len() == 0 {
		return []PathDist{}, errors.New("could not find suitable path") // In case could not find any match
	} else if rankedPaths.Len() == 1 {
		singlePath := PathDist{Path: pathMap[rankedPaths[0].Target], LevDistance: rankedPaths[0].Distance}
		return []PathDist{singlePath}, nil // In case only one match was found, no need to procced to Levenshtein OR Score filters
	} else {
		// Sort the rankedPaths by their Distance attribute and return the two min results.
		var TwoMinDistances []PathDist
		sort.Slice(rankedPaths, func(i,j int) bool {
			return rankedPaths[i].Distance < rankedPaths[j].Distance
		})

		for _, path := range rankedPaths {
			TwoMinDistances = append(TwoMinDistances, PathDist{Path: pathMap[path.Target], LevDistance: path.Distance})
		}
		return []PathDist{TwoMinDistances[0], TwoMinDistances[1]}, nil
		// levenshteinPaths := FindLevenshteinDistance(path, foundPaths, findTwoMin, pathMap) // Should get a list of two elements for future score filtering
		// return levenshteinPaths, nil
	}


	// if len(foundPaths) == 0 {
	// 	// Should exit the program with a message to the user
	// 	return []PathDist{}, errors.New("could not find suitable path") // In case could not find any match

	// } else if len(foundPaths) == 1 {
	// 	singlePath := PathDist{Path: pathMap[foundPaths[0]], LevDistance: 0}
	// 	return []PathDist{singlePath}, nil // In case only one match was found, no need to procced to Levenshtein OR Score filters

	// } else { // In case more then one match was found
	// 	levenshteinPaths := FindLevenshteinDistance(path, foundPaths, findTwoMin, pathMap) // Should get a list of two elements for future score filtering
	// 	return levenshteinPaths, nil
	// }
}
