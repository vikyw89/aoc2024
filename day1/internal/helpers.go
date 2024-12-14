package internal

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NumberPair struct {
	firstNumber  int
	secondNumber int
}
type LeftAndRightNumbers struct {
	leftNumberSlice   []int
	secondNumberSlice []int
}

func GetSortedLeftAndRightNumber() LeftAndRightNumbers {
	fileByte, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileString := string(fileByte)

	inputLines := splitByNewLine(fileString)

	leftNumberSlice := []int{}
	rightNumberSlice := []int{}

	for _, line := range inputLines {
		numberPair := extractNumberPair(line)
		leftNumberSlice = append(leftNumberSlice, numberPair.firstNumber)
		rightNumberSlice = append(rightNumberSlice, numberPair.secondNumber)
	}

	// sort numbers
	sort.Ints(leftNumberSlice)
	sort.Ints(rightNumberSlice)

	return LeftAndRightNumbers{
		leftNumberSlice:   leftNumberSlice,
		secondNumberSlice: rightNumberSlice,
	}
}

func extractNumberPair(lineString string) NumberPair {
	// split by space
	numberPairSplit := splitBySpace(lineString)

	firstNumberString := numberPairSplit[0]
	secondNumberString := numberPairSplit[len(numberPairSplit)-1]

	firstNumber, _ := strconv.Atoi(firstNumberString)
	secondNumber, _ := strconv.Atoi(secondNumberString)

	return NumberPair{
		firstNumber:  firstNumber,
		secondNumber: secondNumber,
	}
}

func splitByNewLine(fileString string) []string {
	return strings.Split(fileString, "\n")
}

func splitBySpace(lineString string) []string {
	return strings.Split(lineString, " ")
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
