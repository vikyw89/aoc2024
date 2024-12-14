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

func TotalDistance() {
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

	// get distance between numbers
	distanceSlice := []int{}
	for i := 0; i < len(leftNumberSlice); i++ {
		distance := leftNumberSlice[i] - rightNumberSlice[i]
		absoluteDistance := abs(distance)
		distanceSlice = append(distanceSlice, absoluteDistance)
	}

	totalDistance := 0
	for _, distance := range distanceSlice {
		totalDistance += distance
	}

	fmt.Println("Total distance: ", totalDistance)
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
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

func sortNumberSlice(numberSlice []int) []int {
	return numberSlice
}
