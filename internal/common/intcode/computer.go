package intcode

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

func ReadCode(filePath string) []int {
	lines, _ := file.ReadFile(filePath)
	code, _ := convert.ToInts(strings.Split(lines[0], ","))
	return code
}

type Computer struct{}

func New() *Computer {
	return &Computer{}
}

type Program struct {
	Memory          map[int]int
	Input           chan<- int
	Output          <-chan int
	input           <-chan int
	output          chan<- int
	address         int
	relativeAddress int
}

func NewProgram(code []int) *Program {
	in := make(chan int)
	out := make(chan int)
	return &Program{
		Memory:          createMemory(code),
		Input:           in,
		Output:          out,
		input:           in,
		output:          out,
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
		//program.log(instruction)
		result := execute(program, instruction)
		program.address = result.NextAddress
		instruction = ParseInstruction(program)
	}
	close(program.output)
}

func createMemory(code []int) map[int]int {
	memory := make(map[int]int)
	for i, v := range code {
		memory[i] = v
	}
	return memory
}

func execute(program *Program, instruction Instruction) (result *InstructionResult) {
	return InstructionHandlers[instruction.Intcode.OpCode](program, instruction)
}

func (p *Program) log(instruction Instruction) {
	line := fmt.Sprintf("[%3d] %s", p.address, instruction.Intcode.OpCode)
	for i := 0; i < instruction.Intcode.NumParameters; i++ {
		line += fmt.Sprintf(" %s(%d)", instruction.Intcode.ParameterModes[i], p.Memory[p.address+i+1])
	}
	fmt.Println(line)
}
