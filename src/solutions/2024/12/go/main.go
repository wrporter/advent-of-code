package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	seen := make(map[geometry.Point]bool)
	total := 0

	for y, row := range grid {
		for x := range row {
			start := *geometry.NewPoint(x, y)
			if seen[start] {
				continue
			}
			seen[start] = true

			queue := []geometry.Point{start}
			area, perimeter := 0, 0
			plant := grid[y][x]

			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				area += 1

				for _, d := range geometry.Directions {
					next := *p.Copy().Move(d)

					if !isInBounds(grid, next) || grid[next.Y][next.X] != plant {
						perimeter += 1
					} else if isInBounds(grid, next) && grid[next.Y][next.X] == plant && !seen[next] {
						seen[next] = true
						queue = append(queue, next)
					}
				}
			}

			total += area * perimeter
		}
	}

	return total
}

func part2(input string, _ ...interface{}) interface{} {
	seen := make(map[geometry.Point]bool)
	grid := make(map[geometry.Point]rune)
	for y, row := range strings.Split(input, "\n") {
		for x, char := range row {
			grid[*geometry.NewPoint(x, y)] = char
		}
	}

	total := 0
	for start := range grid {
		if !seen[start] {
			queue := []geometry.Point{start}
			area, sides := 0, 0
			plant := grid[start]
			seen[start] = true

			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				area += 1

				for _, d := range []*geometry.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
					next := *p.Copy().Add(d)
					rotate := *p.Copy().Add(geometry.NewPoint(-d.Y, d.X))

					if grid[next] != plant && (grid[rotate] != plant || grid[*rotate.Add(d)] == plant) {
						sides += 1
					} else if grid[next] == plant && !seen[next] {
						seen[next] = true
						queue = append(queue, next)
					}
				}
			}

			total += area * sides
		}
	}

	return total
}

func isInBounds(grid []string, next geometry.Point) bool {
	return next.Y >= 0 && next.X >= 0 &&
		next.Y < len(grid) && next.X < len(grid[next.Y])
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 12, Part1: part1, Part2: part2}
}
