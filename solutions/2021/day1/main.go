package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 1
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	depths, _ := convert.ToInts(input)

	numIncreases := 0
	lastDepth := 0

	for i, depth := range depths {
		if i > 0 && depth > lastDepth {
			numIncreases++
		}
		lastDepth = depth
	}

	return numIncreases
}

func part2(input []string) interface{} {
	depths, _ := convert.ToInts(input)

	var windows []int
	numIncreases := 0
	lastWindowIndex := 0

	for i := range depths {
		if i >= 2 {
			windows = append(windows, depths[i]+depths[i-1]+depths[i-2])
			lastWindowIndex = len(windows) - 1
		}
		if lastWindowIndex > 0 && windows[lastWindowIndex] > windows[lastWindowIndex-1] {
			numIncreases++
		}
	}

	return numIncreases
}
