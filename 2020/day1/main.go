package main

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/probability"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	out.Day(2020, 1)
	input, _ := file.ReadFile("./2020/day1/input.txt")
	entries, _ := convert.ToInts(input)
	answer1b := part1b(entries)
	out.Part1(answer1b)
	answer2 := part2(entries)
	out.Part2(answer2)
}

func part1(entries []int) int {
	for i, value1 := range entries {
		for j, value2 := range entries {
			if i == j {
				continue
			}

			if value1+value2 == 2020 {
				return value1 * value2
			}
		}
	}
	return -1
}

func part1b(entries []int) int {
	result := -1
	probability.ComboSize(entries, 2, 2, func(ints []int) {
		if ints[0]+ints[1] == 2020 {
			result = ints[0] * ints[1]
		}
	})
	return result
}

func part2(entries []int) int {
	for i, value1 := range entries {
		for j, value2 := range entries {
			for k, value3 := range entries {
				if i == j || i == k || j == k {
					continue
				}

				if (value1 + value2 + value3) == 2020 {
					return value1 * value2 * value3
				}
			}
		}
	}
	return -1
}
