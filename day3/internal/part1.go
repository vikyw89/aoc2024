package internal

import (
	"fmt"
)

func Part1(path string) {
	fmt.Println("Part 1")
	fileString := LoadFile(path)

	muls := ExtractMulNumbers(fileString)

	multipliedMuls := SumMuls(muls)
	fmt.Println("Total multiplied muls: ", multipliedMuls)

}

