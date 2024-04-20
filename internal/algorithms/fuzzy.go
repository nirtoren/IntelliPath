package algo

import (
	"errors"
	"math"
	"intellipath/internal/db"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type MinFunc func([]int) []int

func FuzzyFind(path string, optionalResults []db.PathRecord) (string, error) {

	var foundPaths []string = fuzzy.Find(path, optionalResults) //Should get a list of optional paths 
	if len(foundPaths) == 0 {
		return "", errors.New("could not find suitable path.")
	} else if len(foundPaths) == 1{
		return foundPaths[0], nil
	} else {
		levenshteinPaths := findMinLevenshteinDistance(path, foundPaths, findTwoMin) // Should get a list of two elements for future score filtering
		heighetScored := 0 // Should filter the "levenshteinPaths" by their score.
	}
	
	return "", nil
}

func findMinLevenshteinDistance(userInput string, foundPaths []string, minFunction MinFunc) {
	// [1,3,5,2,0]
	// Should be {path , distance}
	// Send to minFunction({path, distance})
	// return the objects consisting the path and its distance
	var distances []int
	for _, path := range foundPaths {
		distance := fuzzy.LevenshteinDistance(userInput, path)
		distances = append(distances, distance)
	}

	minimalLevenshtainDistance := minFunction(distances)

}

func findTwoMin(nums []int) []int {
	min1, min2 := math.MaxInt8, math.MaxInt8

	for _, num := range nums {
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	return []int{min1, min2}
}
