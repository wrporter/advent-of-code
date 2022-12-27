package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 6
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	count := 0
	group := ""

	for i, line := range input {
		group += line

		if line == "" || (i+1) == len(input) {
			uniqueAnswers := make(map[rune]bool)
			for _, answer := range group {
				uniqueAnswers[answer] = true
			}

			count += len(uniqueAnswers)
			group = ""
		}
	}

	return count
}

func part2(input []string) interface{} {
	count := 0
	var group []string

	for i, line := range input {
		if line != "" {
			group = append(group, line)
		}

		if line == "" || (i+1) == len(input) {
			uniqueGroupAnswers := make(map[rune]bool)
			for _, answers := range group {
				for _, answer := range answers {
					uniqueGroupAnswers[answer] = true
				}
			}

			for answer := range uniqueGroupAnswers {
				if mystrings.Every(group, func(personAnswers string) bool {
					return strings.ContainsRune(personAnswers, answer)
				}) {
					count++
				}
			}

			group = make([]string, 0)
		}
	}

	return count
}
