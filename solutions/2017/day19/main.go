package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
	"unicode"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 19
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	letters, _ := routePacket(input)
	return letters
}

func part2(input []string) interface{} {
	_, numSteps := routePacket(input)
	return numSteps
}

func routePacket(input []string) (string, int) {
	grid := convert.ToRuneGrid(input)
	start := getStartPosition(grid)
	current := geometry.Vector{
		Point:     start,
		Direction: geometry.Down,
	}

	queue := []geometry.Vector{current}
	visited := make(map[geometry.Point]bool)
	var letters []rune
	numSteps := 0

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		visited[current.Point] = true
		numSteps++

		char := grid[current.Y][current.X]
		if unicode.IsLetter(char) {
			letters = append(letters, char)
		}

		if char == '+' {
			for _, direction := range geometry.Directions {
				next := current.Point.Move(direction)
				if !visited[next] &&
					next.Y >= 0 && next.Y < len(grid) &&
					next.X >= 0 && next.X < len(grid[next.Y]) &&
					grid[next.Y][next.X] != ' ' {
					queue = append(queue, geometry.Vector{
						Point:     next,
						Direction: direction,
					})
				}
			}
		} else {
			next := current.Move(current.Direction)
			if grid[next.Y][next.X] != ' ' {
				queue = append(queue, geometry.Vector{
					Point:     next,
					Direction: current.Direction,
				})
			}
		}
	}

	return string(letters), numSteps
}

func getStartPosition(grid [][]rune) geometry.Point {
	var start geometry.Point
	for x, value := range grid[0] {
		if value == '|' {
			start = geometry.NewPoint(x, 0)
		}
	}
	return start
}
