package main

import (
	"aoc/src/lib/go/file"
	"fmt"
	"strings"
)

func countNiceStringsPart1(words []string) int {
	count := 0
	for _, str := range words {
		if isNicePart1(str) {
			count++
		}
	}
	return count
}

func isNicePart1(str string) bool {
	return hasAtLeastThreeVowels(str) &&
		hasConsecutiveLetter(str) &&
		!hasBadStrings(str)
}

func countNiceStringsPart2(words []string) int {
	count := 0
	for _, str := range words {
		if isNicePart2(str) {
			count++
		}
	}
	return count
}

func isNicePart2(str string) bool {
	return hasRepeatPairWithoutOverlap(str) &&
		hasRepeatLetterWithOneInBetween(str)
}

func hasRepeatLetterWithOneInBetween(str string) bool {
	for i, char := range str {
		if i+2 < len(str) && char == rune(str[i+2]) {
			return true
		}
	}
	return false
}

func hasRepeatPairWithoutOverlap(str string) bool {
	pairs := make(map[string]int)
	var prev rune

	for i, letter := range str {
		if i != 0 {
			pair := string(prev) + string(letter)
			if index, ok := pairs[pair]; ok {
				if index < i-2 {
					return true
				} else {
					continue
				}
			}
			pairs[pair] = i - 1
		}
		prev = letter
	}

	return false
}

func hasConsecutiveLetter(str string) bool {
	var prev rune
	for _, char := range str {
		if prev == char {
			return true
		}
		prev = char
	}
	return false
}

func hasAtLeastThreeVowels(str string) bool {
	vowels := 0
	for _, char := range str {
		if isVowel(char) {
			vowels++
		}
	}
	return vowels >= 3
}

func isVowel(char rune) bool {
	return char == 'a' ||
		char == 'e' ||
		char == 'i' ||
		char == 'o' ||
		char == 'u'
}

func hasBadStrings(str string) bool {
	return strings.Contains(str, "ab") ||
		strings.Contains(str, "cd") ||
		strings.Contains(str, "pq") ||
		strings.Contains(str, "xy")
}

func main() {
	lines, _ := file.ReadFile("./2015/day5/input.txt")
	fmt.Println(countNiceStringsPart1(lines))
	fmt.Println(countNiceStringsPart2(lines))
}
