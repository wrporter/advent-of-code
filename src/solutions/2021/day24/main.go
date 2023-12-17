package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/mystrings"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 24
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var memo = make(map[memory]*int)

func part1(code []string) interface{} {
	c := newALU(code)
	digits := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	registers := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	modelNumber := c.findModelNumber(digits, registers, 0)
	return reverse(modelNumber)
}

func part2(code []string) interface{} {
	c := newALU(code)
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	registers := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	modelNumber := c.findModelNumber(digits, registers, 0)
	return reverse(modelNumber)
}

func (c *alu) findModelNumber(digits []int, registers map[string]int, address int) *int {
	if number, ok := memo[toMemory(registers, address)]; ok {
		return number
	}

	for _, digit := range digits {
		regs := copyMap(registers)
		pc := address
		c.exec(regs, pc, digit)
		pc += 1

		for pc < len(c.instructions) {
			i := c.instructions[pc]
			if i.command == "inp" {
				number := c.findModelNumber(digits, regs, pc)
				if number != nil {
					val := (*number)*10 + digit
					memo[toMemory(regs, pc)] = &val
					return &val
				} else {
					break
				}
			} else {
				c.exec(regs, pc, 0)
				pc += 1
			}
		}

		if pc >= len(c.instructions) && regs["z"] == 0 {
			memo[toMemory(regs, pc)] = &digit
			return &digit
		}
	}

	memo[toMemory(registers, address)] = nil
	return nil
}

func toMemory(registers map[string]int, address int) memory {
	return memory{
		//w:       registers["w"],
		//x:       registers["x"],
		//y:       registers["y"],
		z:       registers["z"],
		address: address,
	}
}

func (c *alu) exec(registers map[string]int, address int, input int) {
	i := c.instructions[address]
	args := i.args

	switch i.command {
	case "inp":
		registers[args[0]] = input
	case "add":
		registers[args[0]] += getValue(registers, args[1])
	case "mul":
		registers[args[0]] *= getValue(registers, args[1])
	case "div":
		registers[args[0]] /= getValue(registers, args[1])
	case "mod":
		registers[args[0]] %= getValue(registers, args[1])
	case "eql":
		if getValue(registers, args[0]) == getValue(registers, args[1]) {
			registers[args[0]] = 1
		} else {
			registers[args[0]] = 0
		}
	}
}

func getValue(registers map[string]int, arg string) int {
	if value, err := strconv.Atoi(arg); err == nil {
		return value
	}
	return registers[arg]
}

func parseInstructions(code []string) []instruction {
	var instructions []instruction
	for _, line := range code {
		parts := strings.Fields(line)
		instructions = append(instructions, instruction{
			command: parts[0],
			args:    parts[1:],
		})
	}
	return instructions
}

type (
	alu struct {
		instructions []instruction
	}
	instruction struct {
		command string
		args    []string
	}
	memory struct {
		//w       int
		//x       int
		//y       int
		z       int
		address int
	}
)

func newALU(code []string) *alu {
	return &alu{
		instructions: parseInstructions(code),
	}
}

func copyMap(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}

func reverse(modelNumber *int) interface{} {
	strVal := strconv.Itoa(*modelNumber)
	strVal = mystrings.Reverse(strVal)
	answer := convert.StringToInt(strVal)

	return answer
}
