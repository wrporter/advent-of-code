package computer

type Instruction struct {
	Intcode Intcode
	Address int
}

type InstructionResult struct {
	NextAddress int
}

type InstructionHandler func(program *Program, instruction Instruction) (result *InstructionResult)

var InstructionHandlers = map[OpCode]InstructionHandler{
	Add: func(program *Program, instruction Instruction) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program.Memory[instruction.Address+3]
		program.Memory[writeAddress] = value1 + value2
		return NewInstructionResult(instruction.nextAddress())
	},
	Multiply: func(program *Program, instruction Instruction) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program.Memory[instruction.Address+3]
		program.Memory[writeAddress] = value1 * value2
		return NewInstructionResult(instruction.nextAddress())
	},
	Input: func(program *Program, instruction Instruction) (result *InstructionResult) {
		//address := program.Memory[instruction.Address+1]
		address := getAddress(program, instruction.Address+1, instruction.Intcode.ParameterModes[0])
		program.Memory[address] = <-program.Input
		return NewInstructionResult(instruction.nextAddress())
	},
	Output: func(program *Program, instruction Instruction) (result *InstructionResult) {
		outputValue := getValue(program, instruction, 0)
		program.Output <- outputValue
		return NewInstructionResult(instruction.nextAddress())
	},
	JumpIfTrue: func(program *Program, instruction Instruction) (result *InstructionResult) {
		value := getValue(program, instruction, 0)
		if value != 0 {
			writeValue := getValue(program, instruction, 1)
			return NewInstructionResult(writeValue)
		}
		return NewInstructionResult(instruction.nextAddress())
	},
	JumpIfFalse: func(program *Program, instruction Instruction) (result *InstructionResult) {
		value := getValue(program, instruction, 0)
		if value == 0 {
			writeValue := getValue(program, instruction, 1)
			return NewInstructionResult(writeValue)
		}
		return NewInstructionResult(instruction.nextAddress())
	},
	LessThan: func(program *Program, instruction Instruction) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		//writeAddress := getAddress(program, instruction.Address + 1, instruction.Intcode.ParameterModes[2])
		writeAddress := program.Memory[instruction.Address+3]
		writeValue := 0
		if value1 < value2 {
			writeValue = 1
		}
		program.Memory[writeAddress] = writeValue
		return NewInstructionResult(instruction.nextAddress())
	},
	Equals: func(program *Program, instruction Instruction) (result *InstructionResult) {
		value1 := getValue(program, instruction, 0)
		value2 := getValue(program, instruction, 1)
		writeAddress := program.Memory[instruction.Address+3]
		writeValue := 0
		if value1 == value2 {
			writeValue = 1
		}
		program.Memory[writeAddress] = writeValue
		return NewInstructionResult(instruction.nextAddress())
	},
	ModifyRelativeBase: func(program *Program, instruction Instruction) (result *InstructionResult) {
		modifier := getValue(program, instruction, 0)
		program.relativeAddress += modifier
		return NewInstructionResult(instruction.nextAddress())
	},
	Exit: func(program *Program, instruction Instruction) (result *InstructionResult) {
		return NewInstructionResult(instruction.nextAddress())
	},
}

func NewInstructionResult(nextAddress int) *InstructionResult {
	return &InstructionResult{nextAddress}
}

func (instruction Instruction) nextAddress() int {
	return instruction.Address + instruction.Intcode.NumParameters + 1
}

func getValue(program *Program, instruction Instruction, parameter int) int {
	address := getAddress(program, instruction.Address+parameter+1, instruction.Intcode.ParameterModes[parameter])
	return program.Memory[address]
}

func getAddress(program *Program, address int, parameterMode ParameterMode) int {
	if parameterMode == Immediate {
		return address
	} else if parameterMode == Position {
		return program.Memory[address]
	} else if parameterMode == Relative {
		return program.relativeAddress + program.Memory[address]
	}
	return address
}
