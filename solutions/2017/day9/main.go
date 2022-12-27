package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 9
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	score, _ := countScoreAndGarbage(input[0])
	return score
}

func countScoreAndGarbage(stream string) (score int, numGarbageCharacters int) {
	garbage := false
	depth := 0

	for i := 0; i < len(stream); i++ {
		char := stream[i]

		if char == '!' {
			i += 1
		} else if char == '>' {
			garbage = false
		} else if garbage {
			numGarbageCharacters += 1
		} else if char == '{' {
			depth += 1
		} else if char == '}' {
			score += depth
			depth -= 1
		} else if char == '<' {
			garbage = true
		}
	}

	return score, numGarbageCharacters
}

func part2(input []string) interface{} {
	_, numGarbageCharacters := countScoreAndGarbage(input[0])
	return numGarbageCharacters
}
