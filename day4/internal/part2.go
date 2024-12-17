package internal

import (
	"fmt"
	"time"
)

func Part2(path string) int {
	startTime := time.Now()
	fileString := LoadFile(path)

	matrix := StringToMatrix(fileString)

	counter := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y])-1; x++ {
			currentCoordinate := Coordinate{
				x: x,
				y: y,
			}

			if IsValidXmas(matrix, currentCoordinate) {
				counter++
			}
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsed)
	println("Total found counter: ", counter)
	return counter
}
