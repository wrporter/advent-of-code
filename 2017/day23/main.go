package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	registers := make(map[string]int)

	countMul := 0

	for i := 0; i < len(input); {
		command := strings.Fields(input[i])
		args := command[1:]

		switch command[0] {
		case "set":
			registers[args[0]] = getValue(registers, args[1])
		case "sub":
			registers[args[0]] -= getValue(registers, args[1])
		case "mul":
			registers[args[0]] *= getValue(registers, args[1])
			countMul++
		case "jnz":
			if getValue(registers, args[0]) != 0 {
				i += getValue(registers, args[1])
				continue
			}
		}
		i++
	}

	return countMul
}

func part2(input []string) interface{} {
	parts := strings.Fields(input[0])
	b := conversion.StringToInt(parts[2])
	b *= 100
	b -= -100000

	// Count all non-prime numbers in the range
	h := 0
	for x := b; x <= b+17000; x += 17 {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				h += 1
				break
			}
		}
	}

	return h
}

func getValue(registers map[string]int, registerOrInt string) int {
	if value, err := strconv.Atoi(registerOrInt); err == nil {
		return value
	}
	return registers[registerOrInt]
}
