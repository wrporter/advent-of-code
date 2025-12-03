package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	total := 0
	for _, bank := range strings.Split(input, "\n") {
		highest := 0
		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				joltage := convert.StringToInt(string(bank[i]) + string(bank[j]))
				if joltage > highest {
					highest = joltage
				}
			}
		}
		total += highest
	}
	return total
}

func part2(input string, _ ...interface{}) interface{} {
	total := 0
	for _, bank := range strings.Split(input, "\n") {
		total += findMax(bank, 12)
	}
	return total
}

// Greedy algorithm that searches for the highest digit for each battery.
// 1. Remove remaining potential batteries from the search space.
// 2. Find the highest digit in the search space.
// This solution can be used for part 1 as well, but I will leave my original
// solution.
func findMax(bank string, numBatteries int) int {
	var builder strings.Builder
	start := 0

	for battery := 0; battery < numBatteries; battery++ {
		remaining := numBatteries - battery
		maxSearchLen := len(bank) - (remaining - 1)

		maxBattery := '0'
		maxIndex := -1
		for i := start; i < maxSearchLen; i++ {
			current := bank[i]
			if current > byte(maxBattery) {
				maxBattery = rune(current)
				maxIndex = i
			}
		}

		builder.WriteByte(bank[maxIndex])
		start = maxIndex + 1
	}

	return convert.StringToInt(builder.String())
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 3, Part1: part1, Part2: part2}
}
