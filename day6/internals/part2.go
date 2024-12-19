package internals

import (
	"fmt"
)

func Part2(fileName string) {
	fmt.Println("Part 2")
	fileString := LoadFile(fileName)

	guardState, mapMatrix := GetMapMatrixAndGuardState(fileString)

	possibleObstacleCountToLoopGuard := GetPossibleObstacleCountToLoopGuard(mapMatrix, guardState)

	fmt.Println("Possible obstacle count to loop guard: ", possibleObstacleCountToLoopGuard)

}
