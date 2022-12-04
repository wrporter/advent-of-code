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
	numberPositions := getNumberPositions(input)
	minSteps := getMinStepsBetweenNumbers(input, numberPositions)
	digits := getNumbers(numberPositions)

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
	digits := getNumbers(numbers)

	min := ints.MaxInt
	probability.Combo(digits, func(nums []int) {
		current := 0
		steps := 0
		for _, next := range nums {
			steps += minSteps[current][next]
			current = next
		}

		// go back to 0
		steps += minSteps[current][0]

		min = ints.Min(min, steps)
	})

	return min
}

func getNumbers(numberPositions map[int]geometry.Point) []int {
	var numbers []int
	for number := range numberPositions {
		if number == 0 {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func getMinStepsBetweenNumbers(grid []string, numbers map[int]geometry.Point) map[int]map[int]int {
	minSteps := make(map[int]map[int]int)
	for from := range numbers {
		minSteps[from] = make(map[int]int)
	}

	for from := range numbers {
		queue := []node{{
			Point: numbers[from],
			steps: 0,
		}}
		visited := map[geometry.Point]bool{numbers[from]: true}
		var current node

		for len(queue) > 0 && len(minSteps[from]) != len(numbers) {
			current, queue = queue[0], queue[1:]
			char := grid[current.Y][current.X]
			if bytes.IsDigit(char) {
				to := bytes.ToInt(char)
				minSteps[from][to] = current.steps
				minSteps[to][from] = current.steps
			}

			for _, dir := range geometry.Directions {
				p := current.Point.Move(dir)
				if !visited[p] && grid[p.Y][p.X] != '#' {
					visited[p] = true
					queue = append(queue, node{
						Point: p,
						steps: current.steps + 1,
					})
				}
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
	steps int
}
