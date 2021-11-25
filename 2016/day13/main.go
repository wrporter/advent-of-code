package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math/bits"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 13
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	favorite, target := parseInput(input)
	start := geometry.NewPoint(1, 1)
	queue := []Node{{start, 0}}
	visited := map[geometry.Point]bool{start: true}
	var current Node

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, next := range getNeighbors(current.Point, favorite) {
			if next == target {
				return current.Steps + 1
			}

			if !visited[next] {
				visited[next] = true
				queue = append(queue, Node{next, current.Steps + 1})
			}
		}
	}

	return -1
}

func part2(input []string) interface{} {
	favorite, _ := parseInput(input)
	start := geometry.NewPoint(1, 1)
	queue := []Node{{start, 0}}
	visited := map[geometry.Point]bool{start: true}
	var current Node
	numOpenSpacesIn50Steps := 0

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		numOpenSpacesIn50Steps++

		for _, next := range getNeighbors(current.Point, favorite) {
			if !visited[next] && current.Steps < 50 {
				visited[next] = true
				queue = append(queue, Node{next, current.Steps + 1})
			}
		}
	}

	return numOpenSpacesIn50Steps
}

func getNeighbors(point geometry.Point, favorite int) []geometry.Point {
	var neighbors []geometry.Point
	for _, direction := range geometry.Directions {
		neighbor := point.Move(direction)
		if neighbor.X >= 0 && neighbor.Y >= 0 && isOpenSpace(neighbor, favorite) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func isOpenSpace(point geometry.Point, favorite int) bool {
	x, y := point.X, point.Y
	value := x*x + 3*x + 2*x*y + y + y*y
	value += favorite
	numOnes := bits.OnesCount(uint(value))
	return numOnes%2 == 0
}

type Node struct {
	Point geometry.Point
	Steps int
}

func parseInput(input []string) (int, geometry.Point) {
	favoriteNumber := conversion.StringToInt(input[0])
	coordinates := strings.Split(input[1], ",")
	target := geometry.NewPoint(conversion.StringToInt(coordinates[0]), conversion.StringToInt(coordinates[1]))
	return favoriteNumber, target
}
