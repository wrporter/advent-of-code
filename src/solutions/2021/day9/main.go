package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 9
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	heightmap := parseInput(input)

	riskLevel := 0
	for y, row := range heightmap {
		for x, height := range row {
			cur := geometry.NewPoint(x, y)
			if isLowest(heightmap, cur) {
				riskLevel += height + 1
			}
		}
	}

	return riskLevel
}

func part2(input []string) interface{} {
	heightmap := parseInput(input)

	var sizes []int
	for y, row := range heightmap {
		for x := range row {
			start := geometry.NewPoint(x, y)
			if isLowest(heightmap, start) {
				size := getBasinSize(heightmap, make(map[geometry.Point]bool), start)
				sizes = append(sizes, size)
			}
		}
	}

	sort.SliceStable(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
}

func getBasinSize(heightmap [][]int, basin map[geometry.Point]bool, point geometry.Point) int {
	basin[point] = true

	for _, direction := range geometry.Directions {
		next := point.Move(direction)
		if isInBounds(heightmap, next) &&
			!basin[next] &&
			heightmap[next.Y][next.X] < 9 {
			getBasinSize(heightmap, basin, next)
		}
	}

	return len(basin)
}

func parseInput(input []string) [][]int {
	heightmap := make([][]int, len(input))
	for i, line := range input {
		heightmap[i], _ = convert.ToInts(strings.Split(line, ""))
	}
	return heightmap
}

func isLowest(heightmap [][]int, point geometry.Point) bool {
	for _, direction := range geometry.Directions {
		next := point.Move(direction)

		if isInBounds(heightmap, next) &&
			heightmap[point.Y][point.X] >= heightmap[next.Y][next.X] {
			return false
		}
	}
	return true
}

func isInBounds(heightmap [][]int, point geometry.Point) bool {
	x, y := point.X, point.Y
	return y >= 0 && y < len(heightmap) &&
		x >= 0 && x < len(heightmap[y])
}
