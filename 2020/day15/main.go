package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

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

func playMemoryGame(startNumbers []int, numTurns int) int {
	history := make([]int, numTurns)
	last := startNumbers[0]

	for turn := range startNumbers {
		history[last] = turn
		last = startNumbers[turn]
	}

	for turn := len(startNumbers); turn < numTurns; turn++ {
		previous := history[last]
		history[last] = turn

		if previous == 0 {
			last = 0
		} else {
			last = turn - previous
		}
	}

	return last
}

func parseNumbers(input []string) []int {
	values := strings.Split(input[0], ",")
	numbers, _ := convert.ToInts(values)
	return numbers
}
