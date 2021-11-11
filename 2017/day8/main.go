package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 8
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^([a-z]+)\s(inc|dec)\s(-?\d+)\sif\s([a-z]+)\s(<|>|<=|>=|!=|==)\s(-?\d+)$`)

type Operation string

var Increment Operation = "inc"
var Decrement Operation = "dec"

type Instruction struct {
	Register          string
	Operation         Operation
	Modifier          int
	ConditionRegister string
	Condition         string
	ConditionValue    int
}

func part1(input []string) interface{} {
	instructions, registers := parseInstructions(input)

	for _, instruction := range instructions {
		meetsCondition := false
		switch instruction.Condition {
		case "<":
			meetsCondition = registers[instruction.ConditionRegister] < instruction.ConditionValue
		case ">":
			meetsCondition = registers[instruction.ConditionRegister] > instruction.ConditionValue
		case "<=":
			meetsCondition = registers[instruction.ConditionRegister] <= instruction.ConditionValue
		case ">=":
			meetsCondition = registers[instruction.ConditionRegister] >= instruction.ConditionValue
		case "==":
			meetsCondition = registers[instruction.ConditionRegister] == instruction.ConditionValue
		case "!=":
			meetsCondition = registers[instruction.ConditionRegister] != instruction.ConditionValue
		}

		if meetsCondition {
			switch instruction.Operation {
			case Increment:
				registers[instruction.Register] += instruction.Modifier
			case Decrement:
				registers[instruction.Register] -= instruction.Modifier
			}
		}
	}

	maxRegisterValue := 0
	for _, value := range registers {
		if value > maxRegisterValue {
			maxRegisterValue = value
		}
	}

	return maxRegisterValue
}

func parseInstructions(input []string) (instructions []Instruction, registers map[string]int) {
	registers = make(map[string]int)

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		instruction := Instruction{
			Register:          match[1],
			Operation:         Operation(match[2]),
			Modifier:          conversion.StringToInt(match[3]),
			ConditionRegister: match[4],
			Condition:         match[5],
			ConditionValue:    conversion.StringToInt(match[6]),
		}
		instructions = append(instructions, instruction)

		registers[instruction.Register] = 0
		registers[instruction.ConditionRegister] = 0
	}
	return instructions, registers
}

func part2(input []string) interface{} {
	instructions, registers := parseInstructions(input)

	maxRegisterValue := 0

	for _, instruction := range instructions {
		meetsCondition := false
		switch instruction.Condition {
		case "<":
			meetsCondition = registers[instruction.ConditionRegister] < instruction.ConditionValue
		case ">":
			meetsCondition = registers[instruction.ConditionRegister] > instruction.ConditionValue
		case "<=":
			meetsCondition = registers[instruction.ConditionRegister] <= instruction.ConditionValue
		case ">=":
			meetsCondition = registers[instruction.ConditionRegister] >= instruction.ConditionValue
		case "==":
			meetsCondition = registers[instruction.ConditionRegister] == instruction.ConditionValue
		case "!=":
			meetsCondition = registers[instruction.ConditionRegister] != instruction.ConditionValue
		}

		if meetsCondition {
			switch instruction.Operation {
			case Increment:
				registers[instruction.Register] += instruction.Modifier
			case Decrement:
				registers[instruction.Register] -= instruction.Modifier
			}

			if registers[instruction.Register] > maxRegisterValue {
				maxRegisterValue = registers[instruction.Register]
			}
		}
	}

	return maxRegisterValue
}
