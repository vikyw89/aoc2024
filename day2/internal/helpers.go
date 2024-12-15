package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ReportLine struct {
	rowSlice []int
}

func GetReportLines() []ReportLine {
	fileByte, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileString := string(fileByte)

	inputLines := splitByNewLine(fileString)

	reportLines := []ReportLine{}
	for i := 0; i < len(inputLines); i++ {
		rowSlice := splitBySpace(inputLines[i])

		reportLines = append(reportLines, ReportLine{
			rowSlice: rowSlice,
		})
	}
	return reportLines
}

func splitByNewLine(fileString string) []string {
	return strings.Split(fileString, "\n")
}

func splitBySpace(lineString string) []int {
	colSlices := strings.Split(lineString, " ")

	// prune empty slices
	parsedColSlices := []int{}
	for _, colSlice := range colSlices {
		if colSlice == "" {
			continue
		}
		intValue, err := strconv.Atoi(colSlice)

		if err != nil {
			fmt.Println(err)
		}

		parsedColSlices = append(parsedColSlices, intValue)
	}

	return parsedColSlices
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
