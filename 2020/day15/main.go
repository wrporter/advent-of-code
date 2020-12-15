package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
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
	defer timeit.Track(time.Now(), "Part 2")
	startNumbers := parseNumbers(input)
	target := 30000000

	numbers := make([]int, target)
	said := make(map[int]*State)
	for i, number := range startNumbers {
		numbers[i] = number
		said[number] = &State{
			Last:       i,
			BeforeLast: -1,
		}
	}

	start := len(startNumbers)
	for turn := start; turn < target; turn++ {
		prev := numbers[turn-1]

		var current int
		if state, ok := said[prev]; ok && state.BeforeLast == -1 {
			current = 0
		} else {
			current = state.Last - state.BeforeLast
		}

		if state, ok := said[current]; ok {
			state.BeforeLast = state.Last
			state.Last = turn
		} else {
			said[current] = &State{
				Last:       turn,
				BeforeLast: -1,
			}
		}

		numbers[turn] = current
	}

	return numbers[target-1]
}

type State struct {
	Last       int
	BeforeLast int
}

func parseNumbers(input []string) []int {
	values := strings.Split(input[0], ",")
	numbers := make([]int, len(values))
	for i, value := range values {
		numbers[i] = conversion.StringToInt(value)
	}
	return numbers
}
