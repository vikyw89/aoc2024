package internal

import (
	"strings"
)

func Part2(path string) int {
	fileString := LoadFile(path)

	validOperatorSlice := ExtractValidOperators(fileString)

	isDo := true
	sum := 0

	for i := 0; i < len(validOperatorSlice); i++ {
		operator := validOperatorSlice[i]
		if strings.HasPrefix(operator, "do") {
			println("Do found")
			isDo = true
		}

		if strings.HasPrefix(operator, "don't") {
			println("Dont found")
			isDo = false
		}

		// skip if isDo is false
		if !isDo {
			continue
		}

		if strings.HasPrefix(operator, "mul") {
			mulValue := CalculateMulString(operator)

			sum += mulValue

		}

	}
	println("Total operator sum: ", sum)
	return sum
}
