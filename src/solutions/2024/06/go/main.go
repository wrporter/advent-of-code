package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid, guard := parse(input)
	steps, _ := precompute(grid, guard)
	return len(steps)
}

func part2(input string, _ ...interface{}) interface{} {
	grid, guard := parse(input)
	count := 0

	potentialObstacleSpots, jumps := precompute(grid, guard.Copy())
	for p := range potentialObstacleSpots {
		x, y := p.X, p.Y
		if grid[y][x] == '.' && !p.Equals(&guard.Point) {
			grid[y][x] = '#'
			if _, hasCycle := moveGuard(grid, guard.Copy(), *geometry.NewPoint(x, y), jumps); hasCycle {
				count++
			}
			grid[y][x] = '.'
		}
	}

	return count
}

func precompute(grid [][]rune, guard *geometry.Vector) (steps map[geometry.Point]bool, jumps map[geometry.Vector]*geometry.Point) {
	steps = make(map[geometry.Point]bool)
	jumps = make(map[geometry.Vector]*geometry.Point)
	lastJump := guard.Copy()

	for isInBounds(grid, guard.Point) {
		steps[guard.Point] = true

		p := guard.Copy().Move().Point

		if isInBounds(grid, p) && grid[p.Y][p.X] == '#' {
			next := guard.Copy()
			jumps[*lastJump] = &next.Point
			lastJump = next
			guard.Rotate(90)
		} else {
			guard.Move()
		}
	}

	return steps, jumps
}

func moveGuard(grid [][]rune, guard *geometry.Vector, obstacle geometry.Point, jumps map[geometry.Vector]*geometry.Point) (steps map[geometry.Point]bool, hasCycle bool) {
	seen := make(map[geometry.Vector]bool)

	for isInBounds(grid, guard.Point) {
		if seen[*guard] {
			return steps, true
		}
		seen[*guard] = true

		if next := jumps[*guard]; next != nil &&
			next.Y != obstacle.Y && next.X != obstacle.X {
			guard.Point = *next
			guard.Rotate(90)
			continue
		}

		p := guard.Copy().Move().Point
		if isInBounds(grid, p) && grid[p.Y][p.X] == '#' {
			guard.Rotate(90)
		} else {
			guard.Move()
		}
	}

	return steps, false
}

func parse(input string) ([][]rune, *geometry.Vector) {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	var guard *geometry.Vector

	for y, row := range grid {
		for x, cell := range row {
			if geometry.IsArrow(cell) {
				grid[y][x] = '.'
				guard = geometry.NewVector(x, y, geometry.NewDirection(cell))
			}
		}
	}
	return grid, guard
}

func isInBounds(grid [][]rune, p geometry.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(grid) && p.X < len(grid[p.Y])
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 6, Part1: part1, Part2: part2}
}
