package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"math"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	games := parseInput(input)
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0

	for id, game := range games {
		possible := true

		for _, set := range game {
			for color, amount := range set {
				if bag[color] < amount {
					possible = false
				}
			}
		}

		if possible {
			sum += id
		}
	}

	return sum
}

func part2(input string, _ ...interface{}) interface{} {
	games := parseInput(input)
	sum := 0

	for _, game := range games {
		high := map[string]int{
			"red":   math.MinInt,
			"green": math.MinInt,
			"blue":  math.MinInt,
		}

		for _, set := range game {
			for color, amount := range set {
				high[color] = max(high[color], amount)
			}
		}

		product := 1
		for _, amount := range high {
			product *= amount
		}
		sum += product
	}

	return sum
}

func parseInput(input string) map[int][]map[string]int {
	games := make(map[int][]map[string]int)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		id := convert.StringToInt(strings.Split(parts[0], " ")[1])
		sets := strings.Split(parts[1], "; ")

		for _, setStr := range sets {
			cubes := strings.Split(setStr, ", ")
			set := map[string]int{"red": 0, "green": 0, "blue": 0}

			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")
				amount := convert.StringToInt(cubeParts[0])
				color := cubeParts[1]

				set[color] = amount
			}

			games[id] = append(games[id], set)
		}
	}

	return games
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 2, Part1: part1, Part2: part2}
}
