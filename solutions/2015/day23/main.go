package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"regexp"
	"strings"
)

func main() {
	year, day := 2015, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^(?:(hlf|tpl|inc) ([ab])|(jmp) ([+-])(\d+)|(jie|jio) ([ab]), ([+-])(\d+))?$`)

type Instruction struct {
	Operation string
	Register  string
	Offset    int
}

func part1(input []string) interface{} {
	instructions := parseInstructions(input)

	registers := map[string]int{
		"a": 0,
		"b": 0,
	}
	run(instructions, registers)

	return registers["b"]
}

func part2(input []string) interface{} {
	instructions := parseInstructions(input)

	registers := map[string]int{
		"a": 1,
		"b": 0,
	}
	run(instructions, registers)

	return registers["b"]
}

func run(instructions []Instruction, registers map[string]int) {
	for i := 0; i < len(instructions) && i >= 0; {
		instruction := instructions[i]
		switch instruction.Operation {
		case "hlf":
			registers[instruction.Register] /= 2
			i++
		case "tpl":
			registers[instruction.Register] *= 3
			i++
		case "inc":
			registers[instruction.Register] += 1
			i++
		case "jmp":
			i += instruction.Offset
		case "jie":
			if isEven(registers[instruction.Register]) {
				i += instruction.Offset
			} else {
				i++
			}
		case "jio":
			if registers[instruction.Register] == 1 {
				i += instruction.Offset
			} else {
				i++
			}
		}
	}
}

func parseInstructions(input []string) []Instruction {
	instructions := make([]Instruction, len(input))
	for i, line := range input {
		parts := strings.Split(line, " ")
		operation := parts[0]
		register := ""
		offset := 0
		sign := "+"

		switch operation {
		case "hlf":
			fallthrough
		case "tpl":
			fallthrough
		case "inc":
			register = parts[1]
		case "jmp":
			sign = string(parts[1][0])
			offset = convert.StringToInt(parts[1][1:])
			if sign == "-" {
				offset = -offset
			}
		case "jie":
			fallthrough
		case "jio":
			register = parts[1][:len(parts[1])-1]
			sign = string(parts[2][0])
			offset = convert.StringToInt(parts[2][1:])
			if sign == "-" {
				offset = -offset
			}
		}

		instructions[i] = Instruction{
			Operation: operation,
			Register:  register,
			Offset:    offset,
		}
	}
	return instructions
}

func isEven(value int) bool {
	return (value % 2) == 0
}
