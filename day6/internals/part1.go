package internals

import (
	"fmt"
)

func Part1(fileName string) {
	fmt.Println("Part 1")
	fileString := LoadFile(fileName)

	guardState, mapMatrix := GetMapMatrixAndGuardState(fileString)

	visitedGuardCoordinates := GetGuardVisitedCoordinates(guardState, mapMatrix)

	uniqueVisitedCoordinates := GetUniqueVisitedCoordinates(visitedGuardCoordinates)

	fmt.Println("Total unique visited coordinates: ", len(uniqueVisitedCoordinates))
}
