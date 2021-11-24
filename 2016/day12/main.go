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

	year, day := 2016, 12
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
	c.Run()
	return c.Registers["a"]
}

func part2(input []string) interface{} {
	c := asm.NewComputer(input)
	c.Registers["c"] = 1
	c.Operations = getOperations(c)
	c.Run()
	return c.Registers["a"]
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
	}
}
