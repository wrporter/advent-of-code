package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/asm"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 25
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	c := asm.NewComputer(input)
	c.Operations = getOperations(c)
	for a := 0; a <= 1_000_000; a++ {
		c.Registers = map[string]int{"a": a}
		c.Output = []interface{}{}
		c.Address = 0
		c.Run()

		isCorrectSignal := true
		for i := 0; i < len(c.Output); i += 2 {
			if c.Output[i] != 0 || c.Output[i+1] != 1 {
				isCorrectSignal = false
				break
			}
		}
		if isCorrectSignal {
			return a
		}
	}
	return -1
}

func part2(input []string) interface{} {
	return "Merry Christmas!!!"
}

func getOperations(c *asm.Computer) map[string]func(args []string) (jumped bool) {
	return map[string]func(args []string) (jumped bool){
		"cpy": func(args []string) (jumped bool) {
			c.Registers[args[1]] = c.GetValue(args[0])
			return false
		},
		"inc": func(args []string) (jumped bool) {
			c.Registers[args[0]] += 1
			return false
		},
		"dec": func(args []string) (jumped bool) {
			c.Registers[args[0]] -= 1
			return false
		},
		"jnz": func(args []string) (jumped bool) {
			if c.GetValue(args[0]) != 0 {
				c.Address += c.GetValue(args[1])
				return true
			}
			return false
		},
		"out": func(args []string) (jumped bool) {
			c.Output = append(c.Output, c.GetValue(args[0]))
			return false
		},
	}
}
