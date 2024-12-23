package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/probability"
	"aoc/src/lib/go/timeit"
	"fmt"
	"sort"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 9
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	values, _ := convert.ToInts(input)
	preambleSize := 25

	for i := preambleSize; i < len(values); i++ {
		value := values[i]
		hasSum := false
		probability.ComboSize(values[i-preambleSize:i], 2, 2, func(ints []int) {
			hasSum = hasSum || ((ints[0] + ints[1]) == value)
		})
		if !hasSum {
			return value
		}
	}

	return 0
}

func part2(input []string) interface{} {
	values, _ := convert.ToInts(input)
	preambleSize := 25

	desiredValue := 0
	index := 0
	for i := preambleSize; i < len(values); i++ {
		value := values[i]
		hasSum := false
		probability.ComboSize(values[i-preambleSize:i], 2, 2, func(ints []int) {
			hasSum = hasSum || ((ints[0] + ints[1]) == value)
		})
		if !hasSum {
			desiredValue = value
			index = i
			break
		}
	}

	for i := 0; i < index; i++ {
		for j := 0; j < i; j++ {
			valueRange := values[j:i]
			if ints.Sum(valueRange) == desiredValue {
				sort.Ints(valueRange)
				return valueRange[0] + valueRange[len(valueRange)-1]
			}
		}
	}

	return 0
}
