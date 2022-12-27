package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/asm"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
	"unicode"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 25
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	c := asm.NewComputer(input)
	c.Operations = getOperations(c)
	optimize(c)

	for a := 0; a <= 1000; a++ {
		c.Registers = map[string]int{"a": a}
		c.Output = []interface{}{}
		c.Address = 0
		c.Run()

		isCorrectSignal := true
		for i := 0; i < len(c.Output); i += 2 {
			if c.Output[i] != 0 || c.Output[i+1] != 1 {
				isCorrectSignal = false
				break
			}
		}
		if isCorrectSignal {
			return a
		}
	}

	return -1
}

func part2(input []string) interface{} {
	return "Merry Christmas!!!"
}

func optimize(cpu *asm.Computer) {
	// Optimize by reducing simplified loop instructions to more complex instructions.

	for i := range cpu.Instructions {
		// mli
		//
		// cpy a b
		// inc c
		// dec b
		// jnz b -2
		// dec d
		// jnz d -5
		//
		// c += a * d
		// b = 0
		// d = 0
		//
		// mli a b c d
		if in := getNext(cpu.Instructions, i, 6); in != nil &&
			in[0].Command == "cpy" &&
			in[1].Command == "inc" &&
			in[2].Command == "dec" &&
			in[3].Command == "jnz" &&
			in[4].Command == "dec" &&
			in[5].Command == "jnz" &&
			cpu.GetValue(in[3].Args[1]) == -2 &&
			cpu.GetValue(in[5].Args[1]) == -5 &&
			in[3].Args[0] == in[2].Args[0] && in[3].Args[0] == in[4].Args[0] {
			cpu.Instructions[i] = &asm.Instruction{
				Command: "mli",
				Args:    []string{in[0].Args[0], in[0].Args[1], in[1].Args[0], in[4].Args[0]},
			}
		}

		// dvm
		//
		// cpy x r
		// jnz x 2
		// jnz 1 6
		// dec x
		// dec r
		// jnz r -4
		// inc q
		// jnz 1 -7
		//
		// q = x / d
		// r = d - (x % d)
		// x = 0
		//
		// dvm x d q r
		if in := getNext(cpu.Instructions, i, 8); in != nil &&
			in[0].Command == "cpy" &&
			in[1].Command == "jnz" &&
			in[2].Command == "jnz" &&
			in[3].Command == "dec" &&
			in[4].Command == "dec" &&
			in[5].Command == "jnz" &&
			in[6].Command == "inc" &&
			in[7].Command == "jnz" &&
			in[2].Args[0] == "1" && in[2].Args[1] == "6" &&
			in[7].Args[0] == "1" && in[7].Args[1] == "-7" &&
			in[1].Args[0] == in[3].Args[0] && in[1].Args[1] == "2" &&
			in[5].Args[0] == in[4].Args[0] && in[5].Args[1] == "-4" &&
			in[0].Args[1] == in[4].Args[0] {
			instruction := &asm.Instruction{
				Command: "dvm",
				Args:    []string{in[3].Args[0], in[0].Args[0], in[6].Args[0], in[4].Args[0]},
			}
			cpu.Instructions[i] = instruction
		}
	}
}

func getNext(instructions []*asm.Instruction, index int, amount int) []*asm.Instruction {
	if index+amount > len(instructions) {
		return nil
	}
	return instructions[index : index+amount]
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
			c.Registers[args[2]] += c.GetValue(args[0]) * c.GetValue(args[3])
			c.Registers[args[1]] = 0
			c.Registers[args[3]] = 0
			c.Address += 6
			return true
		},
		"dvm": func(args []string) (jumped bool) {
			// dvm x d q r
			// -->
			// q = x / d
			// r = (x / d) - (x % d)
			// x = 0
			x := c.GetValue(args[0])
			denominator := c.GetValue(args[1])
			quotient := x / denominator
			remainder := denominator - (x % denominator)

			c.Registers[args[0]] = 0
			c.Registers[args[2]] = quotient
			c.Registers[args[3]] = remainder
			c.Address += 8
			return true
		},
		"out": func(args []string) (jumped bool) {
			c.Output = append(c.Output, c.GetValue(args[0]))
			return false
		},
	}
}

func isRegister(arg string) bool {
	return unicode.IsLetter([]rune(arg)[0])
}
