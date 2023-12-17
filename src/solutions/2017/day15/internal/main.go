package main

import (
	"aoc/src/solutions/2017/day15/lib/generator"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := generator.CountSynchronously()
	elapsed := time.Since(start)
	fmt.Printf("Result: %d, Time: %v\n", result, elapsed)

	start = time.Now()
	result = generator.CountWithChannels()
	elapsed = time.Since(start)
	fmt.Printf("Result: %d, Time: %v\n", result, elapsed)
}
