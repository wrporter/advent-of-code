package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/ints"
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
func bresenhamPlot(x0, y0, x1, y1 int, plot map[geometry.Point]int) {
	dx := ints.Abs(x1 - x0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}

	dy := -ints.Abs(y1 - y0)
	sy := -1
	if y0 < y1 {
		sy = 1
	}

	err := dx + dy

	for {
		updatePlot(x0, y0, plot)

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
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
