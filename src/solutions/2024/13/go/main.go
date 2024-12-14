package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

var numberRegex = regexp.MustCompile(`\d+`)

func part1(input string, _ ...interface{}) interface{} {
	return getTotalTokens(input, 0)
}

func part2(input string, _ ...interface{}) interface{} {
	return getTotalTokens(input, 10000000000000)
}

func getTotalTokens(input string, start int) interface{} {
	machines := strings.Split(input, "\n\n")
	total := 0
	for _, machine := range machines {
		values := lo.Map(numberRegex.FindAllString(machine, -1), func(item string, index int) int {
			return convert.StringToInt(item)
		})

		x1, y1 := values[0], values[1]
		x2, y2 := values[2], values[3]
		x3, y3 := values[4], values[5]

		total += getNumTokens(x1, y1, x2, y2, x3+start, y3+start)
	}
	return total
}

func getNumTokens(x1, y1, x2, y2, x3, y3 int) int {
	a := (y2*x3 - x2*y3) / (x1*y2 - y1*x2)
	b := (y1*x3 - x1*y3) / (y1*x2 - x1*y2)

	x := x1*a + x2*b
	y := y1*a + y2*b

	if x == x3 && y == y3 {
		return a*3 + b
	}

	return 0
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 13, Part1: part1, Part2: part2}
}
