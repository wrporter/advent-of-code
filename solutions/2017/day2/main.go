package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 2
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	checksum := 0

	for _, line := range input {
		values, _ := convert.ToInts(strings.Fields(line))
		min := values[0]
		max := values[0]

		for _, value := range values {
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}

		difference := ints.Abs(min - max)
		checksum += difference
	}

	return checksum
}

func part2(input []string) interface{} {
	checksum := 0

	for _, line := range input {
		values, _ := convert.ToInts(strings.Fields(line))

		checksumFound := false
		for i, value := range values {

			for j, value2 := range values {
				if i == j {
					continue
				}

				if value%value2 == 0 {
					checksum += value / value2
					checksumFound = true
					break
				}
			}

			if checksumFound {
				break
			}
		}
	}

	return checksum
}
