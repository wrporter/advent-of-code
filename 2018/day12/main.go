package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 12
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	state, low, high, spread := parseInput(input)
	numGenerations := 20
	min := low
	max := high

	for generation := 1; generation <= numGenerations; generation++ {
		next := make(map[int]bool)

		for pot := low - 2; pot <= high+2; pot++ {
			from := getPotRange(state, pot-2, pot+2)
			if to, ok := spread[from]; ok && to == plant {
				next[pot] = true
				min = ints.Min(min, pot)
				max = ints.Max(max, pot)
			}
		}

		state = next
		low = min
		high = max
	}

	return sumPlantedPots(state)
}

func part2(input []string) interface{} {
	state, low, high, spread := parseInput(input)
	numGenerations := 2_000
	min := low
	max := high

	var diffs []int
	prevSum := sumPlantedPots(state)

	for generation := 1; generation <= numGenerations; generation++ {
		next := make(map[int]bool)

		for pot := low - 2; pot <= high+2; pot++ {
			from := getPotRange(state, pot-2, pot+2)
			if to, ok := spread[from]; ok && to == plant {
				next[pot] = true
				min = ints.Min(min, pot)
				max = ints.Max(max, pot)
			}
		}

		state = next
		low = min
		high = max

		curSum := sumPlantedPots(state)
		diff := curSum - prevSum
		diffs = append(diffs, diff)
		if len(diffs) > 100 {
			diffs = diffs[1:]
		}
		prevSum = curSum
	}

	// assume a linear progression
	last100diffs := ints.Sum(diffs) / len(diffs)
	sum := (50_000_000_000-numGenerations)*last100diffs + sumPlantedPots(state)

	return sum
}

func parseInput(input []string) (map[int]bool, int, int, map[string]rune) {
	state := make(map[int]bool)
	initialState := strings.Split(input[0], "initial state: ")[1]
	for pot, potState := range initialState {
		if potState == plant {
			state[pot] = true
		}
	}
	low := 0
	high := len(initialState) - 1

	spread := make(map[string]rune)
	for _, line := range input[2:] {
		parts := strings.Split(line, " => ")
		spread[parts[0]] = rune(parts[1][0])
	}
	return state, low, high, spread
}

func sumPlantedPots(state map[int]bool) int {
	sum := 0
	for pot := range state {
		sum += pot
	}
	return sum
}

func getPotRange(state map[int]bool, low, high int) string {
	b := strings.Builder{}
	for i := low; i <= high; i++ {
		if state[i] {
			b.WriteRune(plant)
		} else {
			b.WriteRune(empty)
		}
	}
	return b.String()
}

var (
	plant = '#'
	empty = '.'
)
