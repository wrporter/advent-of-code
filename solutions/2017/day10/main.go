package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/knot"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 10
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	lengths, _ := convert.ToInts(strings.Split(input[0], ","))
	return singleRoundHash(256, lengths)
}

func singleRoundHash(size int, lengths []int) int {
	var list []int
	for value := 0; value < size; value++ {
		list = append(list, value)
	}

	position := 0
	skip := 0

	for _, length := range lengths {
		list = knot.Reverse(list, position, length)

		position = (position + length + skip) % size
		skip++
	}

	return list[0] * list[1]
}

func part2(input []string) interface{} {
	result := knot.Hash(input[0])
	return result
}
