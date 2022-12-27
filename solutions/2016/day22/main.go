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

	year, day := 2016, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%$`)

func part1(input []string) interface{} {
	nodes, _ := parseInput(input)
	viable := identifyViablePairs(nodes)
	return len(viable)
}

func part2(input []string) interface{} {
	nodes, empty := parseInput(input)
	viable := identifyViablePairs(nodes)
	goal := getGoal(nodes)
	//display(nodes, viable)

	numStepsFromEmptyToGoal := getMinSteps(nodes, viable, empty, goal)

	// need to take 5 steps for every spot to shuffle the memory around
	// it is easier to calculate part 2 by doing this all by hand
	return numStepsFromEmptyToGoal + (5 * (goal.point.X - 1))
}

func getMinSteps(nodes map[geometry.Point]node, viable map[geometry.Point]bool, start node, target node) int {
	seen := make(map[geometry.Point]bool)
	queue := []bfs{{0, start}}
	var b bfs

	for len(queue) > 0 {
		b, queue = queue[0], queue[1:]
		if b.node == target {
			return b.steps
		}

		for _, dir := range geometry.Directions {
			p := b.node.point.Move(dir)
			if !seen[p] && viable[p] {
				seen[p] = true
				queue = append(queue, bfs{b.steps + 1, nodes[p]})
			}
		}
	}

	return -1
}

func identifyViablePairs(nodes map[geometry.Point]node) map[geometry.Point]bool {
	viable := make(map[geometry.Point]bool)
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1.isViablePair(n2) {
				viable[n1.point] = true
				viable[n2.point] = true
				//fmt.Printf("[a] used: %d | [b] used: %d\n", n1.used, n2.used)
			}
		}
	}
	return viable
}

func getGoal(nodes map[geometry.Point]node) node {
	maxX := 0
	for p := range nodes {
		maxX = ints.Max(maxX, p.X)
	}
	return nodes[geometry.NewPoint(maxX, 0)]
}

func parseInput(input []string) (nodes map[geometry.Point]node, empty node) {
	nodeInput := input[2:]
	nodes = make(map[geometry.Point]node, len(nodeInput))
	for _, line := range nodeInput {
		match := regex.FindStringSubmatch(line)
		values, _ := convert.ToInts(match[1:])

		point := geometry.NewPoint(values[0], values[1])
		n := node{
			point:      point,
			size:       values[2],
			used:       values[3],
			avail:      values[4],
			usePercent: values[5],
		}
		nodes[point] = n

		if n.used == 0 {
			empty = n
		}
	}
	return nodes, empty
}

type bfs struct {
	steps int
	node  node
}

type node struct {
	point      geometry.Point
	size       int
	used       int
	avail      int
	usePercent int
}

func (n *node) empty() bool {
	return n.used == 0
}

func (n *node) fits(b node) bool {
	return b.avail > n.used
}

func (n *node) isViablePair(b node) bool {
	return n.point != b.point && !n.empty() && n.fits(b)
}

func mapToGrid(m map[geometry.Point]node, viable map[geometry.Point]bool) [][]rune {
	width := 0
	height := 0
	for p := range m {
		width = ints.Max(width, p.X)
		height = ints.Max(height, p.Y)
	}
	width++
	height++

	grid := make([][]rune, height)

	for y := 0; y < height; y++ {
		row := make([]rune, width)
		grid[y] = row

		for x := 0; x < width; x++ {
			n := m[geometry.NewPoint(x, y)]
			if n.used == 0 {
				row[x] = '_'
			} else if n.point.X == width-1 && n.point.Y == 0 {
				row[x] = 'G'
			} else if n.point.X == 0 && n.point.Y == 0 {
				row[x] = 'R'
			} else if !viable[n.point] {
				row[x] = '#'
			} else {
				row[x] = '.'
			}
		}
	}

	return grid
}

func display(m map[geometry.Point]node, viable map[geometry.Point]bool) {
	b := &strings.Builder{}
	//b.WriteString("=====================================================\n")
	b.WriteString("\033c")
	b.WriteString(renderGrid(mapToGrid(m, viable)))
	fmt.Print(b.String())
	//time.Sleep(time.Millisecond * 20)
}

func renderGrid(grid [][]rune) string {
	result := ""
	for _, row := range grid {
		result += string(row)
		result += "\n"
	}
	return result
}
