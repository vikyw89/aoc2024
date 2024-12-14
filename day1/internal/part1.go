package internal

import (
	"fmt"
)

func Part1() {
	sortedLeftAndRightNumber := GetSortedLeftAndRightNumber()

	// get distance between numbers
	distanceSlice := []int{}
	for i := 0; i < len(sortedLeftAndRightNumber.leftNumberSlice); i++ {
		distance := sortedLeftAndRightNumber.leftNumberSlice[i] - sortedLeftAndRightNumber.secondNumberSlice[i]
		absoluteDistance := Abs(distance)
		distanceSlice = append(distanceSlice, absoluteDistance)
	}

	totalDistance := 0
	for _, distance := range distanceSlice {
		totalDistance += distance
	}

	fmt.Println("Total distance: ", totalDistance)
}
