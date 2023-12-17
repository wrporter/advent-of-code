package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 8
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	instructions := parseInstructions(input)
	accumulator, _ := run(instructions)
	return accumulator
}

func part2(input []string) interface{} {
	instructions := parseInstructions(input)

	var tmp Operation
	for i, instruction := range instructions {
		if instruction.Operation == Nop || instruction.Operation == Jmp {
			tmp = instruction.Operation
			instructions[i].Operation = opposite(instruction.Operation)

			accumulator, terminated := run(instructions)
			instructions[i].Operation = tmp

			if terminated {
				return accumulator
			}
		}
	}

	return 0
}

func opposite(operation Operation) Operation {
	if operation == Jmp {
		return Nop
	}
	return Jmp
}

func run(instructions []Instruction) (accumulator int, terminated bool) {
	visited := make(map[int]bool)

	for index := 0; index < len(instructions); {
		if visited[index] {
			return accumulator, false
		}
		visited[index] = true

		instruction := instructions[index]
		if instruction.Operation == Acc {
			accumulator += instruction.Argument
			index++
		} else if instruction.Operation == Jmp {
			index += instruction.Argument
		} else if instruction.Operation == Nop {
			index++
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
	Acc Operation = "acc"
	Jmp Operation = "jmp"
	Nop Operation = "nop"
)

var regex = regexp.MustCompile(`^(nop|acc|jmp) ([-+])(\d+)$`)

func parseInstructions(input []string) []Instruction {
	instructions := make([]Instruction, len(input))
	for i, line := range input {
		instructions[i] = parseInstruction(line)
	}
	return instructions
}

func parseInstruction(line string) Instruction {
	match := regex.FindStringSubmatch(line)
	operation := Operation(match[1])
	argument := parseArgument(match[2], match[3])

	instruction := Instruction{operation, argument}
	return instruction
}

func parseArgument(sign string, value string) int {
	argument := convert.StringToInt(value)
	if sign == "-" {
		return -argument
	}
	return argument
}
