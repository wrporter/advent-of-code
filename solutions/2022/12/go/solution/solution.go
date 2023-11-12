package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	grid, start, end, _ := parseInput(input)
	return getMinSteps(grid, end, start)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	grid, _, end, aStarts := parseInput(input)
	return getMinSteps(grid, end, aStarts...)
}

func getMinSteps(grid [][]rune, end *geometry.Point, starts ...*geometry.Point) interface{} {
	queue := make([]node, len(starts))
	for i, start := range starts {
		queue[i] = node{point: start, steps: 0}
	}
	seen := make(map[geometry.Point]bool)
	current := queue[0]

	for !current.point.Equals(end) {
		current, queue = queue[0], queue[1:]

		if seen[*current.point] {
			continue
		}
		seen[*current.point] = true

		for _, direction := range geometry.Directions {
			next := current.point.Copy()
			next.Move(direction)

			if next.Y < len(grid) && next.Y >= 0 &&
				next.X >= 0 && next.X < len(grid[next.Y]) &&
				!seen[*next] &&
				grid[next.Y][next.X] <= grid[current.point.Y][current.point.X]+1 {
				queue = append(queue, node{
					point: next,
					steps: current.steps + 1,
				})
			}
		}
	}

	return current.steps
}

func parseInput(input string) (grid [][]rune, start *geometry.Point, end *geometry.Point, aStarts []*geometry.Point) {
	for y, line := range strings.Split(input, "\n") {
		grid = append(grid, nil)
		for x, char := range line {
			v := char
			if v == 'S' {
				v = 'a'
				start = geometry.NewPoint(x, y)
			} else if v == 'E' {
				v = 'z'
				end = geometry.NewPoint(x, y)
			}
			if v == 'a' {
				aStarts = append(aStarts, geometry.NewPoint(x, y))
			}
			grid[y] = append(grid[y], v)
		}
	}

	return grid, start, end, aStarts
}

type node struct {
	point *geometry.Point
	steps int
}
