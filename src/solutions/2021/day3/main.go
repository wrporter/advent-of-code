package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strconv"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 3
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	var gamma, epsilon string

	for i := 0; i < len(input[0]); i++ {
		gamma += string(getDesiredBit(input, i, '0'))
		epsilon += string(getDesiredBit(input, i, '1'))
	}

	return binaryToInt(gamma) * binaryToInt(epsilon)
}

func part2(input []string) interface{} {
	oxygenRating := getRatingValue(input, 0, '1')
	co2Rating := getRatingValue(input, 0, '0')

	return binaryToInt(oxygenRating) * binaryToInt(co2Rating)
}

func getRatingValue(report []string, i int, desiredBit byte) string {
	if len(report) == 1 {
		return report[0]
	}

	mostCommonBit := getDesiredBit(report, i, desiredBit)

	var numbersToKeep []string
	for _, number := range report {
		if number[i] == mostCommonBit {
			numbersToKeep = append(numbersToKeep, number)
		}
	}

	return getRatingValue(numbersToKeep, i+1, desiredBit)
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

func binaryToInt(gamma string) int64 {
	g, _ := strconv.ParseInt(gamma, 2, 32)
	return g
}
