package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"aoc/src/lib/go/v2/myslice"
	"fmt"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := strings.Split(input, "\n")
	start := geometry.NewPoint(1, 0)
	goal := geometry.NewPoint(len(grid[0])-2, len(grid)-1)
	seen := map[geometry.Point]bool{*geometry.NewPoint(1, -1): true}

	var dfs func(current *geometry.Point, steps int) int
	dfs = func(current *geometry.Point, steps int) int {
		longest := 0

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
	graph, ids := collapse(grid, start)

	//printMermaid(ids, graph)

	trimmedDistance := graph[ids[*start]].edges[0].distance + graph[ids[*goal]].edges[0].distance
	startId := trim(graph, ids, start)
	goalId := trim(graph, ids, goal)

	//return dfsRecursive(graph, goalId, startId, 0, 1<<startId) + trimmedDistance
	return dfsIterative(graph, goalId, startId) + trimmedDistance
}

func collapse(grid []string, start *geometry.Point) ([]Node, map[geometry.Point]int) {
	graph := []Node{{id: 0}}
	ids := map[geometry.Point]int{*start: 0}
	seen := make(map[geometry.Point]bool)
	queue := []*geometry.Point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

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

			queue = append(queue, next)
		}
	}

	return graph, ids
}

func trim(graph []Node, ids map[geometry.Point]int, current *geometry.Point) int {
	id := ids[*current]
	next := graph[ids[*current]].edges[0].id
	for i, edge := range graph[next].edges {
		if edge.id == id {
			graph[next].edges = myslice.Remove(graph[next].edges, i)
			break
		}
	}
	return next
}

func dfsRecursive(graph []Node, goalId, current, distance, seen int) int {
	if current == goalId {
		return distance
	}

	longest := 0

	for _, edge := range graph[current].edges {
		bit := 1 << edge.id
		if (seen & bit) == 0 {
			longest = max(longest, dfsRecursive(graph, goalId, edge.id, distance+edge.distance, seen|bit))
		}
	}

	return longest
}

func dfsIterative(graph []Node, goalId int, startId int) int {
	totalRemainingDistance := 0
	for _, node := range graph {
		for _, edge := range node.edges {
			totalRemainingDistance += edge.distance
		}
	}
	totalRemainingDistance /= 2

	stack := []Item{{startId, 0, 1 << startId, totalRemainingDistance}}
	longest := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.id == goalId {
			longest = max(longest, current.distance)
			continue
		}

		remainingDistance := current.remainingDistance
		for _, edge := range graph[current.id].edges {
			bit := 1 << edge.id
			if (current.seen & bit) == 0 {
				remainingDistance -= edge.distance
			}
		}

		for _, edge := range graph[current.id].edges {
			bit := 1 << edge.id
			distance := current.distance + edge.distance

			if current.seen&bit == 0 && remainingDistance+distance > longest {
				stack = append(stack, Item{
					id:                edge.id,
					distance:          distance,
					seen:              current.seen | bit,
					remainingDistance: remainingDistance,
				})
			}
		}
	}

	return longest
}

type Item struct {
	id                int
	distance          int
	seen              int
	remainingDistance int
}

func (i Item) Less(item aoc.PriorityQueueItem) bool {
	return i.distance > item.(Item).distance
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

func printMermaid(ids map[geometry.Point]int, graph []Node) {
	var sb strings.Builder
	sb.WriteString("graph LR")

	points := make([]geometry.Point, len(ids))
	for k, v := range ids {
		points[v] = k
	}
	seen := make(map[[2]int]bool)
	queue := []int{0}

	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		p := points[id]
		str := fmt.Sprintf("\n  %d[\"`(%d, %d)`\"]--", id, p.X, p.Y)

		for _, edge := range graph[id].edges {
			parts := []int{id, edge.id}
			sort.Ints(parts)
			key := [2]int{parts[0], parts[1]}

			if !seen[key] {
				ep := points[edge.id]
				sb.WriteString(fmt.Sprintf("%s %d ---%d[\"`(%d, %d)`\"]", str, edge.distance, edge.id, ep.X, ep.Y))
				queue = append(queue, edge.id)
			}
			seen[key] = true
		}
	}

	fmt.Println(sb.String())
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 23, Part1: part1, Part2: part2}
}
