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

	year, day := 2017, 17
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	stepSize := convert.StringToInt(input[0])
	buffer := []int{0}
	iterations := 2017
	position := 0

	for value := 1; value <= iterations; value++ {
		position = ((position + stepSize) % value) + 1
		buffer = insert(buffer, position, value)
	}

	return buffer[position+1]
}

func part2(input []string) interface{} {
	stepSize := convert.StringToInt(input[0])
	iterations := 50_000_000
	position := 0
	valueAfter0 := 0

	for value := 1; value <= iterations; value++ {
		position = ((position + stepSize) % value) + 1
		if position == 1 {
			valueAfter0 = value
		}
	}

	return valueAfter0
}

func insert(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
