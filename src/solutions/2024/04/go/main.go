package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	word := "XMAS"
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'X' {
				for _, delta := range geometry.AllDirectionsModifiers {
					ny := y
					nx := x
					found := true

					for _, char := range word[1:] {
						ny = ny + delta.Y
						nx = nx + delta.X

						if ny < 0 ||
							nx < 0 ||
							ny >= len(grid) ||
							nx >= len(grid[ny]) ||
							grid[ny][nx] != byte(char) {
							found = false
							break
						}
					}

					if found {
						count++
					}
				}
			}
		}
	}

	return count
}

func part2(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'A' && y > 0 && x > 0 && y+1 < len(grid) && x+1 < len(grid[y]) {
				if ((grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M') || (grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S')) &&
					((grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M') || (grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S')) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 4, Part1: part1, Part2: part2}
}
