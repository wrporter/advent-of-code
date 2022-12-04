package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

type OpCode int

const (
	Add      OpCode = 1
	Multiply OpCode = 2
	Exit     OpCode = 99
)

type Instruction struct {
	OpCode     OpCode
	Parameters []int
}

func main() {
	codeLines, _ := file.ReadFile("./2019/day2/input.txt")
	program, _ := convert.ToInts(strings.Split(codeLines[0], ","))
	//program := []int{1,9,10,3,2,3,11,0,99,30,40,50}
	fmt.Println(SolvePart1(copyArray(program)))
	fmt.Println(SolvePart2(copyArray(program), 19690720))
}

func SolvePart2(program []int, target int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			candidate := copyArray(program)
			load(candidate, noun, verb)
			Execute(candidate)
			if candidate[0] == target {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func SolvePart1(program []int) []int {
	load(program, 12, 2)
	return Execute(program)
}

func Execute(program []int) []int {
	index := 0
	opcode := program[index]

	for opcode != 99 {
		op := parseInstruction(program, index)
		execute(program, op)

		index += 4
		opcode = program[index]
	}

	return program
}

func load(program []int, noun int, verb int) {
	program[1] = noun
	program[2] = verb
}

func execute(program []int, instruction Instruction) {
	result := 0
	value1 := program[instruction.Parameters[0]]
	value2 := program[instruction.Parameters[1]]

	if instruction.OpCode == Add {
		result = value1 + value2
	} else if instruction.OpCode == Multiply {
		result = value1 * value2
	}

	program[instruction.Parameters[2]] = result
}

func parseInstruction(program []int, index int) Instruction {
	return Instruction{
		OpCode(program[index]),
		[]int{
			program[index+1],
			program[index+2],
			program[index+3],
		},
	}
}

func copyArray(array []int) []int {
	cpy := make([]int, len(array))
	copy(cpy, array)
	return cpy
}
