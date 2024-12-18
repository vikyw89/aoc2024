package internal

import (
	"strings"
)

// Part1 processes the input and returns the sum of middle numbers in valid updates
func Part1(path string) int {
	input := LoadFile(path)
	lines := strings.Split(input, "\n")
	rules := ParseRules(lines)
	updates := ParseUpdates(lines)

	sum := 0
	for _, update := range updates {
		if IsValid(update, rules) {
			sum += GetMiddleNumber(update)
		}
	}
	return sum
}
