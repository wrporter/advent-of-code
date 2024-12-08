package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid, antennas := parse(input)
	antinodes := make(map[geometry.Point]bool)

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				a1, a2 := points[i], points[j]
				dx, dy := a1.Diff(a2)
				f1 := a1.Copy().Add(geometry.NewPoint(dx, dy))
				f2 := a2.Copy().Add(geometry.NewPoint(-dx, -dy))

				if isInMap(grid, f1) {
					antinodes[*f1] = true
				}
				if isInMap(grid, f2) {
					antinodes[*f2] = true
				}
			}
		}
	}

	return len(antinodes)
}

func part2(input string, _ ...interface{}) interface{} {
	grid, antennas := parse(input)
	antinodes := make(map[geometry.Point]bool)

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				a1, a2 := points[i], points[j]
				dx, dy := a1.Diff(a2)
				decrease := geometry.NewPoint(dx, dy)
				increase := geometry.NewPoint(-dx, -dy)

				for f1 := a1.Copy().Add(decrease); isInMap(grid, f1); f1.Add(decrease) {
					antinodes[*f1] = true
				}
				for f2 := a2.Copy().Add(increase); isInMap(grid, f2); f2.Add(increase) {
					antinodes[*f2] = true
				}

				antinodes[*a1] = true
				antinodes[*a2] = true
			}
		}
	}

	return len(antinodes)
}

func isInMap(grid []string, p *geometry.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(grid) && p.X < len(grid[p.Y])
}

func parse(input string) ([]string, map[rune][]*geometry.Point) {
	antennas := make(map[rune][]*geometry.Point)
	grid := strings.Split(input, "\n")
	for y, line := range grid {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], geometry.NewPoint(x, y))
			}
		}
	}
	return grid, antennas
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 8, Part1: part1, Part2: part2}
}
