package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
)

func part1(input []string) int {
	countSupportingTLS := 0

	for _, ip := range input {
		if supportsTLS(ip) {
			countSupportingTLS++
		}
	}

	return countSupportingTLS
}

func part2(input []string) int {
	countSupportsSSL := 0

	for _, ip := range input {
		if supportsSSL(ip) {
			countSupportsSSL++
		}
	}

	return countSupportsSSL
}

func supportsSSL(ip string) bool {
	outsideBrackets, insideBrackets := extractSequences(ip)
	babs := extractABAsOrBABs(insideBrackets)
	supportsSSL := everyABAHasBAB(outsideBrackets, babs)
	return supportsSSL
}

func everyABAHasBAB(outsideBrackets []string, babs []string) bool {
	hasBAB := func(aba string) bool {
		return abaHasBAB(aba, babs)
	}

	for _, sequence := range outsideBrackets {
		abas := extractABAsOrBABsFromSequence(sequence)
		if some(abas, hasBAB) {
			return true
		}
	}

	return false
}

func abaHasBAB(aba string, babs []string) bool {
	for _, bab := range babs {
		if aba[0] == bab[1] &&
			aba[1] == bab[0] &&
			aba[1] == bab[2] {
			return true
		}
	}
	return false
}

func extractABAsOrBABs(sequences []string) []string {
	var abas []string
	for _, sequence := range sequences {
		abas = append(abas, extractABAsOrBABsFromSequence(sequence)...)
	}
	return abas
}

func extractABAsOrBABsFromSequence(sequence string) []string {
	var abas []string
	for i := 2; i < len(sequence); i++ {
		if sequence[i-2] == sequence[i] &&
			sequence[i] != sequence[i-1] {
			abas = append(abas, sequence[i-2:i+1])
		}
	}
	return abas
}

func supportsTLS(ip string) bool {
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
	answer2 := part2(input)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
