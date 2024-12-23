package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 1
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

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
		if digit == line[(i+1)%len(line)] {
			sum += convert.RuneToInt(digit)
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
			sum += convert.RuneToInt(digit)
		}
	}

	return sum
}
