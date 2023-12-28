package main

import (
	"aoc/src/lib/go/aoc"
	"fmt"
	"strings"
)

func part1(input string, args ...interface{}) interface{} {
	cuts := args[0].(int)
	g := parse(strings.Split(input, "\n"))
	g1, g2, _ := g.split(cuts)
	return len(g1) * len(g2)
}

func part2(_ string, _ ...interface{}) interface{} {
	return "Merry Christmas! 🎄"
}

func (g Graph) split(cuts int) (Graph, Graph, map[string]bool) {
	var g1, g2 Graph
	from := g.getRandomNode()

	for to := range g {
		if from == to {
			continue
		}

		removed := make(map[string]bool)
		if g.isSplit(from, to, cuts, removed) {
			// Group reachable nodes
			g1, _ = g.findShortestPath(from, nil, removed)
			g2, _ = g.findShortestPath(to, nil, removed)

			// Make sure we didn't cut a node completely off the graph
			if len(g1)+len(g2) == len(g) {
				return g1, g2, removed
			}
		}
	}

	return g1, g2, nil
}

func (g Graph) isSplit(from, to *Node, cuts int, removed map[string]bool) bool {
	isSplit := false
	for i := 0; i <= cuts && !isSplit; i++ {
		_, isSplit = g.findShortestPath(from, to, removed)
	}
	return isSplit
}

func (g Graph) findShortestPath(from, to *Node, removed map[string]bool) (Graph, bool) {
	queue := []*Item{{node: from}}
	seen := make(Graph)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		seen[current.node] = true

		if current.node == to {
			// Remove this path
			for i := current; i.edge != ""; i = i.prev {
				removed[i.edge] = true
			}
			return seen, false
		}

		for edge, next := range current.node.edges {
			if !removed[edge] && !seen[next] {
				queue = append(queue, &Item{next, edge, current})
			}
		}
	}

	return seen, true
}

func (g Graph) getRandomNode() *Node {
	for n := range g {
		return n
	}
	return nil
}

func parse(input []string) Graph {
	nodes := make(map[string]*Node)
	graph := make(Graph)

	for _, line := range input {
		fromName, toNames, _ := strings.Cut(line, ": ")
		from := nodes[fromName]
		if from == nil {
			from = NewNode(fromName)
			graph[from] = true
			nodes[fromName] = from
		}
		for _, toName := range strings.Fields(toNames) {
			to := nodes[toName]
			if to == nil {
				to = NewNode(toName)
				graph[to] = true
				nodes[toName] = to
			}
			edge := fmt.Sprintf("%s-%s", fromName, toName)
			nodes[fromName].edges[edge] = to
			to.edges[edge] = nodes[fromName]
		}
	}

	return graph
}

type Node struct {
	name  string
	edges map[string]*Node
}

func NewNode(name string) *Node { return &Node{name: name, edges: make(map[string]*Node)} }

type Graph map[*Node]bool

type Item struct {
	node *Node
	edge string
	prev *Item
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 25, Part1: part1, Part2: part2}
}
