package computer

import (
	"sync"
)

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
		address:         0,
		relativeAddress: 0,
		Input:           make(chan int, 2),
		Output:          make(chan int, 2),
	}
}

func (c *Computer) RunProgram(program *Program) {
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

func (c *Computer) Run(memory []int, input chan int, output chan int) {
	program := &Program{memory, input, output, 0, 0}
	c.RunProgram(program)
}

func (c *Computer) Thread(wg *sync.WaitGroup, program []int, input chan int, output chan int) {
	c.Run(program, input, output)
	wg.Done()
}

func (c *Computer) ThreadProgram(wg *sync.WaitGroup, program *Program) {
	c.RunProgram(program)
	wg.Done()
}

func execute(program *Program, instruction Instruction) (result *InstructionResult) {
	return InstructionHandlers[instruction.Intcode.OpCode](program, instruction)
}
