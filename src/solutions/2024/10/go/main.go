package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := convert.ToIntGrid(input)
	total := 0

	for y, row := range grid {
		for x, value := range row {
			if value == 0 {
				trailhead := *geometry.NewPoint(x, y)
				queue := []geometry.Point{trailhead}
				seen := map[geometry.Point]bool{trailhead: true}
				score := make(map[geometry.Point]bool)

				for len(queue) > 0 {
					current := queue[0]
					queue = queue[1:]

					seen[current] = true
					if grid[current.Y][current.X] == 9 {
						score[current] = true
						continue
					}

					for _, d := range geometry.Directions {
						next := *current.Copy().Move(d)
						if next.Y >= 0 && next.X >= 0 && next.Y < len(grid) &&
							next.X < len(grid[next.Y]) &&
							grid[current.Y][current.X]+1 == grid[next.Y][next.X] &&
							!seen[next] {
							queue = append(queue, next)
						}
					}
				}

				total += len(score)
			}
		}
	}

	return total
}

func part2(input string, _ ...interface{}) interface{} {
	grid := convert.ToIntGrid(input)
	total := 0

	for y, row := range grid {
		for x, value := range row {
			if value == 0 {
				total += getRating(grid, *geometry.NewPoint(x, y), make(map[geometry.Point]bool))
			}
		}
	}

	return total
}

func getRating(grid [][]int, current geometry.Point, seen map[geometry.Point]bool) int {
	rating := 0
	seen[current] = true
	defer func() { seen[current] = false }()

	if grid[current.Y][current.X] == 9 {
		return rating + 1
	}

	for _, d := range geometry.Directions {
		next := *current.Copy().Move(d)
		if next.Y >= 0 && next.X >= 0 && next.Y < len(grid) &&
			next.X < len(grid[next.Y]) &&
			grid[current.Y][current.X]+1 == grid[next.Y][next.X] &&
			!seen[next] {
			rating += getRating(grid, next, seen)
		}
	}

	return rating
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 10, Part1: part1, Part2: part2}
}
