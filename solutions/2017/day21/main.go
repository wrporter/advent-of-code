package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/stringgrid"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 21
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	iterations := 5
	image := iterate(input, iterations)
	return countOn(image)
}

func part2(input []string) interface{} {
	iterations := 18
	image := iterate(input, iterations)
	return countOn(image)
}

func iterate(input []string, iterations int) []string {
	rules := parseInput(input)
	image := []string{
		".#.",
		"..#",
		"###",
	}

	for i := 0; i < iterations; i++ {
		squareSize := 3
		if len(image)%2 == 0 {
			squareSize = 2
		}
		size := len(image)

		steps := size / squareSize
		next := make([]string, (squareSize+1)*steps)
		for y := 0; y < steps; y++ {
			for x := 0; x < steps; x++ {
				square := make([]string, squareSize)
				for row := 0; row < squareSize; row++ {
					square[row] = image[row+(y*squareSize)][x*squareSize : x*squareSize+squareSize]
				}
				enhancedSquare := enhance(rules, square)
				for row := range enhancedSquare {
					next[row+(y*(squareSize+1))] += enhancedSquare[row]
				}
			}
		}

		image = next
	}
	return image
}

func enhance(rules map[int]map[string]string, square []string) []string {
	return expand(rules[len(square)][flatten(square)])
}

func flatten(grid []string) string {
	return strings.Join(grid, "/")
}

func expand(flattened string) []string {
	return strings.Split(flattened, "/")
}

func countOn(grid []string) int {
	numOn := 0
	for _, row := range grid {
		for _, value := range row {
			if value == '#' {
				numOn++
			}
		}
	}
	return numOn
}

func getArrangements(grid []string) [][]string {
	var a [][]string
	a = append(a, grid)
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.FlipLR(grid))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	return a
}

func parseInput(input []string) map[int]map[string]string {
	rules := make(map[int]map[string]string)

	for _, line := range input {
		parts := strings.Fields(line)
		size := len(strings.Split(parts[0], "/")[0])

		if rules[size] == nil {
			rules[size] = make(map[string]string)
		}

		// Memoize the arrangements of rules for fast lookup
		for _, arrangement := range getArrangements(expand(parts[0])) {
			rules[size][flatten(arrangement)] = parts[2]
		}
	}

	return rules
}
