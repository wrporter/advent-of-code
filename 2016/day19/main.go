package main

import (
	"container/ring"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 19
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numElves := conversion.StringToInt(input[0])
	elves := ring.New(numElves)
	for i := 1; i <= numElves; i++ {
		elves.Value = i
		elves = elves.Next()
	}

	for i := 0; i <= numElves; i++ {
		elves.Unlink(1)
		elves = elves.Next()
	}

	return elves.Value
}

func part2(input []string) interface{} {
	numElves := conversion.StringToInt(input[0])
	steal := ring.New(numElves)
	var give *ring.Ring
	for i := 1; i <= numElves; i++ {
		if i == (numElves/2)+1 {
			give = steal
		}
		steal.Value = i
		steal = steal.Next()
	}

	for i := 0; i < numElves; i++ {
		give = give.Prev()
		give.Unlink(1)
		give = give.Next()
		if i%2 == 0 {
			give = give.Next()
		}
		steal = steal.Next()
	}

	return steal.Value
}
