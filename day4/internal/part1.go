package internal

import (
	"fmt"
	"time"
)

var XMAS = []string{"X", "M", "A", "S"}

func Part1(path string) int {
	startTime := time.Now()
	fileString := LoadFile(path)
	matrix := StringToMatrix(fileString)
	// loop the matrix and tally all possible xmas in the matrix
	foundCounter := 0
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			currentCoordinate := Coordinate{
				x: x,
				y: y,
			}

			newCounter := GetCurrentCoordinateLetterCount(matrix, XMAS, currentCoordinate, ANY)
			foundCounter += newCounter
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsed)
	println("Total found counter: ", foundCounter)
	return foundCounter
}
