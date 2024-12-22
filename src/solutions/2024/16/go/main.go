package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/contain"
	"aoc/src/lib/go/v2/geometry"
	"math"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid, start := parse(input)
	best, _ := getBestPath(start, grid)
	return best
}

func part2(input string, _ ...interface{}) interface{} {
	grid, start := parse(input)
	_, sit := getBestPath(start, grid)
	return sit
}

func getBestPath(start geometry.Vector, grid []string) (int, int) {
	distance := make(map[geometry.Vector]int)
	queue := contain.NewPriorityQueue()
	queue.Push(Node{start, 0, []geometry.Point{start.Point}})
	best := math.MaxInt
	sit := make(map[geometry.Point]bool)

	for queue.Length() > 0 {
		current := queue.Pop().(Node)

		if grid[current.Point.Y][current.Point.X] == 'E' && current.score <= best {
			for _, position := range current.path {
				sit[position] = true
			}
			best = current.score
		}

		for _, option := range getNeighbors(current) {
			score, seen := distance[option.Vector]

			if grid[option.Point.Y][option.Point.X] != '#' && (!seen || option.score <= score) {
				distance[option.Vector] = option.score
				option.path = append(copyPath(current.path), option.Point)
				queue.Push(option)
			}
		}
	}

	return best, len(sit)
}

func parse(input string) ([]string, geometry.Vector) {
	grid := strings.Split(input, "\n")
	var start geometry.Vector
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				start = *geometry.NewVector(x, y, geometry.Right)
			}
		}
	}
	return grid, start
}

func getNeighbors(current Node) []Node {
	return []Node{{
		Vector: *current.Copy().Move(),
		score:  current.score + 1,
	}, {
		Vector: *current.Copy().Rotate(90).Move(),
		score:  current.score + 1001,
	}, {
		Vector: *current.Copy().Rotate(-90).Move(),
		score:  current.score + 1001,
	}}
}

func copyPath(path []geometry.Point) []geometry.Point {
	destination := make([]geometry.Point, len(path))
	for i, p := range path {
		destination[i] = *p.Copy()
	}
	return destination
}

type Node struct {
	geometry.Vector
	score int
	path  []geometry.Point
}

func (n Node) Less(item contain.PriorityQueueItem) bool {
	b := item.(Node)
	return n.score < b.score
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 16, Part1: part1, Part2: part2}
}
