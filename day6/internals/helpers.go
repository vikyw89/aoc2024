package internals

import (
	"strings"

	"github.com/jinzhu/copier"
)

func GetMapMatrixAndGuardState(fileString string) (GuardState, MapMatrix) {
	// rows
	rowSlice := strings.Split(fileString, "\n")

	guardState := GuardState{}

	mapMatrix := MapMatrix{}

	for y := 0; y < len(rowSlice); y++ {
		row := rowSlice[y]

		// Skip empty lines
		if len(strings.TrimSpace(row)) == 0 {
			continue
		}

		rowCells := []Cell{}
		for x := 0; x < len(row); x++ {
			char := row[x]

			if char == '#' {
				rowCells = append(rowCells, OBSTACLE)
				continue
			}

			rowCells = append(rowCells, EMPTY)

			if char == '^' {
				guardState.FacingDirection = UP
				guardState.Coordinate = Coordinate{
					x: x,
					y: y,
				}
				continue
			}

			if char == 'v' {
				guardState.FacingDirection = DOWN
				guardState.Coordinate = Coordinate{
					x: x,
					y: y,
				}
				continue
			}

			if char == '>' {
				guardState.FacingDirection = RIGHT
				guardState.Coordinate = Coordinate{
					x: x,
					y: y,
				}
				continue
			}

			if char == '<' {
				guardState.FacingDirection = LEFT
				guardState.Coordinate = Coordinate{
					x: x,
					y: y,
				}
				continue
			}
		}

		mapMatrix = append(mapMatrix, rowCells)
	}
	return guardState, mapMatrix
}

func GetGuardNextState(guardState GuardState, mapMatrix MapMatrix) (GuardState, bool) {

	currentGuardState := GuardState{}
	copier.Copy(&currentGuardState, &guardState)

	turnCounter := 0
	for {

		// check if turned back to original direction
		if turnCounter >= 4 {
			return currentGuardState, false
		}

		// move guard once
		if currentGuardState.FacingDirection == UP {
			currentGuardState.Coordinate.y--
		}
		if currentGuardState.FacingDirection == DOWN {
			currentGuardState.Coordinate.y++
		}
		if currentGuardState.FacingDirection == RIGHT {
			currentGuardState.Coordinate.x++
		}
		if currentGuardState.FacingDirection == LEFT {
			currentGuardState.Coordinate.x--
		}

		// out of bound check
		if currentGuardState.Coordinate.x < 0 || currentGuardState.Coordinate.x > len(mapMatrix[0])-1 {
			return currentGuardState, false
		}
		if currentGuardState.Coordinate.y < 0 || currentGuardState.Coordinate.y > len(mapMatrix)-1 {
			return currentGuardState, false
		}

		// obstacle check
		if mapMatrix[currentGuardState.Coordinate.y][currentGuardState.Coordinate.x] == EMPTY {

			return currentGuardState, true
		}

		// obstacle, need to backtrack coordinate and rotate
		currentGuardState.Coordinate = guardState.Coordinate

		turnCounter++
		if currentGuardState.FacingDirection == UP {
			currentGuardState.FacingDirection = RIGHT
			continue
		}
		if currentGuardState.FacingDirection == RIGHT {
			currentGuardState.FacingDirection = DOWN
			continue
		}
		if currentGuardState.FacingDirection == DOWN {
			currentGuardState.FacingDirection = LEFT
			continue
		}
		if currentGuardState.FacingDirection == LEFT {
			currentGuardState.FacingDirection = UP
			continue
		}

	}
}

func GetGuardVisitedCoordinates(guardState GuardState, mapMatrix MapMatrix) []Coordinate {
	visitedCoordinates := []Coordinate{}

	visitedCoordinates = append(visitedCoordinates, guardState.Coordinate)

	currenGuardState := GuardState{}
	copier.Copy(&currenGuardState, &guardState)

	for {

		nextGuardState, ok := GetGuardNextState(currenGuardState, mapMatrix)

		if !ok {
			return visitedCoordinates
		}

		visitedCoordinates = append(visitedCoordinates, nextGuardState.Coordinate)
		currenGuardState = nextGuardState
	}
}

func GetUniqueVisitedCoordinates(visitedCoordinates []Coordinate) []Coordinate {

	uniqueCoordinatesMap := make(map[Coordinate]bool)

	uniqueCoordinates := []Coordinate{}
	for i := 0; i < len(visitedCoordinates); i++ {
		coordinate := visitedCoordinates[i]

		if uniqueCoordinatesMap[coordinate] {
			continue
		}

		uniqueCoordinatesMap[coordinate] = true
		uniqueCoordinates = append(uniqueCoordinates, coordinate)
	}

	return uniqueCoordinates
}

func IsGuardLooping(guardState GuardState, mapMatrix MapMatrix) bool {
	guardStateMap := make(map[GuardState]bool)

	currentGuardState := GuardState{}
	copier.Copy(&currentGuardState, &guardState)

	for {

		nextGuardState, ok := GetGuardNextState(currentGuardState, mapMatrix)

		if !ok {
			return false
		}

		// if guard state matches existing, return true

		if guardStateMap[nextGuardState] {

			return true
		}

		guardStateMap[nextGuardState] = true

		currentGuardState = nextGuardState

	}
}

func GetPossibleObstacleCountToLoopGuard(mapMatrix MapMatrix, guardState GuardState) int {
	currentGuardState := GuardState{}
	copier.Copy(&currentGuardState, &guardState)

	visitedCoordinates := make(map[Coordinate]bool)
	visitedCoordinates[currentGuardState.Coordinate] = true

	possibleObstacleCount := 0
	for {
		// place a new obstacle
		// create a cloned map with new obstacle
		nextGuardState, ok := GetGuardNextState(currentGuardState, mapMatrix)
		currentGuardState = nextGuardState

		if !ok {
			return possibleObstacleCount
		}

		// if already visited, continue
		if visitedCoordinates[currentGuardState.Coordinate] {
			continue
		}

		visitedCoordinates[currentGuardState.Coordinate] = true

		// place obstacle
		mapMatrix[currentGuardState.Coordinate.y][currentGuardState.Coordinate.x] = OBSTACLE

		// check if guard is looping
		// use initial guard state
		if IsGuardLooping(guardState, mapMatrix) {
			possibleObstacleCount++
		}

		// remove obstacle
		mapMatrix[currentGuardState.Coordinate.y][currentGuardState.Coordinate.x] = EMPTY

	}
}
