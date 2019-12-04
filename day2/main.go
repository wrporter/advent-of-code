package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

type OpType int

const (
	Add      OpType = 1
	Multiply OpType = 2
)

type Op struct {
	Type          OpType
	InputOneIndex int
	InputTwoIndex int
	OutputIndex   int
}

func main() {
	codeLines, _ := file.ReadFile("./day2/input.txt")
	program, _ := conversion.ToInts(strings.Split(codeLines[0], ","))
	//program := []int{1,9,10,3,2,3,11,0,99,30,40,50}
	restoreGravityAssist(program)
	fmt.Println(Solve(program))
}

func Solve(program []int) []int {
	index := 0
	opcode := program[index]

	for opcode != 99 {
		op := parseOp(program, index)
		executeOp(program, op)

		index += 4
		opcode = program[index]
	}

	return program
}

func restoreGravityAssist(program []int) {
	program[1] = 12
	program[2] = 2
}

func executeOp(program []int, op Op) {
	result := 0
	value1 := program[op.InputOneIndex]
	value2 := program[op.InputTwoIndex]

	if op.Type == Add {
		result = value1 + value2
	} else if op.Type == Multiply {
		result = value1 * value2
	}

	program[op.OutputIndex] = result
}

func parseOp(program []int, index int) Op {
	return Op{
		OpType(program[index]),
		program[index+1],
		program[index+2],
		program[index+3],
	}
}
