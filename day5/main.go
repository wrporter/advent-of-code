package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strconv"
	"strings"
)

type Code int
type OpCode struct {
	Code          Code
	NumParameters int
}

const (
	Add      Code = 1
	Multiply Code = 2
	Input    Code = 3
	Output   Code = 4
	Exit     Code = 99
)

var OpCodes = map[Code]OpCode{
	Add:      {Add, 3},
	Multiply: {Multiply, 3},
	Input:    {Input, 1},
	Output:   {Output, 1},
	Exit:     {Exit, 0},
}

type ParameterMode int

const (
	Position  = 0
	Immediate = 1
)

type Instruction struct {
	Intcode Intcode
	Address int
}

type Intcode struct {
	OpCode         OpCode
	ParameterModes []ParameterMode
}

func main() {
	codeLines, _ := file.ReadFile("./day5/input.txt")
	//codeLines := []string{"3,0,4,0,99"}
	//codeLines := []string{"1002,4,3,4,33"}
	//codeLines := []string{"1101,100,-1,4,0"}
	program, _ := conversion.ToInts(strings.Split(codeLines[0], ","))
	input := 1
	fmt.Println(Run(program, input))
}

func Run(program []int, input int) []int {
	address := 0
	intcode := parseIntcode(program[address])

	for intcode.OpCode.Code != Exit {
		instruction := parseInstruction(program, address)
		execute(program, instruction, input)

		address += instruction.Intcode.OpCode.NumParameters + 1
		intcode = parseIntcode(program[address])
	}

	return program
}

func execute(program []int, instruction Instruction, input int) {
	switch instruction.Intcode.OpCode.Code {
	case Add:
		value1 := getValue(program, instruction.Address+1, instruction.Intcode.ParameterModes[0])
		value2 := getValue(program, instruction.Address+2, instruction.Intcode.ParameterModes[1])
		writeAddress := program[instruction.Address+3]
		program[writeAddress] = value1 + value2
	case Multiply:
		value1 := getValue(program, instruction.Address+1, instruction.Intcode.ParameterModes[0])
		value2 := getValue(program, instruction.Address+2, instruction.Intcode.ParameterModes[1])
		writeAddress := program[instruction.Address+3]
		program[writeAddress] = value1 * value2
	case Input:
		address := program[instruction.Address+1]
		program[address] = input
	case Output:
		address := program[instruction.Address+1]
		fmt.Println(program[address])
	case Exit:
	}
}

func parseIntcode(intcode int) Intcode {
	intcodeString := strconv.Itoa(intcode)

	opCodeSpace := getOpCodeSpace(intcodeString)
	code := Code(StringToInt(intcodeString[opCodeSpace:]))
	opCode := OpCodes[code]

	parameterModeCodes := Reverse(intcodeString[:opCodeSpace])
	parameterModes := make([]ParameterMode, 3)
	for i := 0; i < len(parameterModeCodes); i++ {
		parameterModes[i] = ParameterMode(RuneToInt(parameterModeCodes[i]))
	}

	return Intcode{opCode, parameterModes}
}

func getOpCodeSpace(intcode string) int {
	space := len(intcode) - 2
	if space < 0 {
		return 0
	}
	return space
}

func parseInstruction(program []int, index int) Instruction {
	return Instruction{
		parseIntcode(program[index]),
		index,
	}
}

func getValue(program []int, address int, mode ParameterMode) int {
	valueAddress := getAddress(program, address, mode)
	return program[valueAddress]
}

func getAddress(program []int, address int, mode ParameterMode) int {
	if mode == Immediate {
		return address
	} else if mode == Position {
		return program[address]
	}
	return address
}

func RuneToInt(rune uint8) int {
	return int(rune - '0')
}

func StringToInt(value string) int {
	valueInt64, _ := strconv.ParseInt(value, 10, 64)
	return int(valueInt64)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
