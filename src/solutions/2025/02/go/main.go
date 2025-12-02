package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	ranges := strings.Split(input, ",")
	sum := 0

	for _, r := range ranges {
		ids := strings.Split(r, "-")
		start := ids[0]
		end := ids[1]

		for idInt := convert.StringToInt(start); idInt <= convert.StringToInt(end); idInt++ {
			id := strconv.Itoa(idInt)
			if id[:len(id)/2] == id[len(id)/2:] {
				sum += idInt
			}
		}
	}

	return sum
}

func part2(input string, _ ...interface{}) interface{} {
	ranges := strings.Split(input, ",")
	sum := 0

	for _, r := range ranges {
		ids := strings.Split(r, "-")
		start := ids[0]
		end := ids[1]

		for idInt := convert.StringToInt(start); idInt <= convert.StringToInt(end); idInt++ {
			id := strconv.Itoa(idInt)

			for i := 1; i < len(id); i++ {
				sequence := id[0:i]
				next, isFound := strings.CutPrefix(id, sequence)
				for isFound && next != "" {
					next, isFound = strings.CutPrefix(next, sequence)
				}
				if next == "" {
					sum += idInt
					break
				}
			}
		}
	}

	return sum
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 2, Part1: part1, Part2: part2}
}
