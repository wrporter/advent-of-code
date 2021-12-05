package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 5
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

func part1(input []string) interface{} {
	plot := make(map[geometry.Point]int)

	for _, line := range input {
		x1, y1, x2, y2 := parseLine(line)
		if x1 == x2 || y1 == y2 {
			bresenhamPlot(x1, y1, x2, y2, plot)
		}
	}

	return countOverlappingPoints(plot)
}

func part2(input []string) interface{} {
	plot := make(map[geometry.Point]int)

	for _, line := range input {
		x1, y1, x2, y2 := parseLine(line)
		bresenhamPlot(x1, y1, x2, y2, plot)
	}

	return countOverlappingPoints(plot)
}

func countOverlappingPoints(points map[geometry.Point]int) interface{} {
	numOverlappingPoints := 0
	for _, overlap := range points {
		if overlap >= 2 {
			numOverlappingPoints++
		}
	}

	return numOverlappingPoints
}

func parseLine(in string) (int, int, int, int) {
	match := regex.FindStringSubmatch(in)
	nums, _ := convert.ToInts(match[1:])
	x1, y1 := nums[0], nums[1]
	x2, y2 := nums[2], nums[3]
	return x1, y1, x2, y2
}

// bresenhamPlot uses [Bresenham's Line Algorithm](https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm) to
// plot all integer points on a line.
func bresenhamPlot(x1, y1, x2, y2 int, plot map[geometry.Point]int) {
	dx := x2 - x1
	if dx < 0 {
		dx = -dx
	}

	dy := y2 - y1
	if dy < 0 {
		dy = -dy
	}

	var sx, sy int
	if x1 < x2 {
		sx = 1
	} else {
		sx = -1
	}

	if y1 < y2 {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	for {
		updatePlot(x1, y1, plot)

		if x1 == x2 && y1 == y2 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}

		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func updatePlot(x int, y int, plot map[geometry.Point]int) {
	point := geometry.NewPoint(x, y)
	if _, ok := plot[point]; ok {
		plot[point]++
	} else {
		plot[point] = 1
	}
}
