package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	chunks := strings.Split(input, "\n\n")
	rules := strings.Split(chunks[0], "\n")
	updates := strings.Split(chunks[1], "\n")

	before := make(map[int]map[int]bool)
	for _, rule := range rules {
		parts := convert.ToIntsV2(strings.Split(rule, "|"))
		if before[parts[0]] == nil {
			before[parts[0]] = make(map[int]bool)
		}
		before[parts[0]][parts[1]] = true
	}

	sum := 0

	for _, updateStr := range updates {
		update := convert.ToIntsV2(strings.Split(updateStr, ","))
		correct := true

		for i := len(update) - 1; i >= 0 && correct; i-- {
			page := update[i]
			for _, after := range update[i+1:] {
				if !before[page][after] {
					correct = false
					break
				}
			}
		}

		if correct {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func part2(input string, _ ...interface{}) interface{} {
	chunks := strings.Split(input, "\n\n")
	rules := strings.Split(chunks[0], "\n")
	updates := strings.Split(chunks[1], "\n")

	before := make(map[int]map[int]bool)
	for _, rule := range rules {
		parts := convert.ToIntsV2(strings.Split(rule, "|"))
		if before[parts[0]] == nil {
			before[parts[0]] = make(map[int]bool)
		}
		before[parts[0]][parts[1]] = true
	}

	sum := 0

	for _, updateStr := range updates {
		update := convert.ToIntsV2(strings.Split(updateStr, ","))
		correct := true

		for i := len(update) - 1; i >= 0 && correct; i-- {
			page := update[i]
			for _, after := range update[i+1:] {
				if !before[page][after] {
					correct = false
					break
				}
			}
		}

		if !correct {
			sort.Slice(update, func(i, j int) bool {
				return before[update[i]][update[j]]
			})
			sum += update[len(update)/2]
		}
	}

	return sum
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 5, Part1: part1, Part2: part2}
}
