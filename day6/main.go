package main

import (
	"day6/internals"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	internals.Part1("input.txt")
	duration := time.Since(start)
	fmt.Printf("Time taken: %v\n", duration)

	start = time.Now()
	internals.Part2("input.txt")
	duration = time.Since(start)
	fmt.Printf("Time taken: %v\n", duration)
}
