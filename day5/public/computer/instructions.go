package computer

type Instruction struct {
	Intcode Intcode
	Address int
}

type InstructionResult struct {
	NextAddress int
}

type InstructionHandler func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult)

var InstructionHandlers = map[OpCode]InstructionHandler{
	Add: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		program[writeAddress] = value1 + value2
		return NewInstructionResult(jumpAddress(instruction))
	},
	Multiply: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		program[writeAddress] = value1 * value2
		return NewInstructionResult(jumpAddress(instruction))
	},
	Input: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		address := program[instruction.Address+1]
		program[address] = <-input
		return NewInstructionResult(jumpAddress(instruction))
	},
	Output: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		outputValue := getValue(program, instruction, 0)
		output <- outputValue
		return NewInstructionResult(jumpAddress(instruction))
	},
	JumpIfTrue: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		value := getValue(program, instruction, 0)
		if value != 0 {
			writeValue := getValue(program, instruction, 1)
			return NewInstructionResult(writeValue)
		}
		return NewInstructionResult(jumpAddress(instruction))
	},
	JumpIfFalse: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		value := getValue(program, instruction, 0)
		if value == 0 {
			writeValue := getValue(program, instruction, 1)
			return NewInstructionResult(writeValue)
		}
		return NewInstructionResult(jumpAddress(instruction))
	},
	LessThan: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		writeValue := 0
		if value1 < value2 {
			writeValue = 1
		}
		program[writeAddress] = writeValue
		return NewInstructionResult(jumpAddress(instruction))
	},
	Equals: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program[instruction.Address+3]
		writeValue := 0
		if value1 == value2 {
			writeValue = 1
		}
		program[writeAddress] = writeValue
		return NewInstructionResult(jumpAddress(instruction))
	},
	Exit: func(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
		return NewInstructionResult(jumpAddress(instruction))
	},
}

func NewInstructionResult(nextAddress int) *InstructionResult {
	return &InstructionResult{nextAddress}
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
