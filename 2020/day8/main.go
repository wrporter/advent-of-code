package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"regexp"
)

func main() {
	year, day := 2020, 8
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	instructions := parseInstructions(input)

	visited := make(map[int]bool)
	accumulator := 0
	for index := 0; ; index++ {
		if visited[index] {
			break
		}
		visited[index] = true

		instruction := instructions[index]
		if instruction.Operation == acc {
			accumulator += instruction.Argument
		} else if instruction.Operation == jmp {
			index += instruction.Argument - 1
		} else if instruction.Operation == nop {
			// do nothing
		}
	}

	return accumulator
}

func part2(input []string) interface{} {
	instructions := parseInstructions(input)

	var tmp Operation
	for i, instruction := range instructions {
		if instruction.Operation == nop || instruction.Operation == jmp {
			tmp = instruction.Operation
			instructions[i].Operation = opposite(instruction.Operation)

			accumulator, terminated := tryInstructionSet(instructions)
			instructions[i].Operation = tmp
			if terminated {
				return accumulator
			}
		}
	}

	return 0
}

func opposite(operation Operation) Operation {
	if operation == jmp {
		return nop
	}
	return jmp
}

func tryInstructionSet(instructions []Instruction) (int, bool) {
	visited := make(map[int]bool)
	accumulator := 0

	for index := 0; index < len(instructions); index++ {
		if visited[index] {
			return accumulator, false
		}
		visited[index] = true

		instruction := instructions[index]
		if instruction.Operation == acc {
			accumulator += instruction.Argument
		} else if instruction.Operation == jmp {
			index += instruction.Argument - 1
		} else if instruction.Operation == nop {
			// do nothing
		}
	}

	return accumulator, true
}

type (
	Operation   string
	Instruction struct {
		Operation Operation
		Argument  int
	}
)

const (
	acc Operation = "acc"
	jmp Operation = "jmp"
	nop Operation = "nop"
)

var regex = regexp.MustCompile(`^(nop|acc|jmp) (-|\+)(\d+)$`)

func parseInstructions(input []string) []Instruction {
	var instructions []Instruction
	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		operation := Operation(match[1])
		sign := match[2]
		argument := conversion.StringToInt(match[3])
		if sign == "-" {
			argument = -argument
		}

		instructions = append(instructions, Instruction{operation, argument})
	}
	return instructions
}
