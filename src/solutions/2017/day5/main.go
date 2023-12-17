package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 5
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	offsets, _ := convert.ToInts(input)
	steps := 0

	for position := 0; position >= 0 && position < len(offsets); {
		jump := offsets[position]
		offsets[position]++
		position += jump
		steps++
	}

	return steps
}

func part2(input []string) interface{} {
	offsets, _ := convert.ToInts(input)
	steps := 0

	for position := 0; position >= 0 && position < len(offsets); {
		jump := offsets[position]
		if jump >= 3 {
			offsets[position]--
		} else {
			offsets[position]++
		}
		position += jump
		steps++
	}

	return steps
}
