package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/ints"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	regionLines := lines[30:]
	total := 0

	for _, regionLine := range regionLines {
		parts := strings.Split(regionLine, ": ")

		sizes := strings.Split(parts[0], "x")
		height, _ := strconv.Atoi(sizes[0])
		width, _ := strconv.Atoi(sizes[1])
		area := height * width

		counts, _ := convert.ToInts(strings.Fields(parts[1]))
		sum := ints.Sum(counts)
		requiredArea := sum * 3 * 3

		if area >= requiredArea {
			total++
		}
	}

	return total
}

func part2(_ string, _ ...interface{}) interface{} {
	return "Merry Christmas! 🎄"
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 12, Part1: part1, Part2: part2}
}
