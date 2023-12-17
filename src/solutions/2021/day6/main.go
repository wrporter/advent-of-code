package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"container/list"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 6
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	fishInput, _ := convert.ToInts(strings.Split(input[0], ","))
	fishes := list.New()
	for _, fish := range fishInput {
		fishes.PushBack(fish)
	}

	numDays := 80
	for day := 0; day < numDays; day++ {
		numNew := 0
		for e := fishes.Front(); e != nil; e = e.Next() {
			if e.Value == 0 {
				numNew++
				e.Value = 6
			} else {
				e.Value = e.Value.(int) - 1
			}
		}
		for n := 0; n < numNew; n++ {
			fishes.PushBack(8)
		}
	}

	return fishes.Len()
}

func part2(input []string) interface{} {
	fishes, _ := convert.ToInts(strings.Split(input[0], ","))
	// keep track of how many fish are in each age bucket
	fishAge := make([]int, 9)
	for _, fish := range fishes {
		fishAge[fish]++
	}

	numDays := 256
	for day := 0; day < numDays; day++ {
		fishAge[(day+7)%9] += fishAge[day%9]
	}

	return ints.Sum(fishAge)
}
