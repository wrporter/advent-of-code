package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"fmt"
)

func main() {
	year, day := 2018, 2
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	countTwo := 0
	countThree := 0

	for _, boxID := range input {
		counts := make(map[rune]int)
		for _, char := range boxID {
			if _, ok := counts[char]; ok {
				counts[char]++
			} else {
				counts[char] = 1
			}
		}

		twoCounted := false
		threeCounted := false
		for _, count := range counts {
			if count == 2 && !twoCounted {
				countTwo++
				twoCounted = true
			} else if count == 3 && !threeCounted {
				countThree++
				threeCounted = true
			}
		}
	}

	return countTwo * countThree
}

func part2(input []string) interface{} {
	for _, boxID1 := range input {
		for _, boxID2 := range input {
			countDiff := 0
			diffCharIndex := 0

			if boxID1 != boxID2 {
				for i := 0; i < len(boxID1); i++ {
					if boxID1[i] != boxID2[i] {
						countDiff++
						diffCharIndex = i
					}
				}

				if countDiff == 1 {
					return boxID1[:diffCharIndex] + boxID1[diffCharIndex+1:]
				}
			}
		}
	}

	return ""
}
