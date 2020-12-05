package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"sort"
)

func main() {
	year, day := 2020, 5
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) int {
	maxSeatID := 0
	for _, space := range input {
		seatID := getSeatID(space)
		maxSeatID = ints.Max(maxSeatID, seatID)
	}
	return maxSeatID
}

func getSeatID(space string) int {
	row := binaryPartitionSpace(space[:7], 128)
	column := binaryPartitionSpace(space[7:], 8)
	return (row * 8) + column
}

func part2(input []string) interface{} {
	seatIDs := make([]int, len(input))

	for i, space := range input {
		seatID := getSeatID(space)
		seatIDs[i] = seatID
	}

	sort.Ints(seatIDs)

	previous := seatIDs[0] - 1
	for _, seatID := range seatIDs {
		if seatID-previous != 1 {
			return seatID - 1
		}
		previous = seatID
	}

	return -1
}

func binaryPartitionSpace(space string, size int) int {
	low := 0
	high := size - 1

	for _, char := range space {
		if char == 'F' || char == 'L' {
			high = ((high - low) / 2) + low
		} else if char == 'B' || char == 'R' {
			low = ((high - low) / 2) + low + 1
		}
	}

	return low
}
