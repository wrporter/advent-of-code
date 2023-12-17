package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/mystrings"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/runes"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 21
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	password := "abcdefgh"
	p := []rune(password)

	for _, line := range input {
		op := parseOperation(line)
		args := op.args
		switch op.command {
		case swapPosition:
			x := convert.StringToInt(args[0])
			y := convert.StringToInt(args[1])
			p[x], p[y] = p[y], p[x]
		case swapLetter:
			i := indexOf(p, rune(args[0][0]))
			j := indexOf(p, rune(args[1][0]))
			p[i], p[j] = p[j], p[i]
		case rotateDirection:
			x := convert.StringToInt(args[0])
			p = runes.Rotate(p, x)
		case rotatePosition:
			i := indexOf(p, rune(args[0][0]))
			if i >= 4 {
				i += 1
			}
			i += 1
			p = runes.Rotate(p, i)
		case Reverse:
			i := convert.StringToInt(args[0])
			j := convert.StringToInt(args[1])
			p = reverse(p, i, j)
		case Move:
			x := convert.StringToInt(args[0])
			y := convert.StringToInt(args[1])
			p = move(p, x, y)
		}
		password = string(p)
	}

	return password
}

func part2(input []string) interface{} {
	password := "fbgdceah"
	p := []rune(password)
	answers, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/answers.txt", 2016, 21))
	input = mystrings.ReverseList(input)

	for a, line := range input {
		op := parseOperation(line)
		args := op.args
		prev := password
		switch op.command {
		case swapPosition:
			x := convert.StringToInt(args[0])
			y := convert.StringToInt(args[1])
			p[x], p[y] = p[y], p[x]
		case swapLetter:
			i := indexOf(p, rune(args[0][0]))
			j := indexOf(p, rune(args[1][0]))
			p[i], p[j] = p[j], p[i]
		case rotateDirection:
			// reversed by negating x
			x := convert.StringToInt(args[0])
			p = runes.Rotate(p, -x)
		case rotatePosition:
			// reversed by negating the pattern
			// definitely the most challenging part to reverse
			i := indexOf(p, rune(args[0][0]))
			if i != 0 && i%2 == 0 {
				i += len(p)
			}
			i = (i/2 + 1) % len(p)
			p = runes.Rotate(p, -i)
		case Reverse:
			i := convert.StringToInt(args[0])
			j := convert.StringToInt(args[1])
			p = reverse(p, i, j)
		case Move:
			// reversed by swapping x and y
			x := convert.StringToInt(args[0])
			y := convert.StringToInt(args[1])
			p = move(p, y, x)
		}
		password = string(p)
		if answers[a] != password {
			fmt.Printf("%3d: [%s] %s -> %s | %s\n", a+1, prev, password, answers[a], line)
		}
	}

	return password
}

func move(values []rune, from int, to int) []rune {
	tmp := values[from]
	result := values[:]
	result = runes.Remove(result, from)
	result = runes.Insert(result, to, tmp)
	return result
}

func reverse(values []rune, start int, end int) []rune {
	values = runes.Concat([][]rune{
		values[:start],
		runes.Reverse(values[start : end+1]),
		values[end+1:],
	})
	return values
}

var (
	swapPosition    = "swapp"
	swapLetter      = "swapl"
	rotateDirection = "rotated"
	rotatePosition  = "rotatep"
	Reverse         = "reverse"
	Move            = "move"
)

type operation struct {
	command string
	full    string
	args    []string
}

// swap position X with position Y
// swap letter X with letter Y
// rotate left/right X steps
// rotate based on position of letter X
// reverse positions X through Y
// move position X to position Y
func parseOperation(full string) operation {
	parts := strings.Fields(full)
	var args []string
	var command string
	if strings.Contains(full, "swap position") {
		command = swapPosition
		args = append(args, parts[2])
		args = append(args, parts[5])
	} else if strings.Contains(full, "swap letter") {
		command = swapLetter
		args = append(args, parts[2])
		args = append(args, parts[5])
	} else if strings.Contains(full, "rotate based") {
		command = rotatePosition
		args = append(args, parts[6])
	} else if strings.Contains(full, "rotate") {
		command = rotateDirection
		arg := parts[2]
		if strings.Contains(full, "left") {
			arg = "-" + arg
		}
		args = append(args, arg)
	} else if strings.Contains(full, "reverse") {
		command = Reverse
		args = append(args, parts[2])
		args = append(args, parts[4])
	} else if strings.Contains(full, "move") {
		command = Move
		args = append(args, parts[2])
		args = append(args, parts[5])
	}

	return operation{
		command: command,
		full:    full,
		args:    args,
	}
}

func indexOf(chars []rune, char rune) int {
	for i := range chars {
		if char == chars[i] {
			return i
		}
	}
	return -1
}
