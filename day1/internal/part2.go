package internal

import (
	"fmt"
)

func Part2() {
	sortedLeftAndRightNumber := GetSortedLeftAndRightNumber()

	similarityMap := make(map[int]int)

	for i := 0; i < len(sortedLeftAndRightNumber.leftNumberSlice); i++ {
		firstNumber := sortedLeftAndRightNumber.leftNumberSlice[i]

		similarityMap[firstNumber] = 0
	}

	// find similarity
	for i := 0; i < len(sortedLeftAndRightNumber.leftNumberSlice); i++ {
		secondNumber := sortedLeftAndRightNumber.secondNumberSlice[i]

		number, ok := similarityMap[secondNumber]
		if !ok {
			continue
		}

		similarityMap[secondNumber] = number + 1
	}

	// get max similarity
	maxSimilarity := 0
	for key, similarity := range similarityMap {
		maxSimilarity += key * similarity
	}

	fmt.Println("Total similarity: ", maxSimilarity)
}
