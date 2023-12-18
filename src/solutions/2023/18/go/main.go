package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	return calculateArea(input, func(line string) (geometry.Point, int) {
		parts := strings.Fields(line)
		delta := directions[parts[0]]
		distance := convert.StringToInt(parts[1])
		return delta, distance
	})
}

func part2(input string, _ ...interface{}) interface{} {
	return calculateArea(input, func(line string) (geometry.Point, int) {
		parts := strings.Fields(line)
		delta := directions[string(parts[2][7])]
		distance := hexToInt(parts[2][2:7])
		return delta, distance
	})
}

func calculateArea(input string, parse func(line string) (geometry.Point, int)) interface{} {
	current := geometry.NewPoint(0, 0)
	perimeter := 0
	area := 0

	for _, line := range strings.Split(input, "\n") {
		delta, distance := parse(line)

		current.X += delta.X * distance
		current.Y += delta.Y * distance
		perimeter += distance

		// Green's Theorem (rate of change bounded by the curve)
		area += current.X * delta.Y * distance
	}

	// Pick's Theorem for interior points
	return area + perimeter/2 + 1
}

var directions = map[string]geometry.Point{
	"R": {1, 0},
	"D": {0, 1},
	"L": {-1, 0},
	"U": {0, -1},

	// Hex encoding
	"0": {1, 0},
	"1": {0, 1},
	"2": {-1, 0},
	"3": {0, -1},
}

func hexToInt(hex string) int {
	value, _ := strconv.ParseInt(hex, 16, 64)
	return int(value)
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 18, Part1: part1, Part2: part2}
}
