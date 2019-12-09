package computer

import "github.com/wrporter/advent-of-code-2019/internal/common/arrays"

type Instruction struct {
	Intcode Intcode
	Address int
}

type InstructionResult struct {
	NextAddress int
	ReadInput   bool
	Output      []int
}

type InstructionHandler func(program []int, instruction Instruction, input []int) (result *InstructionResult)

var InstructionHandlers = map[OpCode]InstructionHandler{
	Add: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		program[writeAddress] = value1 + value2
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
	Multiply: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		program[writeAddress] = value1 * value2
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
	Input: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		address := program[instruction.Address+1]
		program[address], _ = arrays.Poll(input)
		return NewInstructionResult(jumpAddress(instruction), true, nil)
	},
	Output: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		outputValue := getValue(program, instruction, 0)
		return NewInstructionResult(jumpAddress(instruction), false, []int{outputValue})
	},
	JumpIfTrue: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		value := getValue(program, instruction, 0)
		if value != 0 {
			writeValue := getValue(program, instruction, 1)
			program[instruction.Address] = writeValue
			return NewInstructionResult(writeValue, false, nil)
		}
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
	JumpIfFalse: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		value := getValue(program, instruction, 0)
		if value == 0 {
			writeValue := getValue(program, instruction, 1)
			program[instruction.Address] = writeValue
			return NewInstructionResult(writeValue, false, nil)
		}
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
	LessThan: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		writeValue := 0
		if value1 < value2 {
			writeValue = 1
		}
		program[writeAddress] = writeValue
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
	Equals: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		writeValue := 0
		if value1 == value2 {
			writeValue = 1
		}
		program[writeAddress] = writeValue
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
	Exit: func(program []int, instruction Instruction, input []int) (result *InstructionResult) {
		return NewInstructionResult(jumpAddress(instruction), false, nil)
	},
}

func NewInstructionResult(nextAddress int, readInput bool, output []int) *InstructionResult {
	return &InstructionResult{nextAddress, readInput, output}
}

func jumpAddress(instruction Instruction) int {
	return instruction.Address + instruction.Intcode.NumParameters + 1
}

func getValue(program []int, instruction Instruction, parameter int) int {
	valueAddress := getAddress(program, instruction.Address+parameter+1, instruction.Intcode.ParameterModes[parameter])
	return program[valueAddress]
}

func getAddress(program []int, address int, parameterMode ParameterMode) int {
	if parameterMode == Immediate {
		return address
	} else if parameterMode == Position {
		return program[address]
	}
	return address
}
