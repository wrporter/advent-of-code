package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, args ...interface{}) interface{} {
	size, exit, start, byteList := parse(input, args)
	numBytes := args[1].(int)

	bytes := make(map[geometry.Point]bool)
	for i := 0; i < len(byteList) && i < numBytes; i++ {
		bytes[*byteList[i]] = true
	}

	return getShortestPath(start, exit, size, bytes)
}

func part2(input string, args ...interface{}) interface{} {
	size, exit, start, byteList := parse(input, args)

	bytes := make(map[geometry.Point]bool)
	for _, b := range byteList {
		bytes[*b] = true
		steps := getShortestPath(start, exit, size, bytes)
		if steps == -1 {
			return b.String()
		}
	}

	return -1
}

func parse(input string, args []interface{}) (int, *geometry.Point, *geometry.Point, []*geometry.Point) {
	size := args[0].(int)
	exit := geometry.NewPoint(size, size)
	start := geometry.NewPoint(0, 0)
	byteList := geometry.ToPoints(strings.Split(input, "\n"))
	return size, exit, start, byteList
}

func getShortestPath(start *geometry.Point, exit *geometry.Point, size int, bytes map[geometry.Point]bool) interface{} {
	queue := []Node{{*start, 0}}
	seen := make(map[geometry.Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Equals(exit) {
			return current.steps
		}

		for _, d := range geometry.Directions {
			next := *current.Copy().Move(d)
			if next.Y >= 0 && next.X >= 0 &&
				next.Y <= size && next.X <= size &&
				!bytes[next] && !seen[next] {
				seen[next] = true
				queue = append(queue, Node{next, current.steps + 1})
			}
		}
	}

	return -1
}

type Node struct {
	geometry.Point
	steps int
}

func main() {
	New().Run([]interface{}{70, 1024}, []interface{}{70})
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 18, Part1: part1, Part2: part2}
}
