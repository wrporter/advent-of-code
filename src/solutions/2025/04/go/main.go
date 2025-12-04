package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	count := 0
	for y, row := range grid {
		for x, char := range row {
			if char == '@' {
				numRolls := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 {
							continue
						}
						newY, newX := y+dy, x+dx
						if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[newY]) {
							continue
						}
						if grid[newY][newX] == '@' {
							numRolls++
						}
					}
				}
				if numRolls < 4 {
					count++
				}
			}
		}
	}
	return count
}

func part2(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	rolls := make(map[geometry.Point]bool)

	for y, row := range grid {
		for x, char := range row {
			if char == '@' {
				rolls[*geometry.NewPoint(x, y)] = true
			}
		}
	}

	total := 0
	next := rolls
	var removed int

	for {
		next, removed = removeRolls(next)
		total += removed
		if removed == 0 {
			return total
		}
	}

	return -1
}

func removeRolls(rolls map[geometry.Point]bool) (map[geometry.Point]bool, int) {
	removed := 0
	next := make(map[geometry.Point]bool)

	for roll := range rolls {
		numNeighbors := 0

		for _, delta := range geometry.AllDirectionsModifiers {
			neighbor := geometry.NewPoint(roll.X+delta.X, roll.Y+delta.Y)
			if rolls[*neighbor] {
				numNeighbors++
			}
		}

		if numNeighbors < 4 {
			removed++
		} else {
			next[roll] = true
		}
	}

	return next, removed
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 4, Part1: part1, Part2: part2}
}
