package main

import (
	"aoc/src/lib/go/asm"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"time"
	"unicode"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	c := asm.NewComputer(input)
	c.Registers["a"] = 7
	c.Operations = getOperations(c)
	c.Run()
	return c.Registers["a"]
}

func part2(input []string) interface{} {
	c := asm.NewComputer(input)
	optimizeMultipliers(c)
	c.Registers["a"] = 12
	c.Operations = getOperations(c)
	c.Run()
	return c.Registers["a"]
}

func optimizeMultipliers(cpu *asm.Computer) {
	// Optimize multiplication instructions by reducing the following 6 instructions into a multiply instruction,
	// followed by nop instructions to maintain the number of instructions in the code.
	//
	// cpy b c
	// inc a
	// dec c
	// jnz c -2
	// dec d
	// jnz d -5
	//
	// mul b d a
	// jnz 0 0
	// jnz 0 0
	// jnz 0 0
	// jnz 0 0
	// jnz 0 0
	for i := range cpu.Instructions {
		if i+5 >= len(cpu.Instructions) {
			continue
		}

		c := cpu.Instructions[i]
		i1 := cpu.Instructions[i+1]
		d1 := cpu.Instructions[i+2]
		j1 := cpu.Instructions[i+3]
		d2 := cpu.Instructions[i+4]
		j2 := cpu.Instructions[i+5]

		if c.Command == "cpy" &&
			i1.Command == "inc" &&
			d1.Command == "dec" &&
			j1.Command == "jnz" &&
			d2.Command == "dec" &&
			j2.Command == "jnz" &&
			cpu.GetValue(j1.Args[1]) == -2 &&
			cpu.GetValue(j2.Args[1]) == -5 &&
			j1.Args[0] == d1.Args[0] && j2.Args[0] == d2.Args[0] {
			cpu.Instructions[i] = &asm.Instruction{
				Command: "mul",
				Args:    []string{c.Args[0], j2.Args[0], i1.Args[0]},
			}
			cpu.Instructions[i+1] = &asm.Instruction{
				Command: "jnz",
				Args:    []string{"0", "0"},
			}
			cpu.Instructions[i+2] = &asm.Instruction{
				Command: "jnz",
				Args:    []string{"0", "0"},
			}
			cpu.Instructions[i+3] = &asm.Instruction{
				Command: "jnz",
				Args:    []string{"0", "0"},
			}
			cpu.Instructions[i+4] = &asm.Instruction{
				Command: "jnz",
				Args:    []string{"0", "0"},
			}
			cpu.Instructions[i+5] = &asm.Instruction{
				Command: "jnz",
				Args:    []string{"0", "0"},
			}
		}
	}
}

func getOperations(c *asm.Computer) map[string]func(args []string) (jumped bool) {
	return map[string]func(args []string) (jumped bool){
		"cpy": func(args []string) (jumped bool) {
			if !isRegister(args[1]) {
				// invalid instruction
				return false
			}
			c.Registers[args[1]] = c.GetValue(args[0])
			return false
		},
		"inc": func(args []string) (jumped bool) {
			if !isRegister(args[0]) {
				// invalid instruction
				return false
			}
			c.Registers[args[0]] += 1
			return false
		},
		"dec": func(args []string) (jumped bool) {
			if !isRegister(args[0]) {
				// invalid instruction
				return false
			}
			c.Registers[args[0]] -= 1
			return false
		},
		"jnz": func(args []string) (jumped bool) {
			if c.GetValue(args[0]) != 0 {
				c.Address += c.GetValue(args[1])
				return true
			}
			return false
		},
		"tgl": func(args []string) (jumped bool) {
			x := c.GetValue(args[0])
			address := c.Address + x
			if address < 0 || address >= len(c.Instructions) {
				return false
			}

			instruction := c.Instructions[address]
			if len(instruction.Args) == 1 {
				if instruction.Command == "inc" {
					instruction.Command = "dec"
				} else {
					instruction.Command = "inc"
				}
			} else if len(instruction.Args) == 2 {
				if instruction.Command == "jnz" {
					instruction.Command = "cpy"
				} else {
					instruction.Command = "jnz"
				}
			}
			return false
		},
		"mul": func(args []string) (jumped bool) {
			if !isRegister(args[2]) {
				// invalid instruction
				return false
			}
			c.Registers[args[2]] = c.GetValue(args[0]) * c.GetValue(args[1])
			return false
		},
	}
}

func isRegister(arg string) bool {
	return unicode.IsLetter([]rune(arg)[0])
}
