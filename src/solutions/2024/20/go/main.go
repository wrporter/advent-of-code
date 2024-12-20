package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, args ...interface{}) interface{} {
	return countCheats(input, args[0].(int), 2)
}

func part2(input string, args ...interface{}) interface{} {
	return countCheats(input, args[0].(int), 20)
}

func countCheats(input string, minimumDistance int, maxCheatDistance int) int {
	var start geometry.Point
	var path []int

	grid := strings.Split(input, "\n")
	height, width := len(grid), len(grid[0])

	for y, row := range grid {
		for x, cell := range row {
			if cell != '#' {
				path = append(path, y*height+x)
			}
			if cell == 'S' {
				start = *geometry.NewPoint(x, y)
			}
		}
	}

	distance := make([]int, height*width)
	seen := make([]bool, height*width)

	startIndex := toIndex(start, height)
	distance[startIndex] = 0
	seen[startIndex] = true

	queue := []geometry.Point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, d := range geometry.Directions {
			next := *current.Copy().Move(d)
			index := toIndex(next, height)

			if next.Y >= 0 && next.X >= 0 &&
				next.Y < height && next.X < width &&
				grid[next.Y][next.X] != '#' && !seen[index] {
				seen[index] = true
				distance[index] = distance[toIndex(current, height)] + 1
				queue = append(queue, next)
			}
		}
	}

	count := 0
	for _, i1 := range path {
		for _, i2 := range path {
			x1, y1 := i1%height, i1/height
			x2, y2 := i2%height, i2/height

			cheatDistance := ints.Abs(x1-x2) + ints.Abs(y1-y2)
			savedDistance := distance[i2] - distance[i1] - cheatDistance

			if cheatDistance <= maxCheatDistance && savedDistance >= minimumDistance {
				count++
			}
		}
	}
	return count
}

func toIndex(p geometry.Point, height int) int {
	return p.Y*height + p.X
}

func main() {
	New().Run([]interface{}{100}, []interface{}{100})
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 20, Part1: part1, Part2: part2}
}
