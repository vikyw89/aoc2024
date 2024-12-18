package internal

import (
	"strconv"
	"strings"
)

// ruleMap stores the mapping of Before values to their After values
var ruleMap map[int][]int

// BuildRuleMap creates a map from the rules for faster lookups
func BuildRuleMap(rules []Rule) {
	ruleMap = make(map[int][]int)
	for _, rule := range rules {
		ruleMap[rule.Before] = append(ruleMap[rule.Before], rule.After)
	}
}

// ParseRules parses the rules from input lines
func ParseRules(lines []string) []Rule {
	rules := make([]Rule, 0)
	for _, line := range lines {
		if !strings.Contains(line, "|") {
			continue
		}
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		after, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		rules = append(rules, Rule{Before: before, After: after})
	}
	BuildRuleMap(rules)
	return rules
}

func GetMiddleNumber(update []int) int {
	return update[len(update)/2]
}

// ParseUpdates parses the updates from input lines
func ParseUpdates(lines []string) [][]int {
	updates := make([][]int, 0)
	for _, line := range lines {
		if strings.Contains(line, "|") {
			continue
		}
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		update := make([]int, 0)
		for _, part := range parts {
			num, _ := strconv.Atoi(strings.TrimSpace(part))
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return updates
}

// IsValid checks if an update is valid according to the rules
func IsValid(update []int, rules []Rule) bool {
	for _, rule := range rules {
		beforeIdx := -1
		afterIdx := -1
		for i, num := range update {
			if num == rule.Before {
				beforeIdx = i
			}
			if num == rule.After {
				afterIdx = i
			}
		}
		if beforeIdx != -1 && afterIdx != -1 && beforeIdx > afterIdx {
			return false
		}
	}
	return true
}

// sortUpdate sorts a single update according to the rules
func SortUpdate(update []int, rules []Rule) []int {
	sorted := make([]int, len(update))
	copy(sorted, update)

	// Bubble sort with rules
	n := len(sorted)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// If j should come after j+1 according to rules, swap them
			if ShouldComeBefore(sorted[j+1], sorted[j], rules) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}

// shouldComeBefore checks if a should come before b according to rules
func ShouldComeBefore(a, b int, rules []Rule) bool {
	// Check if a should come before b
	if afters, exists := ruleMap[a]; exists {
		for _, after := range afters {
			if after == b {
				return true
			}
		}
	}

	// Check if b should come before a
	if afters, exists := ruleMap[b]; exists {
		for _, after := range afters {
			if after == a {
				return false
			}
		}
	}

	// If no rules apply, larger number should come first
	return a > b
}
