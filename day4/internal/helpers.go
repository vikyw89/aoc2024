package internal

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	DIAGONAL_UP_LEFT
	DIAGONAL_UP_RIGHT
	DIAGONAL_DOWN_LEFT
	DIAGONAL_DOWN_RIGHT
	ANY
)

var DirectionName = map[Direction]string{
	UP:                  "up",
	DOWN:                "down",
	LEFT:                "left",
	RIGHT:               "right",
	DIAGONAL_UP_LEFT:    "diagonal up left",
	DIAGONAL_UP_RIGHT:   "diagonal up right",
	DIAGONAL_DOWN_LEFT:  "diagonal down left",
	DIAGONAL_DOWN_RIGHT: "diagonal down right",
}

func LoadFile(path string) string {
	fileByte, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	fileString := string(fileByte)

	return fileString
}

func StringToMatrix(stringContent string) [][]string {

	rows := strings.Split(stringContent, "\n")

	matrix := make([][]string, len(rows))

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		rowSlice := strings.Split(row, "")
		matrix[i] = rowSlice
	}

	return matrix
}

type Coordinate struct {
	x int
	y int
}

func GetSuroundingExplorableCoordinates(matrix [][]string, currentCoordinate Coordinate, allowedDirections Direction) []Coordinate {

	explorableCoordinates := []Coordinate{}

	if allowedDirections == ANY {
		for xAxis := -1; xAxis < 2; xAxis++ {
			for yAxis := -1; yAxis < 2; yAxis++ {
				// skip current coordinate
				if xAxis == 0 && yAxis == 0 {
					continue
				}

				// if out of bound, skip
				if currentCoordinate.x+xAxis < 0 || currentCoordinate.x+xAxis >= len(matrix[0]) {
					continue
				}
				if currentCoordinate.y+yAxis < 0 || currentCoordinate.y+yAxis >= len(matrix) {
					continue
				}

				// add
				explorableCoordinates = append(explorableCoordinates, Coordinate{
					x: currentCoordinate.x + xAxis,
					y: currentCoordinate.y + yAxis,
				})
			}
		}
		return explorableCoordinates
	}

	if allowedDirections == UP {
		// if out of bound, skip
		if currentCoordinate.y-1 < 0 {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x,
			y: currentCoordinate.y - 1,
		})
		return explorableCoordinates
	}

	if allowedDirections == DOWN {
		// if out of bound, skip
		if currentCoordinate.y+1 >= len(matrix) {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x,
			y: currentCoordinate.y + 1,
		})
		return explorableCoordinates
	}

	if allowedDirections == LEFT {
		// if out of bound, skip
		if currentCoordinate.x-1 < 0 {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x - 1,
			y: currentCoordinate.y,
		})
		return explorableCoordinates
	}

	if allowedDirections == RIGHT {
		// if out of bound, skip
		if currentCoordinate.x+1 >= len(matrix[0]) {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x + 1,
			y: currentCoordinate.y,
		})
		return explorableCoordinates
	}

	if allowedDirections == DIAGONAL_UP_LEFT {
		// if out of bound, skip
		if currentCoordinate.x-1 < 0 || currentCoordinate.y-1 < 0 {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x - 1,
			y: currentCoordinate.y - 1,
		})
		return explorableCoordinates
	}

	if allowedDirections == DIAGONAL_UP_RIGHT {
		// if out of bound, skip
		if currentCoordinate.x+1 >= len(matrix[0]) || currentCoordinate.y-1 < 0 {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x + 1,
			y: currentCoordinate.y - 1,
		})
		return explorableCoordinates
	}

	if allowedDirections == DIAGONAL_DOWN_LEFT {
		// if out of bound, skip
		if currentCoordinate.x-1 < 0 || currentCoordinate.y+1 >= len(matrix) {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x - 1,
			y: currentCoordinate.y + 1,
		})
		return explorableCoordinates
	}

	if allowedDirections == DIAGONAL_DOWN_RIGHT {
		// if out of bound, skip
		if currentCoordinate.x+1 >= len(matrix[0]) || currentCoordinate.y+1 >= len(matrix) {
			return explorableCoordinates
		}
		explorableCoordinates = append(explorableCoordinates, Coordinate{
			x: currentCoordinate.x + 1,
			y: currentCoordinate.y + 1,
		})
		return explorableCoordinates
	}

	return explorableCoordinates
}

