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

	year, day := 2020, 13
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	timestamp := convert.StringToInt(input[0])
	buses, _ := convert.ToInts(strings.Split(strings.ReplaceAll(input[1], ",x", ""), ","))

	var minBus int
	soonestTime := ints.MaxInt
	for _, bus := range buses {
		remainder := timestamp % bus

		if (bus - remainder) < soonestTime {
			soonestTime = bus - remainder
			minBus = bus
		}
	}
	return minBus * soonestTime
}

func part2(input []string) interface{} {
	busIDs := strings.Split(input[1], ",")

	var n, moduli []int
	for i, busIDString := range busIDs {
		if busIDString != "x" {
			busID := convert.StringToInt(busIDString)
			n = append(n, busID-i)
			moduli = append(moduli, busID)
		}
	}

	return ints.ChineseRemainderTheorem(n, moduli)
}
