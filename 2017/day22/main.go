package main

import (
	"fmt"
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

	year, day := 2017, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	grid, carrier := parseInput(input)

	//visualize(grid)

	bursts := 10_000
	numBurstsCausedAnInfection := 0
	for burst := 0; burst < bursts; burst++ {
		if grid[carrier.Point] == Infected {
			carrier.Direction = carrier.Direction.Rotate(90)
			delete(grid, carrier.Point)
		} else {
			carrier.Direction = carrier.Direction.Rotate(-90)
			grid[carrier.Point] = Infected
			numBurstsCausedAnInfection++
		}
		carrier.Point = carrier.Add(carrier.Direction)
		//visualize(grid)
	}

	return numBurstsCausedAnInfection
}

func part2(input []string) interface{} {
	grid, carrier := parseInput(input)

	bursts := 10_000_000
	numBurstsCausedAnInfection := 0
	for burst := 0; burst < bursts; burst++ {
		if state, ok := grid[carrier.Point]; state == Clean || !ok {
			carrier.Direction = carrier.Direction.Rotate(-90)
			grid[carrier.Point] = Weakened
		} else if state == Weakened {
			grid[carrier.Point] = Infected
			numBurstsCausedAnInfection++
		} else if state == Infected {
			carrier.Direction = carrier.Direction.Rotate(90)
			grid[carrier.Point] = Flagged
		} else if state == Flagged {
			carrier.Direction = carrier.Direction.Rotate(180)
			delete(grid, carrier.Point)
		}
		carrier.Point = carrier.Add(carrier.Direction)
	}

	return numBurstsCausedAnInfection
}

func parseInput(input []string) (map[geometry.Point]rune, *geometry.Vector) {
	grid := make(map[geometry.Point]rune)
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				grid[geometry.NewPoint(x, y)] = Infected
			}
		}
	}

	carrier := &geometry.Vector{
		Point:     geometry.NewPoint(len(input[0])/2, len(input)/2),
		Direction: geometry.Up,
	}
	return grid, carrier
}

func visualize(grid map[geometry.Point]rune) {
	var top, right, bottom, left int
	for point := range grid {
		if point.Y < top {
			top = point.Y
		}
		if point.Y > bottom {
			bottom = point.Y
		}
		if point.X > right {
			right = point.X
		}
		if point.X < left {
			left = point.X
		}
	}

	numRows := ints.Abs(top) + ints.Abs(bottom) + 1
	grid2D := make([][]rune, numRows)
	numCols := ints.Abs(right) + ints.Abs(left) + 1
	for y := range grid2D {
		grid2D[y] = make([]rune, numCols)
	}

	for point, state := range grid {
		grid2D[point.Y-top][point.X-left] = state
	}

	fmt.Println(toString(grid2D), "")
}

func toString(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		for _, state := range row {
			visual := '.'
			if state == Infected {
				visual = '#'
			}
			sb.WriteRune(visual)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

const (
	Clean    = '.'
	Infected = '#'
	Flagged  = 'F'
	Weakened = 'W'
)
