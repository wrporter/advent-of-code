package main

import (
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
)

func main() {
	out.Day(2020, 4)
	input, _ := file.ReadFile("./2020/day4/input.txt")

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	return 0
}

func part2(input []string) interface{} {
	return 0
}
