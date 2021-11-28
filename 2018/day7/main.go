package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"sort"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 7
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$`)

func part1(input []string) interface{} {
	dependencies := make(map[rune]map[rune]bool)
	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		step := rune(match[1][0])
		before := rune(match[2][0])
		if dependencies[step] == nil {
			dependencies[step] = map[rune]bool{before: true}
		} else {
			dependencies[step][before] = true
		}
		if dependencies[before] == nil {
			dependencies[before] = make(map[rune]bool)
		}
	}

	var instructions []rune
	for ins := range dependencies {
		instructions = append(instructions, ins)
	}
	sort.SliceStable(instructions, func(i, j int) bool {
		return less(dependencies, instructions[i], instructions[j])
	})
	return string(instructions)
}

func part2(input []string) interface{} {
	return 0
}

func less(m map[rune]map[rune]bool, step1 rune, step2 rune) bool {
	step1IsBeforeStep2 := isBefore(m, step1, step2)
	step2IsBeforeStep1 := isBefore(m, step2, step1)
	if !step1IsBeforeStep2 && !step2IsBeforeStep1 {
		return step1 < step2
	}

	if step2IsBeforeStep1 {
		return false
	}
	return step1IsBeforeStep2
}

func isBefore(m map[rune]map[rune]bool, step1 rune, step2 rune) bool {
	for before := range m[step1] {
		if before == step2 || isBefore(m, before, step2) {
			return true
		}
	}
	return false
}
