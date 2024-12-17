package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	mymath "aoc/src/lib/go/ints"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var programRegex = regexp.MustCompile(`Register A: (\d+)
Register B: (\d+)
Register C: (\d+)

Program: (.+)`)

func part1(input string, _ ...interface{}) interface{} {
	A, B, C, program := parse(input)
	out := run(A, B, C, program)
	return strings.Join(strings.Split(out, ""), ",")
}

func part2(input string, _ ...interface{}) interface{} {
	_, B, C, program := parse(input)
	programStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(program)), ""), "[]")

	type node struct {
		count   int
		aBounds int
	}

	queue := []node{{1, 0}}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for A := n.aBounds; A < n.aBounds+8; A++ {
			output := run(A, B, C, program)
			if output == programStr {
				return A
			} else if output == programStr[len(program)-n.count:] {
				queue = append(queue, node{n.count + 1, A * 8})
			}
		}
	}

	return -1
}

func run(A int, B int, C int, program []int) string {
	combo := func(operand int) int {
		if operand == 4 {
			return A
		} else if operand == 5 {
			return B
		} else if operand == 6 {
			return C
		}
		return operand
	}

	pointer := 0
	output := &strings.Builder{}

	for pointer < len(program) {
		opcode, operand := program[pointer], program[pointer+1]

		switch opcode {
		case 0: // adv: aBounds = aBounds / 2^combo
			A = A / mymath.Pow(2, combo(operand))
		case 1: // bxl: B = B xor literal
			B = B ^ operand
		case 2: // bst: B = combo % 8
			B = combo(operand) % 8
		case 3: // jnz: jump to literal if aBounds != 0
			if A != 0 {
				pointer = operand
				continue
			}
		case 4: // bxc: B = B ^ C
			B = B ^ C
		case 5: // out: -> combo % 8
			output.WriteString(strconv.Itoa(combo(operand) % 8))
			//output = append(output, combo(operand)%8)
		case 6: // bdv: B = aBounds / 2^combo
			B = A / mymath.Pow(2, combo(operand))
		case 7: // cdv: C = aBounds / 2^combo
			C = A / mymath.Pow(2, combo(operand))
		}

		pointer += 2
	}

	return output.String()
}

func parse(input string) (int, int, int, []int) {
	match := programRegex.FindStringSubmatch(input)
	A := convert.StringToInt(match[1])
	B := convert.StringToInt(match[2])
	C := convert.StringToInt(match[3])
	program := convert.ToIntsV2(strings.Split(match[4], ","))
	return A, B, C, program
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 17, Part1: part1, Part2: part2}
}
