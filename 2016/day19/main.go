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
		elves.Value = &elf{i, 1}
		elves = elves.Next()
	}

	for e := elves; e.Value != numElves; e = e.Next() {
		if e.Value.(*elf).presents == 0 {
			continue
		}

		numPresents := e.Value.(*elf).presents + e.Next().Value.(*elf).presents
		e.Value.(*elf).presents = numPresents
		e.Unlink(1)

		if numPresents == numElves {
			return e.Value.(*elf).id
		}
	}

	return 0
}

func part2(input []string) interface{} {
	return 0
}

type elf struct {
	id       int
	presents int
}
