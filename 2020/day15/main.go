package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"strings"
)

func main() {
	year, day := 2020, 15
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numbers := parseNumbers(input)
	target := 2020

	said := make(map[int]int)
	for _, number := range numbers {
		said[number] = 1
	}

	start := len(numbers)
	for i := start; i < target; i++ {
		prev := numbers[i-1]
		var current int
		if said[prev] == 1 {
			current = 0
		} else {
			last := -1
			for j := i - 2; j >= 0 && last == -1; j-- {
				if numbers[j] == prev {
					last = j
				}
			}
			current = i - (last + 1)
		}
		numbers = append(numbers, current)
		said[current]++
	}

	return numbers[target-1]
}

func part2(input []string) interface{} {
	numbers := parseNumbers(input)
	target := 30000000

	said := make(map[int][]int)
	for i, number := range numbers {
		said[number] = []int{i}
	}

	start := len(numbers)
	for turn := start; turn < target; turn++ {
		prev := numbers[turn-1]
		var current int
		if len(said[prev]) == 1 {
			current = 0
		} else {
			last := said[prev][len(said[prev])-2]
			current = turn - (last + 1)
		}
		numbers = append(numbers, current)
		said[current] = append(said[current], turn)
	}

	return numbers[target-1]
}

func parseNumbers(input []string) []int {
	values := strings.Split(input[0], ",")
	numbers := make([]int, len(values))
	for i, value := range values {
		numbers[i] = conversion.StringToInt(value)
	}
	return numbers
}
