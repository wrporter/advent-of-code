package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

func countNiceStrings(words []string) int {
	count := 0
	for _, str := range words {
		if isNice(str) {
			count++
		}
	}
	return count
}

func isNice(str string) bool {
	return hasAtLeastThreeVowels(str) &&
		hasConsecutiveLetter(str) &&
		!hasBadStrings(str)
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
	fmt.Println(countNiceStrings(lines))
}
