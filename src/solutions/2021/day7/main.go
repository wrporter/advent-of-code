package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 7
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	crabs, _ := convert.ToInts(strings.Split(input[0], ","))
	min := ints.Min(crabs...)
	max := ints.Max(crabs...)
	minFuel := ints.MaxInt

	for x := min; x <= max; x++ {
		fuel := 0
		for _, crab := range crabs {
			fuel += ints.Abs(crab - x)
		}
		minFuel = ints.Min(minFuel, fuel)
	}

	return minFuel
}

func part2(input []string) interface{} {
	crabs, _ := convert.ToInts(strings.Split(input[0], ","))
	min := ints.Min(crabs...)
	max := ints.Max(crabs...)
	minFuel := ints.MaxInt

	for x := min; x <= max; x++ {
		fuel := 0
		for _, crab := range crabs {
			n := ints.Abs(crab - x)
			fuel += (n * (n + 1)) / 2
		}
		minFuel = ints.Min(minFuel, fuel)
	}

	return minFuel
}
