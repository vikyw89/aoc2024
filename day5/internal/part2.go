package internal

import "strings"

func Part2(path string) int {
	input := LoadFile(path)
	lines := strings.Split(input, "\n")
	rules := ParseRules(lines)
	updates := ParseUpdates(lines)

	// Find invalid updates, sort them, and sum their middle numbers
	sum := 0
	for _, update := range updates {
		if !IsValid(update, rules) {
			sortedUpdate := SortUpdate(update, rules)
			sum += GetMiddleNumber(sortedUpdate)
		}
	}

	return sum
}
