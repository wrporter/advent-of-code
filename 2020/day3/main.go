package main

import (
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
)

type Slope struct {
	DX int
	DY int
}

func main() {
	out.Day(2020, 3)
	input, _ := file.ReadFile("./2020/day3/input.txt")
	answer1 := part1(input, 3, 1)
	out.Part1(answer1)
	answer2 := part2(input, []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	})
	out.Part2(answer2)
}

func part1(input []string, dx int, dy int) int {
	x := dx
	y := dy
	numTrees := 0

	for y < len(input) {
		if input[y][x%len(input[y])] == '#' {
			numTrees++
		}
		x += dx
		y += dy
	}

	return numTrees
}

func part2(input []string, slopes []Slope) int {
	result := 1
	for _, slope := range slopes {
		numTrees := part1(input, slope.DX, slope.DY)
		result *= numTrees
	}
	return result
}