func GetNewAllowedDirections(currentCoordinate Coordinate, newCoordinate Coordinate, allowedDirections Direction) Direction {
	// Check diagonal moves
	if newCoordinate.x > currentCoordinate.x && newCoordinate.y > currentCoordinate.y {
		return DIAGONAL_DOWN_RIGHT
	}
	if newCoordinate.x < currentCoordinate.x && newCoordinate.y > currentCoordinate.y {
		return DIAGONAL_DOWN_LEFT
	}
	if newCoordinate.x > currentCoordinate.x && newCoordinate.y < currentCoordinate.y {
		return DIAGONAL_UP_RIGHT
	}
	if newCoordinate.x < currentCoordinate.x && newCoordinate.y < currentCoordinate.y {
		return DIAGONAL_UP_LEFT
	}

	// Check straight moves
	if newCoordinate.x > currentCoordinate.x {
		return RIGHT
	}
	if newCoordinate.x < currentCoordinate.x {
		return LEFT
	}
	if newCoordinate.y > currentCoordinate.y {
		return DOWN
	}
	if newCoordinate.y < currentCoordinate.y {
		return UP
	}

	return allowedDirections
}

func GetCurrentCoordinateLetterCount(matrix [][]string, remainingLetterSlice []string, currentCoordinate Coordinate, allowedDirections Direction) int {
	// break case
	// if remaining letters is empty, we've found a complete path
	nextLetter := remainingLetterSlice[0]
	currentLetter := matrix[currentCoordinate.y][currentCoordinate.x]

	if nextLetter == currentLetter && len(remainingLetterSlice) == 1 {
		return 1
	}

	if nextLetter != currentLetter {
		return 0
	}

	// recursive case
	// get new remaining letters
	newRemainingLetterSlice := remainingLetterSlice[1:]

	// get new coordinates
	nextCoordinates := GetSuroundingExplorableCoordinates(matrix, currentCoordinate, allowedDirections)

	// recursively call algo
	// total counter
	foundCounter := 0
	for _, val := range nextCoordinates {
		newAllowedDirections := GetNewAllowedDirections(currentCoordinate, val, allowedDirections)

		// call algo
		foundCounter += GetCurrentCoordinateLetterCount(matrix, newRemainingLetterSlice, val, newAllowedDirections)
	}

	return foundCounter
}

func IsValidXmas(matrix [][]string, currentCoordinate Coordinate) bool {
	// Check if we can safely check diagonals (need at least 1 space on all sides)
	if currentCoordinate.y <= 0 || currentCoordinate.y >= len(matrix)-1 ||
		currentCoordinate.x <= 0 || currentCoordinate.x >= len(matrix[currentCoordinate.y])-1 {
		return false
	}

	// invalidate if current coordinate is not A
	if matrix[currentCoordinate.y][currentCoordinate.x] != "A" {
		return false
	}

	// retrieve diagonals
	topLeftChar := matrix[currentCoordinate.y-1][currentCoordinate.x-1]
	topRightChar := matrix[currentCoordinate.y-1][currentCoordinate.x+1]
	bottomLeftChar := matrix[currentCoordinate.y+1][currentCoordinate.x-1]
	bottomRightChar := matrix[currentCoordinate.y+1][currentCoordinate.x+1]

	if !((topLeftChar == "M" && bottomRightChar == "S") || (topLeftChar == "S" && bottomRightChar == "M")) {
		return false
	}

	if !((topRightChar == "M" && bottomLeftChar == "S") || (topRightChar == "S" && bottomLeftChar == "M")) {
		return false
	}

	return true
}
