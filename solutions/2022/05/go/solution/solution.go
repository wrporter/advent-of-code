package solution

import (
	"bytes"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/v2/contain"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	moves, stacks := parseInput(input)

	for _, move := range strings.Split(moves, "\n") {
		amount, from, to := parseMove(move)

		for i := 0; i < amount; i++ {
			//stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			//stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to].Push(stacks[from].Pop())
		}
	}

	return joinTopCrates(stacks)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	moves, stacks := parseInput(input)

	for _, move := range strings.Split(moves, "\n") {
		amount, from, to := parseMove(move)

		//stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:]...)
		//stacks[from] = stacks[from][:len(stacks[from])-amount]
		stacks[to].PushMany(stacks[from].PopMany(amount)...)
	}

	return joinTopCrates(stacks)
}

// func joinTopCrates(stacks [][]string) interface{} {
func joinTopCrates(stacks []*contain.Stack[string]) interface{} {
	var buffer bytes.Buffer
	for _, stack := range stacks {
		//buffer.WriteString(stack[len(stack)-1])
		buffer.WriteString(stack.Peek())
	}
	return buffer.String()
}

func parseInput(input string) (string, []*contain.Stack[string]) {
	splitInput := strings.Split(input, "\n\n")
	stackInput := splitInput[0]
	stackLines := strings.Split(stackInput, "\n")
	numStacks := (len(stackLines[0]) + 1) / 4

	//stacks := make([][]string, numStacks)
	stacks := make([]*contain.Stack[string], numStacks)
	for i := 0; i < numStacks; i++ {
		stacks[i] = contain.NewStack[string]()
	}

	for i := len(stackLines) - 1; i >= 0; i-- {
		line := stackLines[i]
		for stack := range stacks {
			index := stack * 4
			if line[index] == '[' {
				crate := line[index+1 : index+2]
				//stacks[stack] = append(stacks[stack], crate)
				stacks[stack].Push(crate)
			}
		}
	}

	moves := splitInput[1]
	return moves, stacks
}

func parseMove(move string) (int, int, int) {
	parts := strings.Split(move, " ")
	amount := convert.StringToInt(parts[1])
	from := convert.StringToInt(parts[3]) - 1
	to := convert.StringToInt(parts[5]) - 1
	return amount, from, to
}
