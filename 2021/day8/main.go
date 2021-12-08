package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 8
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	displays := parseInput(input)

	sum := 0
	for _, d := range displays {
		for _, digit := range d.output {
			size := len(digit)
			if size == 2 || size == 4 || size == 3 || size == 7 {
				sum++
			}
		}
	}

	return sum
}

func part2(input []string) interface{} {
	displays := parseInput(input)
	sum := 0

	for _, d := range displays {
		mapping := d.getMapping()
		value := ""

		for _, pattern := range d.output {
			value += strconv.Itoa(mapping[pattern])
		}

		sum += convert.StringToInt(value)
	}

	return sum
}

func parseInput(input []string) []display {
	displays := make([]display, len(input))
	for i, line := range input {
		parts := strings.Split(line, " | ")
		displays[i] = display{
			input:  sortStrings(strings.Fields(parts[0])),
			output: sortStrings(strings.Fields(parts[1])),
		}
	}
	return displays
}

func sortStrings(values []string) []string {
	sorted := make([]string, len(values))
	for i, value := range values {
		sorted[i] = mystrings.SortString(value)
	}
	return sorted
}

type display struct {
	input  []string
	output []string
}

func valuesToSet(m map[int]string) map[string]bool {
	result := make(map[string]bool)
	for _, v := range m {
		result[v] = true
	}
	return result
}

func (d display) getMapping() map[string]int {
	digits := make(map[int]string)

	for _, pattern := range d.input {
		switch len(pattern) {
		case 2:
			digits[1] = pattern
		case 4:
			digits[4] = pattern
		case 3:
			digits[7] = pattern
		case 7:
			digits[8] = pattern
		}
	}

	digits[6] = d.getDigit(6, func(pattern string) bool {
		return len(intersect(pattern, digits[1])) == 1
	})
	digits[0] = d.getDigit(6, func(pattern string) bool {
		return len(intersect(pattern, digits[4])) == 3 && !valuesToSet(digits)[pattern]
	})
	digits[9] = d.getDigit(6, func(pattern string) bool {
		return !valuesToSet(digits)[pattern]
	})
	digits[5] = d.getDigit(5, func(pattern string) bool {
		return len(intersect(pattern, digits[6])) == 5
	})
	digits[3] = d.getDigit(5, func(pattern string) bool {
		return !valuesToSet(digits)[pattern] && len(intersect(pattern, digits[9])) == 5
	})
	digits[2] = d.getDigit(5, func(pattern string) bool {
		return !valuesToSet(digits)[pattern]
	})

	mapping := make(map[string]int)
	for k, v := range digits {
		mapping[v] = k
	}

	return mapping
}

func (d display) getDigit(size int, test func(pattern string) bool) string {
	for _, pattern := range d.input {
		if len(pattern) == size && test(pattern) {
			return pattern
		}
	}
	return ""
}

func intersect(a string, b string) string {
	set := make(map[rune]bool)
	hash := make(map[rune]bool)

	for _, char := range a {
		hash[char] = true
	}

	for _, char := range b {
		if _, found := hash[char]; found {
			set[char] = true
		}
	}

	result := strings.Builder{}
	for char := range set {
		result.WriteRune(char)
	}

	return result.String()
}
