package main

import (
	"aoc/src/lib/go/aoc"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grammar, words := parse(input)

	sum := 0
	for _, word := range words {
		if matches(grammar, word) {
			sum += 1
		}
	}
	return sum

	//p := pda.NewPDA("S")
	//p.AddBNFRules([]string{"S: " + strings.Join(grammar, " S | ") + " | !"})
	//
	//sum := 0
	//for _, design := range words {
	//	if p.Match(design) {
	//		sum++
	//	}
	//}
	//return sum
}

func part2(input string, _ ...interface{}) interface{} {
	grammar, words := parse(input)

	sum := 0
	for _, word := range words {
		sum += countMatches(make(map[string]int), grammar, word)
	}
	return sum
}

func matches(grammar []string, input string) bool {
	if input == "" {
		return true
	}

	for _, prefix := range grammar {
		if strings.HasPrefix(input, prefix) {
			if matches(grammar, strings.TrimPrefix(input, prefix)) {
				return true
			}
		}
	}

	return false
}

func countMatches(cache map[string]int, grammar []string, input string) int {
	if val, found := cache[input]; found {
		return val
	}

	if input == "" {
		return 1
	}

	sum := 0
	for _, prefix := range grammar {
		if strings.HasPrefix(input, prefix) {
			sum += countMatches(cache, grammar, strings.TrimPrefix(input, prefix))
		}
	}

	cache[input] = sum
	return sum
}

func parse(input string) ([]string, []string) {
	parts := strings.Split(input, "\n\n")
	grammar := strings.Split(parts[0], ", ")
	words := strings.Split(parts[1], "\n")
	return grammar, words
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 19, Part1: part1, Part2: part2}
}
