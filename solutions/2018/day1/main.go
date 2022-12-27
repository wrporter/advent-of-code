package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
)

func main() {
	year, day := 2018, 1
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	frequency := 0

	for _, changeStr := range input {
		change := convert.StringToInt(changeStr[1:])
		sign := changeStr[0]
		if sign == '-' {
			change = -change
		}
		frequency += change
	}

	return frequency
}

func part2(input []string) interface{} {
	frequency := 0
	frequencies := make(map[int]bool)

	for i := 0; ; i++ {
		changeStr := input[i%len(input)]
		change := convert.StringToInt(changeStr[1:])
		sign := changeStr[0]
		if sign == '-' {
			change = -change
		}
		frequency += change
		if frequencies[frequency] {
			return frequency
		}
		frequencies[frequency] = true
	}
}
