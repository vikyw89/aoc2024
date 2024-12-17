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

// algo steps
// track counter
// track remaining letters
// track allowed directions (up, down, left, right, ur, ul, dl, dr, any)
// go to 1st coordinate
// check if remaining letter is in the current coordinate
// if no, reset remaining letters and go to next coordinate
// if yes, remove from remaining letters
// if remaining letters is empty, return counter
// get surounding explorable coordinates
// recursively call algo, with new coordinate, remaining letters, allowed directions
// collect counters returned by the algo, sum it all

// break case
// if remaining letter is empty, return 1
