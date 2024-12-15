package internal

import "fmt"

type SafeTally struct {
	safeReportLinesCount   int
	unsafeReportLinesCount int
}

func Part1() {
	reportLines := GetReportLines()

	tally := SafeTally{
		safeReportLinesCount:   0,
		unsafeReportLinesCount: 0,
	}

	// check per lines
	for i := 0; i < len(reportLines); i++ {

		if isSafe(reportLines[i]) {
			tally.safeReportLinesCount++
		} else {
			tally.unsafeReportLinesCount++
		}
	}

	fmt.Println("Total safe report lines: ", tally.safeReportLinesCount)

}

func isSafe(reportLine ReportLine) bool {
	// if len(reportLine.rowSlice) < 2 {
	// 	return false
	// }
	previousChange := 0

	for i := 0; i < len(reportLine.rowSlice)-1; i++ {
		change := reportLine.rowSlice[i+1] - reportLine.rowSlice[i]

		// 1st condition, all increasing or all decreasing
		if (previousChange > 0 && change < 0) || (previousChange < 0 && change > 0) {
			return false
		}

		absChange := abs(change)

		// change is more than 3 or less than 1
		if absChange > 3 || absChange < 1 {
			return false
		}

		previousChange = change
	}
	return true
}
