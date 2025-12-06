package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		grid[i], _ = convert.ToInts(strings.Fields(line))
	}

	operators := strings.Fields(lines[len(lines)-1])

	total := 0
	for x, operator := range operators {
		result := 0
		if operator == "*" {
			result = 1
		}

		for y := 0; y < len(grid); y++ {
			num := grid[y][x]

			if operator == "*" {
				result *= num
			} else if operator == "+" {
				result += num
			}
		}

		total += result
	}

	return total
}

func part2(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	total := 0

	var nums []int
	for x := len(grid[0]) - 1; x >= 0; x-- {
		var sb strings.Builder
		for y := 0; y < len(grid)-1; y++ {
			if grid[y][x] != ' ' {
				sb.WriteRune(rune(grid[y][x]))
			}
		}

		if sb.String() == "" {
			// This is an empty column between math problems
			continue
		}
		nums = append(nums, convert.StringToInt(sb.String()))

		operator := grid[len(grid)-1][x]
		if operator != ' ' {
			result := 0
			if operator == '*' {
				result = 1
			}

			for _, num := range nums {
				if operator == '*' {
					result *= num
				} else if operator == '+' {
					result += num
				}
			}

			nums = make([]int, 0)
			total += result
		}
	}

	return total
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 6, Part1: part1, Part2: part2}
}
