package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
	"aoc/src/lib/go/v2/mymath"
	"strings"
)

func part1(input string, args ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	steps := args[0].(int)
	return countPlots(grid, -1, steps)
}

func part2(input string, args ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	cycle := args[0].(int)
	steps := args[1].(int)
	return countPlots(grid, cycle, steps)
}

func countPlots(grid [][]rune, cycle, steps int) interface{} {
	size := len(grid)

	var start *geometry.Point
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
				start = geometry.NewPoint(x, y)
				grid[y][x] = '.'
			}
		}
	}

	plots := map[geometry.Point]struct{}{*start: {}}
	var polynomial []int

	for step := 0; step <= steps; step++ {
		// Part 1
		if step == steps {
			return len(plots)
		}

		// Part 2
		if step%size == cycle {
			polynomial = append(polynomial, len(plots))
		}
		if len(polynomial) == 3 {
			break
		}

		nextPlots := make(map[geometry.Point]struct{})
		for plot := range plots {
			for _, direction := range geometry.Directions {
				next := plot.Copy().Move(direction)

				if grid[mymath.WrapMod(next.Y, size)][mymath.WrapMod(next.X, size)] != '#' {
					nextPlots[*next] = struct{}{}
				}
			}
		}
		plots = nextPlots
	}

	// Part 2
	if len(polynomial) == 3 {
		return lagrange(steps/size, polynomial[0], polynomial[1], polynomial[2])
	}

	return -1
}

func lagrange(n, a, b, c int) int {
	return a + n*(b-a+(n-1)*(c-b-b+a)/2)
}

func main() {
	New().Run([]interface{}{64}, []interface{}{65, 26501365})
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 21, Part1: part1, Part2: part2}
}
