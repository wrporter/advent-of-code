package computer

import (
	"sync"
)

type Computer struct{}

func New() *Computer {
	return &Computer{}
}

func (c *Computer) Run(program []int, input chan int, output chan int) {
	address := 0
	instruction := ParseInstruction(program, address)

	for instruction.Intcode.OpCode != Exit {
		result := execute(program, instruction, input, output)
		address = result.NextAddress

		instruction = ParseInstruction(program, address)
	}
	close(output)
}

func (c *Computer) Thread(wg *sync.WaitGroup, program []int, input chan int, output chan int) {
	c.Run(program, input, output)
	wg.Done()
}

func execute(program []int, instruction Instruction, input chan int, output chan int) (result *InstructionResult) {
	return InstructionHandlers[instruction.Intcode.OpCode](program, instruction, input, output)
}
