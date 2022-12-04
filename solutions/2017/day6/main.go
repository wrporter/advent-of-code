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

	year, day := 2017, 6
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	banks, _ := convert.ToInts(strings.Fields(input[0]))
	seen := make(map[string]bool)
	cycles := 0

	for !seen[fmt.Sprint(banks)] || cycles == 0 {
		seen[fmt.Sprint(banks)] = true

		max := banks[0]
		maxIndex := 0
		for i, bank := range banks {
			if bank > max {
				max = bank
				maxIndex = i
			}
		}
		banks[maxIndex] = 0

		for i := (maxIndex + 1) % len(banks); max > 0; max-- {
			banks[i%len(banks)]++
			i++
		}

		cycles++
	}

	return cycles
}

func part2(input []string) interface{} {
	banks, _ := convert.ToInts(strings.Fields(input[0]))
	seen := make(map[string]int)
	cycles := 0
	loopSize := 0

	for {
		bankString := fmt.Sprint(banks)

		if cycle, ok := seen[bankString]; ok {
			loopSize = cycles - cycle
			break
		}

		seen[bankString] = cycles

		max := banks[0]
		maxIndex := 0
		for i, bank := range banks {
			if bank > max {
				max = bank
				maxIndex = i
			}
		}
		banks[maxIndex] = 0

		for i := (maxIndex + 1) % len(banks); max > 0; max-- {
			banks[i%len(banks)]++
			i++
		}

		cycles++
	}

	return loopSize
}
