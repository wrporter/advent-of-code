package main

import (
	"aoc/src/lib/go/bytes"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 6
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	coordinates, dim := parseInput(input)
	areas := make(map[coordinate]map[geometry.Point]bool)

	// calculate area for each coordinate
	for y := dim.top; y <= dim.bottom; y++ {
		for x := dim.left; x <= dim.right; x++ {
			point := geometry.NewPoint(x, y)
			closest := getClosest(coordinates, point)

			if closest != nil {
				if areas[*closest] == nil {
					areas[*closest] = make(map[geometry.Point]bool)
				}
				areas[*closest][point] = true
			}
		}
	}

	//display(areas, dim)

	// identify areas that are infinite that touch the edge
	infinite := make(map[coordinate]bool)
	for coord, area := range areas {
		for point := range area {
			if point.X == dim.left ||
				point.X == dim.right ||
				point.Y == dim.top ||
				point.Y == dim.bottom {
				infinite[coord] = true
			}
		}
	}

	// get the max area that is not infinite
	maxArea := 0
	for coord, area := range areas {
		if !infinite[coord] {
			maxArea = ints.Max(maxArea, len(area))
		}
		//fmt.Printf("%s = %d\n", coord.name, len(area))
	}

	return maxArea
}

func part2(input []string) interface{} {
	coordinates, dim := parseInput(input)
	safe := make(map[geometry.Point]bool)
	total := 10_000
	dim.expand(100)

	for y := dim.top; y <= dim.bottom; y++ {
		for x := dim.left; x <= dim.right; x++ {
			point := geometry.NewPoint(x, y)
			distance := getTotalDistance(coordinates, point)
			if distance < total {
				safe[point] = true
			}
		}
	}

	return len(safe)
}

func getTotalDistance(coordinates map[coordinate]bool, point geometry.Point) int {
	total := 0
	for coord := range coordinates {
		total += coord.ManhattanDistance(point)
	}
	return total
}

func getClosest(coords map[coordinate]bool, point geometry.Point) (closest *coordinate) {
	min := ints.MaxInt
	for coord := range coords {
		distance := coord.ManhattanDistance(point)
		if distance < min {
			min = distance
			closest = &coordinate{
				name:  coord.name,
				Point: geometry.NewPoint(coord.X, coord.Y),
			}
		}
	}
	for coord := range coords {
		distance := coord.ManhattanDistance(point)
		if distance == min && coord != *closest {
			return nil
		}

	}
	return closest
}

func parseInput(input []string) (points map[coordinate]bool, dim *dimension) {
	points = make(map[coordinate]bool)
	dim = &dimension{left: ints.MaxInt, top: ints.MaxInt}
	name := "A"
	for _, line := range input {
		parts := strings.Split(line, ", ")
		x := convert.StringToInt(parts[0])
		y := convert.StringToInt(parts[1])

		coord := coordinate{
			name:  name,
			Point: geometry.NewPoint(x, y),
		}
		points[coord] = true
		name = string(name[0] + 1)

		dim.left = ints.Min(dim.left, x)
		dim.right = ints.Max(dim.left, x)
		dim.top = ints.Min(dim.top, y)
		dim.bottom = ints.Max(dim.bottom, y)
	}
	dim.expand(1)
	return points, dim
}

type dimension struct {
	left   int
	right  int
	top    int
	bottom int
}

func (d *dimension) expand(amount int) {
	d.left -= amount
	d.right += amount
	d.top -= amount
	d.bottom += amount
}

func (d *dimension) height() int {
	return d.bottom - d.top + 1
}

func (d *dimension) width() int {
	return d.right - d.left + 1
}

type coordinate struct {
	name string
	geometry.Point
}

func mapToGrid(m map[coordinate]map[geometry.Point]bool, dim *dimension) [][]rune {
	height, width := dim.height(), dim.width()
	grid := make([][]rune, height)

	for y := 0; y < height; y++ {
		row := make([]rune, width)
		grid[y] = row
		for x := 0; x < width; x++ {
			row[x] = '.'
		}
	}

	for coord, points := range m {
		for point := range points {
			grid[point.Y-dim.top][point.X-dim.left] = rune(bytes.ToLower(coord.name[0]))
		}
		grid[coord.Y-dim.top][coord.X-dim.left] = rune(coord.name[0])
	}

	return grid
}

func display(m map[coordinate]map[geometry.Point]bool, dim *dimension) {
	b := &strings.Builder{}
	//b.WriteString("=====================================================\n")
	b.WriteString("\033c")
	b.WriteString(renderGrid(mapToGrid(m, dim)))
	fmt.Print(b.String())
	//time.Sleep(time.Millisecond * 20)
}

func renderGrid(grid [][]rune) string {
	b := &strings.Builder{}
	for _, row := range grid {
		b.WriteString(string(row))
		b.WriteString("\n")
	}
	return b.String()
}
