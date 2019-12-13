package computer

type Computer struct{}

func New() *Computer {
	return &Computer{}
}

type Program struct {
	Memory          []int
	Input           chan int
	Output          chan int
	address         int
	relativeAddress int
}

func NewProgram(code []int) *Program {
	return &Program{
		Memory:          addMemory(code),
		Input:           make(chan int),
		Output:          make(chan int),
		address:         0,
		relativeAddress: 0,
	}
}

func (c *Computer) Run(program *Program) {
	go c.executeProgram(program)
}

func (c *Computer) executeProgram(program *Program) {
	instruction := ParseInstruction(program)

	for instruction.Intcode.OpCode != Exit {
		result := execute(program, instruction)
		program.address = result.NextAddress
		instruction = ParseInstruction(program)
	}
	close(program.Output)
}

func addMemory(code []int) []int {
	newMemory := make([]int, len(code)<<20)
	copy(newMemory, code)
	return newMemory
}

func execute(program *Program, instruction Instruction) (result *InstructionResult) {
	return InstructionHandlers[instruction.Intcode.OpCode](program, instruction)
}
