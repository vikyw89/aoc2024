package main

import (
	"day6/internals"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	internals.Part1("test.txt")
	duration := time.Since(start)
	fmt.Printf("Time taken: %v\n", duration)
	// internal.Part2("input.txt")

	// start = time.Now()
	// internals.Part2("test.txt")
	// duration = time.Since(start)
	// fmt.Printf("Time taken: %v\n", duration)
}
