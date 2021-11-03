package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 1
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	sum := 0
	line := input[0]

	for i := 0; i < len(line); i++ {
		digit := line[i]
		if (i == len(line)-1 && line[0] == digit) || (i < len(line)-1 && line[i+1] == digit) {
			sum += conversion.RuneToInt(digit)
		}
	}

	return sum
}

func part2(input []string) interface{} {
	sum := 0
	line := input[0]
	distance := len(line) / 2

	for i := 0; i < len(line); i++ {
		digit := line[i]
		next := ints.WrapMod(i+distance, len(line))

		if digit == line[next] {
			sum += conversion.RuneToInt(digit)
		}
	}

	return sum
}
