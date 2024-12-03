package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"regexp"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0

	for _, match := range matches {
		num1 := convert.StringToInt(match[1])
		num2 := convert.StringToInt(match[2])
		sum += num1 * num2
	}

	return sum
}

func part2(input string, _ ...interface{}) interface{} {
	regex := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	enabled := true

	for _, match := range matches {
		instruction := match[1]
		if strings.HasPrefix(instruction, "don't") {
			enabled = false
		} else if strings.HasPrefix(instruction, "do") {
			enabled = true
		} else if strings.HasPrefix(instruction, "mul") && enabled {
			param1 := convert.StringToInt(match[2])
			param2 := convert.StringToInt(match[3])
			sum += param1 * param2
		}
	}

	return sum
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 3, Part1: part1, Part2: part2}
}
