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
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 10
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>$`)

func part1(input []string) interface{} {
	vectors := parseInput(input)
	_, message := simulate(vectors)
	return message
}

func part2(input []string) interface{} {
	vectors := parseInput(input)
	t, _ := simulate(vectors)
	return t
}

func simulate(vectors map[vector]bool) (int, string) {
	for t := 1; t <= 20_000; t++ {
		next := make(map[vector]bool)
		points := make(map[geometry.Point]bool)

		for v := range vectors {
			v.X += v.velocity.X
			v.Y += v.velocity.Y
			next[v] = true
			points[v.Point] = true
		}
		vectors = next

		if every(points, isAdjacentToAnotherPoint(points)) {
			return t, "\n" + renderGrid(mapToGrid(vectors))
		}
	}

	return 0, ""
}

func isAdjacentToAnotherPoint(points map[geometry.Point]bool) func(p geometry.Point) bool {
	return func(p geometry.Point) bool {
		for _, direction := range geometry.AllDirections {
			neighbor := p.Add(direction)
			if points[neighbor] {
				return true
			}
		}
		return false
	}
}

func parseInput(input []string) map[vector]bool {
	vectors := make(map[vector]bool)
	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		point := geometry.NewPoint(convert.StringToInt(match[1]), convert.StringToInt(match[2]))
		velocity := geometry.NewPoint(convert.StringToInt(match[3]), convert.StringToInt(match[4]))
		v := vector{
			Point:    point,
			velocity: velocity,
		}
		vectors[v] = true
	}
	return vectors
}

type vector struct {
	geometry.Point
	velocity geometry.Point
}

func (v *vector) String() string {
	return fmt.Sprintf("p=<%d, %d> v=<%d, %d>", v.X, v.Y, v.velocity.X, v.velocity.Y)
}

func mapToGrid(vectors map[vector]bool) [][]rune {
	dim := &dimension{left: ints.MaxInt, top: ints.MaxInt}
	m := make(map[geometry.Point]bool)
	for v := range vectors {
		dim.left = ints.Min(dim.left, v.X)
		dim.right = ints.Max(dim.right, v.X)
		dim.top = ints.Min(dim.top, v.Y)
		dim.bottom = ints.Max(dim.bottom, v.Y)
		m[v.Point] = true
	}
	width, height := dim.width(), dim.height()

	grid := make([][]rune, height)
	for y := 0; y < height; y++ {
		row := make([]rune, width)
		grid[y] = row
		for x := 0; x < width; x++ {
			if m[geometry.NewPoint(x+dim.left, y+dim.top)] {
				row[x] = '#'
			} else {
				row[x] = ' '
			}
		}
	}

	return grid
}

func display(m map[vector]bool) {
	b := &strings.Builder{}
	//b.WriteString("=====================================================\n")
	b.WriteString("\033c")
	b.WriteString(renderGrid(mapToGrid(m)))
	fmt.Print(b.String())
	//time.Sleep(time.Millisecond * 20)
}

func renderGrid(grid [][]rune) string {
	result := ""
	delimiter := ""
	for _, row := range grid {
		result += delimiter
		result += string(row)
		delimiter = "\n"
	}
	return result
}

type dimension struct {
	left   int
	right  int
	top    int
	bottom int
}

func (d *dimension) height() int {
	return d.bottom - d.top + 1
}

func (d *dimension) width() int {
	return d.right - d.left + 1
}

func every(values map[geometry.Point]bool, test func(key geometry.Point) bool) bool {
	for key := range values {
		if !test(key) {
			return false
		}
	}
	return true
}
