package asm

import (
	"strconv"
	"strings"
)

type (
	Computer struct {
		Registers    map[string]int
		Address      int
		Instructions []*Instruction
		Operations   map[string]func(args []string) (jumped bool)
	}
	Instruction struct {
		Command string
		Args    []string
	}
)

func NewComputer(code []string) *Computer {
	return &Computer{
		Registers:    make(map[string]int),
		Address:      0,
		Instructions: parseInstructions(code),
	}
}

func (c *Computer) Run() {
	for c.Address >= 0 && c.Address < len(c.Instructions) {
		instruction := c.Instructions[c.Address]
		jumped := c.Operations[instruction.Command](instruction.Args)
		if !jumped {
			c.Address += 1
		}
	}
}

func (c *Computer) GetValue(arg string) int {
	if value, err := strconv.Atoi(arg); err == nil {
		return value
	}
	return c.Registers[arg]
}

func parseInstructions(code []string) []*Instruction {
	var instructions []*Instruction
	for _, line := range code {
		parts := strings.Fields(line)
		instructions = append(instructions, &Instruction{
			Command: parts[0],
			Args:    parts[1:],
		})
	}
	return instructions
}
