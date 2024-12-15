package internal

import "fmt"

type SafeTallyPart2 struct {
	safeReportLinesCount   int
	unsafeReportLinesCount int
}

func Part2() {
	reportLines := GetReportLines()

	safeTallyPart2 := SafeTallyPart2{
		safeReportLinesCount:   0,
		unsafeReportLinesCount: 0,
	}

	for i := 0; i < len(reportLines); i++ {
		reportLine := reportLines[i]

		newMap := make(map[string]bool)

		if isSafeWithDampener(reportLine, 1, newMap) {
			safeTallyPart2.safeReportLinesCount++
		} else {
			safeTallyPart2.unsafeReportLinesCount++
		}
	}
	fmt.Println("Total safe report lines with dampener: ", safeTallyPart2.safeReportLinesCount)
}

func isSafeWithDampener(reportLine ReportLine, remainingDampener int, memo map[string]bool) bool {
	// cache
	cacheKey := fmt.Sprintf("%v-%v", reportLine, remainingDampener)

	if memo[cacheKey] {
		return memo[cacheKey]
	}

	// break case
	fmt.Println("reportLine: ", reportLine)
	if len(reportLine.rowSlice) < 2 {
		return false
	}

	if remainingDampener < 0 {
		return false
	}

	diffSlice := []int{}
	for i := 1; i < len(reportLine.rowSlice); i++ {
		previousValue := reportLine.rowSlice[i-1]
		currentValue := reportLine.rowSlice[i]
		diff := currentValue - previousValue
		diffSlice = append(diffSlice, diff)
	}

	// all increasing or all decreasing check
	// we grouped all positive and negative diffs together
	isSafe := true

	isIncreasing := true

	if diffSlice[0] < 0 {
		isIncreasing = false
	}

	for i := 0; i < len(diffSlice); i++ {
		if isIncreasing && diffSlice[i] < 0 {
			isSafe = false
			break
		}
		if !isIncreasing && diffSlice[i] > 0 {
			isSafe = false
			break
		}
		absoluteDiff := abs(diffSlice[i])
		if absoluteDiff > 3 || absoluteDiff < 1 {
			isSafe = false
			break
		}
	}

	if isSafe {
		return true
	}

	// recursive case
	for i := 0; i < len(reportLine.rowSlice); i++ {
		copiedSlice := make([]int, len(reportLine.rowSlice))
		copy(copiedSlice, reportLine.rowSlice)
		newReportLineRowSlice := append(copiedSlice[:i], copiedSlice[i+1:]...)

		newReportLine := ReportLine{
			rowSlice: newReportLineRowSlice,
		}

		isSafeResult := isSafeWithDampener(newReportLine, remainingDampener-1, memo)

		// fast bubling up
		if isSafeResult {
			return true
		}
	}

	memo[cacheKey] = isSafe
	return isSafe
}
