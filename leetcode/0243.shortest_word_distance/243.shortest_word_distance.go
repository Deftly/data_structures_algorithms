package data_structures_algorithms

import "math"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	min := len(wordsDict)
	index1, index2 := -1, -1

	for i, w := range wordsDict {
		if w == word1 {
			index1 = i
		}

		if w == word2 {
			index2 = i
		}

		if index1 != -1 && index2 != -1 {
			abs := math.Abs(float64(index1 - index2))
			min = int(math.Min(float64(min), abs))
		}
	}

	return min
}
