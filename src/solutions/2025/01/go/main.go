package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"strings"
)

const DialStart = 50
const DialMod = 100

func part1(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	password := 0
	dial := DialStart

	for _, line := range lines {
		rotation := line[0]
		amount := convert.StringToInt(line[1:])

		if rotation == 'R' {
			dial = (dial + amount) % DialMod
		} else if rotation == 'L' {
			dial = (dial - amount) % DialMod
			if dial < 0 {
				dial = dial + DialMod
			}
		}

		if dial == 0 {
			password++
		}
	}

	return password
}

func part2(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	password := 0
	dial := DialStart

	for _, line := range lines {
		rotation := line[0]
		amount := convert.StringToInt(line[1:])

		if rotation == 'R' {
			for num := 1; num <= amount; num++ {
				dial = (dial + 1) % DialMod
				if dial == 0 {
					password++
				}
			}
		} else if rotation == 'L' {
			for num := 1; num <= amount; num++ {
				dial = dial - 1
				if dial == 0 {
					password++
				}
				if dial == -1 {
					dial = 99
				}
			}
		}
	}

	return password
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 1, Part1: part1, Part2: part2}
}
