package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")

	beams := make(map[geometry.Point]bool)
	splitters := make(map[geometry.Point]bool)
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
				beams[*geometry.NewPoint(x, 0)] = true
			}
			if grid[y][x] == '^' {
				splitters[*geometry.NewPoint(x, y)] = true
			}
		}
	}

	totalSplits := 0
	for len(beams) > 0 {
		var splitCount int
		beams, splitCount = moveBeams(grid, splitters, beams)
		totalSplits += splitCount
	}

	return totalSplits
}

func moveBeams(grid []string, splitters map[geometry.Point]bool, beams map[geometry.Point]bool) (map[geometry.Point]bool, int) {
	splits := 0
	nextBeams := make(map[geometry.Point]bool)

	for beam := range beams {
		next := beam.Copy()
		next.Down()

		if next.Y < len(grid) {
			if splitters[*next] {
				splits++

				left, right := next.Copy(), next.Copy()
				left.Left()
				right.Right()

				nextBeams[*left] = true
				nextBeams[*right] = true
			} else {
				nextBeams[*next] = true
			}
		}
	}

	return nextBeams, splits
}

func part2(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	beams := make([]int, len(grid[0]))
	for x := range grid[0] {
		if grid[0][x] == 'S' {
			beams[x] = 1
		}
	}

	for y := 1; y < len(grid); y++ {
		row := grid[y]
		for x := range row {
			if row[x] == '^' {
				beams[x-1] += beams[x]
				beams[x+1] += beams[x]
				beams[x] = 0
			}
		}
	}

	return ints.Sum(beams)
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 7, Part1: part1, Part2: part2}
}
