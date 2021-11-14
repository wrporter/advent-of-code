package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 16
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	moves := strings.Split(input[0], ",")
	dance := []rune("abcdefghijklmnop")

	for _, move := range moves {
		dance = performMove(move, dance)
	}

	return string(dance)
}

func part2(input []string) interface{} {
	moveSplit := strings.Split(input[0], ",")
	start := "abcdefghijklmnop"
	dance := []rune(start)
	iterations := 1_000_000_000

	// tiny speed up since the cycle for this input is small (60)
	memo := make(map[int][]rune)

	for i := 1; i <= iterations; i++ {
		for _, moveStr := range moveSplit {
			dance = performMove(moveStr, dance)
		}

		memo[i] = dance

		// did we encounter a cycle?
		if string(dance) == start {
			// jump to the last cycle
			i = iterations / i * i
			return string(memo[iterations-i])
		}
	}

	return string(dance)
}

func rotate(values []rune, amount int) []rune {
	if amount <= 0 || len(values) == 0 {
		return values
	}

	rotation := len(values) - (amount % len(values))
	values = append(values[rotation:], values[:rotation]...)

	return values
}

const (
	Spin     = 's'
	Exchange = 'x'
	Partner  = 'p'
)

func performMove(move string, dance []rune) []rune {
	moveType := move[0]
	switch moveType {
	case Spin:
		amount := conversion.StringToInt(move[1:])
		dance = rotate(dance, amount)
	case Exchange:
		args := strings.Split(move[1:], "/")
		i := conversion.StringToInt(args[0])
		j := conversion.StringToInt(args[1])
		dance[i], dance[j] = dance[j], dance[i]
	case Partner:
		args := strings.Split(move[1:], "/")
		name1 := rune(args[0][0])
		name2 := rune(args[1][0])
		var j, k int
		for i, program := range dance {
			if name1 == program {
				j = i
			}
			if name2 == program {
				k = i
			}
		}
		dance[j], dance[k] = dance[k], dance[j]
	}
	return dance
}
