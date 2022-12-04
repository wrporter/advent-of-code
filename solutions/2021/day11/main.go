package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 11
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	octopuses := parseInput(input)
	numSteps := 100
	numFlashes := 0

	for step := 1; step <= numSteps; step++ {
		numFlashes += performStep(octopuses)
		//fmt.Printf("After step %d\n", step)
		//display(octopuses)
	}

	return numFlashes
}

func part2(input []string) interface{} {
	octopuses := parseInput(input)

	numSteps := 1000
	for step := 1; step <= numSteps; step++ {
		numFlashes := performStep(octopuses)
		if numFlashes == 100 {
			return step
		}
	}

	return -1
}

func performStep(octopuses [][]int) int {
	for y, row := range octopuses {
		for x := range row {
			octopuses[y][x]++
		}
	}

	flashed := make(map[geometry.Point]bool)
	for y, row := range octopuses {
		for x, level := range row {
			if level > 9 {
				flash(octopuses, flashed, x, y)
			}
		}
	}

	return len(flashed)
}

func flash(octopuses [][]int, flashed map[geometry.Point]bool, x int, y int) {
	point := geometry.NewPoint(x, y)
	flashed[point] = true

	for _, direction := range geometry.AllDirections {
		neighbor := point.Add(direction)
		if neighbor.Y >= 0 && neighbor.Y < len(octopuses) && neighbor.X >= 0 && neighbor.X < len(octopuses[neighbor.Y]) {
			if !flashed[neighbor] {
				octopuses[neighbor.Y][neighbor.X]++

				if octopuses[neighbor.Y][neighbor.X] > 9 {
					flash(octopuses, flashed, neighbor.X, neighbor.Y)
				}
			}
		}
	}

	octopuses[y][x] = 0
}

func parseInput(input []string) [][]int {
	octopuses := make([][]int, len(input))
	for i, line := range input {
		octopuses[i], _ = convert.ToInts(strings.Split(line, ""))
	}
	return octopuses
}

func display(grid [][]int) {
	b := strings.Builder{}
	for _, row := range grid {
		b.WriteString(ints.Join(row, ""))
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}
