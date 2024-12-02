package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/mymath"
	"aoc/src/lib/go/v2/myslice"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	numSafe := 0

	for _, line := range lines {
		report := convert.ToIntsV2(strings.Fields(line))
		if isSafe(report) {
			numSafe += 1
		}
	}

	return numSafe
}

func part2(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	numSafe := 0

	for _, line := range lines {
		report := convert.ToIntsV2(strings.Fields(line))

		if isSafe(report) {
			numSafe += 1
		} else {
			for i := range report {
				toleratedReport := remove(myslice.Copy(report), i)
				if isSafe(toleratedReport) {
					numSafe += 1
					break
				}
			}
		}
	}

	return numSafe
}

func isSafe(report []int) bool {
	previous := report[0]
	direction := 0
	safe := true

	for i := 1; i < len(report) && safe; i += 1 {
		current := report[i]

		if direction == 0 && current > previous {
			direction = 1
		} else if direction == 0 && current < previous {
			direction = -1
		}

		diff := mymath.Abs(current - previous)
		if current == previous ||
			diff < 1 ||
			diff > 3 ||
			(direction == 1 && current < previous) ||
			(direction == -1 && current > previous) {
			safe = false
		}

		previous = current
	}
	return safe
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 2, Part1: part1, Part2: part2}
}
