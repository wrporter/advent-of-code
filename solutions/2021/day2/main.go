package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 2
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	commands := parseInput(input)
	depth := 0
	horizontal := 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			horizontal += c.amount
		case "up":
			depth -= c.amount
		case "down":
			depth += c.amount
		}
	}

	return depth * horizontal
}

func part2(input []string) interface{} {
	commands := parseInput(input)
	aim := 0
	depth := 0
	horizontal := 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			horizontal += c.amount
			depth += c.amount * aim
		case "up":
			aim -= c.amount
		case "down":
			aim += c.amount
		}
	}

	return depth * horizontal
}

func parseInput(input []string) []command {
	commands := make([]command, len(input))
	for i, line := range input {
		parts := strings.Fields(line)
		commands[i] = command{
			direction: parts[0],
			amount:    convert.StringToInt(parts[1]),
		}
	}
	return commands
}

type command struct {
	direction string
	amount    int
}
