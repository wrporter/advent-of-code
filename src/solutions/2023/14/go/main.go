package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/runegrid"
	"aoc/src/lib/go/v2/myslice"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	rollRocks(grid)
	return calculateNorthLoad(grid)
}

func part2(input string, _ ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	seen := make(map[string]int)
	cycles := 1_000_000_000

	for cycle := 1; cycle <= cycles; cycle++ {
		for tilt := 1; tilt <= 4; tilt++ {
			rollRocks(grid)
			myslice.Rotate90(grid)
		}

		key := runegrid.String(grid)
		if start, ok := seen[key]; ok {
			cycle += getCycleJumpLength(start, cycle, cycles)
		}

		seen[key] = cycle
	}

	return calculateNorthLoad(grid)
}

func getCycleJumpLength(start, current, total int) int {
	length := current - start
	remaining := total - current
	return remaining / length * length
}

func rollRocks(grid [][]rune) {
	for x := 0; x < len(grid[0]); x++ {
		empty := 0

		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 'O' {
				grid[y][x] = '.'
				grid[empty][x] = 'O'
				empty += 1
			} else if grid[y][x] == '#' {
				empty = y + 1
			}
		}
	}
}

func calculateNorthLoad(grid [][]rune) interface{} {
	sum := 0
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'O' {
				load := len(grid) - y
				sum += load
			}
		}
	}
	return sum
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 14, Part1: part1, Part2: part2}
}
