package main

import (
	"day5/internal"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Day 5")

	// Run Part 1
	start := time.Now()
	result := internal.Part1("input.txt")
	duration := time.Since(start)
	fmt.Printf("Part 1 Result: %d (took %v)\n", result, duration)

	// Run Part 2
	start = time.Now()
	result = internal.Part2("input.txt")
	duration = time.Since(start)
	fmt.Printf("Part 2 Result: %d (took %v)\n", result, duration)
}
