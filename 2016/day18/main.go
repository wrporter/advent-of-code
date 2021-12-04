package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 18
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	firstRow := input[0]
	height := convert.StringToInt(input[1])
	grid := expand(firstRow, height)
	return countSafeTiles(grid)
}

func part2(input []string) interface{} {
	firstRow := input[0]
	height := 400_000
	grid := expand(firstRow, height)
	return countSafeTiles(grid)
}

func countSafeTiles(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, tile := range row {
			if tile == safe {
				count++
			}
		}
	}
	return count
}

func expand(firstRow string, height int) [][]rune {
	grid := make([][]rune, height)
	grid[0] = []rune(firstRow)
	width := len(firstRow)

	for y := 1; y < height; y++ {
		grid[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			grid[y][x] = getTile(grid, x, y)
		}
	}

	return grid
}

func getTile(grid [][]rune, x int, y int) rune {
	left := x-1 >= 0 && grid[y-1][x-1] == trap
	center := grid[y-1][x] == trap
	right := x+1 < len(grid[y-1]) && grid[y-1][x+1] == trap

	if (left && center && !right) ||
		(!left && center && right) ||
		(left && !center && !right) ||
		(!left && !center && right) {
		return trap
	}

	return safe
}

var (
	trap = '^'
	safe = '.'
)
