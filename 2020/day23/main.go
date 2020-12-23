package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/sample-input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	cups := parse(input)

	play(cups, 10)

	return 0
}

func play(cups []int, moves int) []int {
	current := 0
	var pickup [3]int

	for move := 1; move <= moves; move++ {
		fmt.Printf("-- move %d --\n", move)
		fmt.Printf("cups: %s\n", render(cups, current))

		for pick := 0; pick < 3; pick++ {
			cupPosition := (current + 1 + pick) % len(cups)
			// pickup the cups
			pickup[pick] = cups[cupPosition]
			// shift the cups to fill in the space that was taken
			cups[cupPosition] = cups[(current+4)%len(cups)]
		}
	}

	return cups
}

func render(cups []int, current int) string {
	var sb strings.Builder
	delimiter := ' '

	for i, cup := range cups {
		if i == current {
			sb.WriteString(fmt.Sprintf("(%d)", cup))
		} else {
			sb.WriteString(fmt.Sprintf("%d", cup))
		}
		if i < len(cups)-1 {
			sb.WriteRune(delimiter)
		}
	}

	return sb.String()
}

func part2(input []string) interface{} {
	return 0
}

func parse(input []string) []int {
	labels := strings.Split(input[0], "")
	cups, _ := conversion.ToInts(labels)
	return cups
}
