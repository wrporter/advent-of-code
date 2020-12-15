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
	startNumbers := parseNumbers(input)
	return playMemoryGame(startNumbers, 2020)
}

func part2(input []string) interface{} {
	startNumbers := parseNumbers(input)
	return playMemoryGame(startNumbers, 30000000)
}

func playMemoryGame(startNumbers []int, numTurns int) interface{} {
	numbers := make([]int, numTurns)
	said := make(map[int]int)
	for i, number := range startNumbers {
		numbers[i] = number
		if i != len(startNumbers)-1 {
			said[number] = i
		}
	}

	for turn := len(startNumbers); turn < numTurns; turn++ {
		prev := numbers[turn-1]

		var next int
		if beforeLast, ok := said[prev]; !ok {
			next = 0
		} else {
			next = turn - 1 - beforeLast
		}

		said[prev] = turn - 1
		numbers[turn] = next
	}

	return numbers[numTurns-1]
}

func parseNumbers(input []string) []int {
	values := strings.Split(input[0], ",")
	numbers := make([]int, len(values))
	for i, value := range values {
		numbers[i] = conversion.StringToInt(value)
	}
	return numbers
}
