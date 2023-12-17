package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/mystrings"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 16
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	state := input[0]
	length := convert.StringToInt(input[1])
	curve := getDragonCurve(state, length)
	return getChecksum(curve)
}

func part2(input []string) interface{} {
	state := input[0]
	length := 35651584
	curve := getDragonCurve(state, length)
	return getChecksum(curve)
}

func getChecksum(curve string) string {
	checksum := []rune(curve)
	for len(checksum)%2 == 0 {
		var next []rune
		for i := 0; i < len(checksum); i += 2 {
			if checksum[i] == checksum[i+1] {
				next = append(next, '1')
			} else {
				next = append(next, '0')
			}
		}
		checksum = next
	}
	return string(checksum)
}

func getDragonCurve(state string, length int) string {
	curve := state
	for len(curve) < length {
		a := curve
		b := a
		b = mystrings.Reverse(b)
		b = not(b)
		curve = a + "0" + b
	}
	return curve[:length]
}

func not(bits string) string {
	result := []rune(bits)
	for i, bit := range result {
		if bit == '0' {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}
