package main

import (
	"aoc/src/lib/go/aoc"
	"math"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last byte

		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				first = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				last = line[i]
				break
			}
		}

		value, _ := strconv.Atoi(string(first) + string(last))
		sum += value
	}

	return sum
}

var numbers = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// MAX_NUM_LENGTH is the max size any of the digits can be in word form.
const MAX_NUM_LENGTH = 5

// Part2 uses a sliding window algorithm to find the first and last digit. This
// is more performant, but there is a much higher chance for error due to off-
// by-one errors.
func part2(input string, _ ...interface{}) interface{} {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last string

		for i := 0; i < len(line) && first == ""; i++ {
			for size := 1; size <= MAX_NUM_LENGTH && i+size <= len(line) && first == ""; size++ {
				window := line[i : i+size]
				if digit, ok := numbers[window]; ok {
					first = digit
				}
			}
		}

		for i := len(line); i >= 1 && last == ""; i-- {
			for size := 1; size <= MAX_NUM_LENGTH && i-size >= 0 && last == ""; size++ {
				window := line[i-size : i]
				if digit, ok := numbers[window]; ok {
					last = digit
				}
			}
		}

		value, _ := strconv.Atoi(first + last)
		sum += value
	}

	return sum
}

// Part2_alternative is less performant, but easier to think through with less
// potential for bugs.
func part2_alternative(input string, _ ...interface{}) interface{} {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last string
		locations := make(map[int]string)

		low := math.MaxInt
		high := math.MinInt
		for num := range numbers {
			if i := strings.Index(line, num); i >= 0 {
				locations[i] = num
				low = min(low, i)
				high = max(high, i)
			}
			if i := strings.LastIndex(line, num); i >= 0 {
				locations[i] = num
				low = min(low, i)
				high = max(high, i)
			}
		}

		first = numbers[locations[low]]
		last = numbers[locations[high]]

		value, _ := strconv.Atoi(first + last)
		sum += value
	}

	return sum
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 1, Part1: part1, Part2: part2}
}
