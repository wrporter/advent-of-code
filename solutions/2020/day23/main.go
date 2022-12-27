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

	year, day := 2020, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	result := play(input[0], 100, 9, 8)
	return ints.Join(result, "")
}

func part2(input []string) interface{} {
	result := play(input[0], 10_000_000, 1_000_000, 2)
	return ints.Product(result)
}

func play(input string, moves int, size int, take int) []int {
	start, _ := convert.ToInts(strings.Split(input, ""))
	ring := make([]int, size+1)
	for i := range ring {
		ring[i] = i + 1
	}

	ring[len(ring)-1] = start[0]

	for i, num := range start {
		if i < len(start)-1 {
			ring[num] = start[i+1]
		} else if len(start) < len(ring)-1 {
			ring[num] = len(start) + 1
		} else {
			ring[num] = start[0]
		}
	}

	var pickup [3]int
	current := start[0]

	for move := 1; move <= moves; move++ {
		pickup[0] = ring[current]
		pickup[1] = ring[pickup[0]]
		pickup[2] = ring[pickup[1]]
		ring[current] = ring[pickup[2]]

		destination := getDestination(pickup, current, size+1)
		end := ring[destination]
		ring[destination] = pickup[0]
		ring[pickup[2]] = end

		current = ring[current]
	}

	result := []int{ring[1]}
	for taken := 1; taken < take; taken++ {
		result = append(result, ring[result[len(result)-1]])
	}
	return result
}

func getDestination(pickup [3]int, current int, size int) int {
	destination := 0
	for next := current - 1; pickup[0] == destination ||
		pickup[1] == destination ||
		pickup[2] == destination ||
		destination == 0; next-- {
		destination = ints.WrapMod(next, size)
	}
	return destination
}
