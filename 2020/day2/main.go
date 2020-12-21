package main

import (
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	out.Day(2020, 2)
	input, _ := file.ReadFile("./2020/day2/input.txt")
	answer1 := part1(input)
	out.Part1(answer1)
	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]+): ([a-z]+)$`)

func part1(input []string) int {
	countValid := 0

	for _, passwordString := range input {
		match := regex.FindStringSubmatch(passwordString)
		lowest := conversion.StringToInt(match[1])
		highest := conversion.StringToInt(match[2])
		letter := rune(match[3][0])
		password := match[4]

		letterCount := 0
		for _, char := range password {
			if char == letter {
				letterCount++
			}
		}
		if letterCount >= lowest && letterCount <= highest {
			countValid++
		}
	}

	return countValid
}

func part2(input []string) int {
	countValid := 0

	for _, passwordString := range input {
		match := regex.FindStringSubmatch(passwordString)
		position1 := conversion.StringToInt(match[1])
		position2 := conversion.StringToInt(match[2])
		letter := rune(match[3][0])
		password := match[4]

		countContainsLetter := 0
		if containsLetter(password, letter, position1-1) {
			countContainsLetter++
		}
		if containsLetter(password, letter, position2-1) {
			countContainsLetter++
		}
		if countContainsLetter == 1 {
			countValid++
		}
	}

	return countValid
}

func containsLetter(password string, letter rune, index int) bool {
	return index >= 0 && index < len(password) && rune(password[index]) == letter
}
