package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strconv"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 3
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	gamma := ""
	epsilon := ""

	for i := 0; i < len(input[0]); i++ {
		numOnes := 0
		for _, number := range input {
			if number[i] == '1' {
				numOnes++
			}
		}
		if numOnes > len(input)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 32)
	e, _ := strconv.ParseInt(epsilon, 2, 32)
	return g * e
}

func part2(input []string) interface{} {
	oxygenRating := getRatingValue(input, 0, '1')
	co2Rating := getRatingValue(input, 0, '0')

	o, _ := strconv.ParseInt(oxygenRating, 2, 32)
	c, _ := strconv.ParseInt(co2Rating, 2, 32)

	return o * c
}

func getRatingValue(report []string, i int, desiredEqualBit byte) string {
	if len(report) == 1 {
		return report[0]
	}

	mostCommonBit := getDesiredBit(report, i, desiredEqualBit)

	var numbersToKeep []string
	for _, number := range report {
		if number[i] == mostCommonBit {
			numbersToKeep = append(numbersToKeep, number)
		}
	}

	return getRatingValue(numbersToKeep, i+1, desiredEqualBit)
}

func getDesiredBit(report []string, i int, desiredBit byte) byte {
	numOnes := 0
	numZeroes := 0
	for _, number := range report {
		if number[i] == '1' {
			numOnes++
		} else {
			numZeroes++
		}
	}

	if desiredBit == '0' && numZeroes < numOnes {
		return '0'
	} else if desiredBit == '0' && numZeroes > numOnes {
		return '1'
	} else if numOnes == numZeroes {
		return desiredBit
	} else if numOnes > numZeroes {
		return '1'
	} else {
		return '0'
	}
}
