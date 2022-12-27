package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"sort"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 10
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	score := 0
	for _, line := range input {
		got, _, _, index := parse(line)
		if index < len(line) {
			score += part1Points[got]
		}
	}
	return score
}

func part2(input []string) interface{} {
	var scores []int

	for _, line := range input {
		_, _, remaining, index := parse(line)
		if index >= len(line) {
			total := 0
			for _, char := range remaining[:len(remaining)-1] {
				total = (5 * total) + part2Points[char]
			}
			scores = append(scores, total)
		}
	}

	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}

func parse(input string) (string, string, []string, int) {
	return parseRec(strings.Split(input, ""), []string{EOF}, 0)
}

func parseRec(line []string, stack []string, index int) (string, string, []string, int) {
	if index >= len(line) {
		if len(stack) == 1 && stack[0] == EOF {
			// We reached the end and the syntax is correct.
			return EOF, EOF, stack, index
		}
		// We reached the end and the syntax is incorrect.
		return EOF, stack[0], stack, index
	}

	cur := line[index]

	if cur == stack[0] {
		// We reached a matching closing character and pop it off the stack.
		return parseRec(line, stack[1:], index+1)
	}

	next, ok := pairs[cur]
	if !ok {
		// We reached an invalid character (corrupted).
		return cur, stack[0], stack, index
	}

	// We reached a new opening character and continue parsing.
	return parseRec(line, append([]string{next}, stack...), index+1)
}

const EOF = "EOF"

var pairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var part1Points = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var part2Points = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}
