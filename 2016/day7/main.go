package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
)

func part1(input []string) int {
	countSupportingTLS := 0

	for _, ip := range input {
		if hasABBA(ip) {
			countSupportingTLS++
		}
	}

	return countSupportingTLS
}

func hasABBA(ip string) bool {
	outsideBrackets, insideBrackets := extractSequences(ip)

	if !some(outsideBrackets, hasRepeatBackwardsPair) ||
		some(insideBrackets, hasRepeatBackwardsPair) {
		return false
	}

	return true
}

func extractSequences(ip string) ([]string, []string) {
	var outsideBrackets []string
	var insideBrackets []string
	var current []rune

	for i, char := range ip {
		if char == '[' {
			outsideBrackets = append(outsideBrackets, string(current))
			current = make([]rune, 0)
		} else if char == ']' {
			insideBrackets = append(insideBrackets, string(current))
			current = make([]rune, 0)
		} else if (i + 1) == len(ip) {
			current = append(current, char)
			outsideBrackets = append(outsideBrackets, string(current))
		} else {
			current = append(current, char)
		}
	}

	return outsideBrackets, insideBrackets
}

func some(values []string, test func(string) bool) bool {
	for _, value := range values {
		if test(value) {
			return true
		}
	}
	return false
}

func hasRepeatBackwardsPair(str string) bool {
	for i := range str {
		if i >= 3 {
			char1 := str[i-3]
			char2 := str[i-2]
			char3 := str[i-1]
			char4 := str[i]

			if char1 == char4 && char2 == char3 && char1 != char2 {
				return true
			}
		}
	}
	return false
}

func main() {
	input, _ := file.ReadFile("./2016/day7/input.txt")
	answer1 := part1(input)
	fmt.Println(answer1)
}
