package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	start := geometry.NewPoint(1, 0)
	goal := geometry.NewPoint(len(grid[0])-2, len(grid)-1)
	seen := map[geometry.Point]bool{*geometry.NewPoint(1, -1): true}

	var dfs func(current *geometry.Point, steps int) int
	dfs = func(current *geometry.Point, steps int) int {
		longest := -1

		if current.Equals(goal) {
			return steps
		}

		for _, d := range geometry.Directions {
			next := current.Copy().Move(d)
			if seen[*next] {
				continue
			}
			c := grid[next.Y][next.X]
			if c == '.' ||
				(c == '<' && d == geometry.Left) ||
				(c == '>' && d == geometry.Right) ||
				(c == '^' && d == geometry.Up) ||
				(c == 'v' && d == geometry.Down) {
				seen[*next] = true
				longest = max(longest, dfs(next, steps+1))
				seen[*next] = false
			}
		}

		return longest
	}

	return dfs(start, 0)
}

func part2(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	start := geometry.NewPoint(1, 0)
	goal := geometry.NewPoint(len(grid[0])-2, len(grid)-1)

	graph := []Node{{id: 0}}
	ids := map[geometry.Point]int{*start: 0}

	seen := make(map[geometry.Point]bool)

	var collapse func(current *geometry.Point)
	collapse = func(current *geometry.Point) {
		for _, next := range getEdges(grid, current, current) {
			if seen[*next] {
				continue
			}

			seen[*next] = true
			prev := current
			distance := 1

			for edges := getEdges(grid, next, prev); len(edges) == 1; edges = getEdges(grid, next, prev) {
				distance += 1
				prev = next
				next = edges[0]
			}

			currentId := ids[*current]
			nextId, ok := ids[*next]
			if !ok {
				nextId = len(graph)
				graph = append(graph, Node{id: nextId})
			}

			graph[currentId].edges = append(graph[currentId].edges, Edge{nextId, distance})
			graph[nextId].edges = append(graph[nextId].edges, Edge{currentId, distance})

			ids[*next] = nextId
			seen[*prev] = true

			collapse(next)
		}
	}

	collapse(start)

	goalIndex := ids[*goal]
	seen2 := make(map[int]bool)

	var dfs func(current int, steps int) int
	dfs = func(current int, steps int) int {
		longest := -1

		if current == goalIndex {
			return steps
		}

		for _, edge := range graph[current].edges {
			next, distance := edge.id, edge.distance
			if !seen2[next] {
				seen2[next] = true
				longest = max(longest, dfs(next, steps+distance))
				seen2[next] = false
			}
		}

		return longest
	}

	return dfs(0, 0)
}

type Node struct {
	id    int
	edges []Edge
}

type Edge struct {
	id       int
	distance int
}

func getEdges(grid []string, current, prev *geometry.Point) []*geometry.Point {
	var edges []*geometry.Point
	for _, direction := range geometry.Directions {
		next := current.Copy().Move(direction)
		if !next.Equals(prev) &&
			next.X >= 0 && next.Y >= 0 &&
			next.Y < len(grid) && next.X < len(grid[next.Y]) &&
			grid[next.Y][next.X] != '#' {
			edges = append(edges, next)
		}
	}
	return edges
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 23, Part1: part1, Part2: part2}
}
