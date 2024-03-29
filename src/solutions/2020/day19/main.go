package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"aoc/src/solutions/2020/day19/pda"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 19
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	messages, rules := parse(input)
	return countValidMessages(rules, messages)
}

func part2(input []string) interface{} {
	messages, rules := parse(input)

	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"

	return countValidMessages(rules, messages)
}

func countValidMessages(rules map[string]string, messages []string) int {
	p := pda.NewPDA("0")
	for ruleID, rule := range rules {
		for _, sequence := range strings.Split(rule, " | ") {
			p.AddRule(ruleID, strings.Split(sequence, " "))
		}
	}

	count := 0
	for _, message := range messages {
		if p.Match(message) {
			count++
		}
	}

	return count
}

func parse(input []string) ([]string, map[string]string) {
	var messages []string
	rules := make(map[string]string)

	section := 0
	for _, line := range input {
		if line == "" {
			section++
			continue
		}

		if section == 0 {
			splitRule := strings.Split(line, ": ")
			rules[splitRule[0]] = strings.ReplaceAll(splitRule[1], "\"", "")
		} else if section == 1 {
			messages = append(messages, line)
		}
	}

	return messages, rules
}
