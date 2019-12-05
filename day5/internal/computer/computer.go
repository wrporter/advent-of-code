package computer

type Computer struct{}

func New() *Computer {
	return &Computer{}
}

func (c *Computer) Run(program []int, input int) (output []int) {
	address := 0
	instruction := ParseInstruction(program, address)

	for instruction.Intcode.OpCode.Code != Exit {
		result := execute(program, instruction, input)
		output = append(output, result.Output...)
		address = result.NextAddress

		instruction = ParseInstruction(program, address)
	}

	return output
}

func execute(program []int, instruction Instruction, input int) (result *InstructionResult) {
	return InstructionHandlers[instruction.Intcode.OpCode.Code](program, instruction, input)
}
