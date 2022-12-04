package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/2017/day15/internal/generator"
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
