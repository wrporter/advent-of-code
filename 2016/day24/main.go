package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/bytes"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/probability"
	"github.com/wrporter/advent-of-code/internal/common/runes"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
	"unicode"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 24
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numbers := getNumberPositions(input)
	minSteps := getMinStepsBetweenNumbers(input, numbers)
	digits := getKeys(numbers)

	min := ints.MaxInt
	probability.Combo(digits, func(nums []int) {
		current := 0
		steps := 0
		for _, next := range nums {
			steps += minSteps[current][next]
			current = next
		}
		min = ints.Min(min, steps)
	})

	return min
}

func part2(input []string) interface{} {
	numbers := getNumberPositions(input)
	minSteps := getMinStepsBetweenNumbers(input, numbers)
	digits := getKeys(numbers)

	min := ints.MaxInt
	probability.Combo(digits, func(nums []int) {
		current := 0
		steps := 0
		for _, next := range nums {
			steps += minSteps[current][next]
			current = next
		}
		steps += minSteps[current][0]
		min = ints.Min(min, steps)
	})

	return min
}

func getKeys(numbers map[int]geometry.Point) []int {
	var digits []int
	for number := range numbers {
		if number == 0 {
			continue
		}
		digits = append(digits, number)
	}
	return digits
}

func getMinStepsBetweenNumbers(input []string, numbers map[int]geometry.Point) map[int]map[int]int {
	minStepsBetweenNumbers := make(map[int]map[int]int)
	for from := range numbers {
		minStepsBetweenNumbers[from] = make(map[int]int)
		minSteps := getMinSteps(input, from, numbers)
		minStepsBetweenNumbers[from] = minSteps
	}
	return minStepsBetweenNumbers
}

// TODO: 2 slight optimizations:
//     1. Stop the BFS once all numbers have been visited.
//     2. Don't bother calculating the distance twice between numbers.
func getMinSteps(grid []string, from int, numbers map[int]geometry.Point) map[int]int {
	minSteps := make(map[int]int)
	queue := []node{{
		Point: numbers[from],
		from:  from,
		steps: 0,
	}}
	seen := map[geometry.Point]bool{numbers[from]: true}
	var current node

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		char := grid[current.Y][current.X]
		if bytes.IsDigit(char) {
			minSteps[bytes.ToInt(char)] = current.steps
		}

		for _, dir := range geometry.Directions {
			p := current.Point.Move(dir)
			if !seen[p] && grid[p.Y][p.X] != '#' {
				seen[p] = true
				queue = append(queue, node{
					Point: p,
					from:  current.from,
					steps: current.steps + 1,
				})
			}
		}
	}

	return minSteps
}

func getNumberPositions(input []string) map[int]geometry.Point {
	numbers := make(map[int]geometry.Point)
	for y, row := range input {
		for x, char := range row {
			if unicode.IsDigit(char) {
				number := runes.ToInt(char)
				position := geometry.NewPoint(x, y)
				numbers[number] = position
			}
		}
	}
	return numbers
}

type node struct {
	geometry.Point
	from  int
	steps int
}
