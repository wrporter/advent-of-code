package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
	"aoc/src/lib/go/v2/myslice"
)

func part1(input string, _ ...interface{}) interface{} {
	grid := convert.ToIntGrid(input)
	return getMinHeatLoss(grid, 1, 3)
}

func part2(input string, _ ...interface{}) interface{} {
	grid := convert.ToIntGrid(input)
	return getMinHeatLoss(grid, 4, 10)
}

func getMinHeatLoss(grid [][]int, minDistance int, maxDistance int) int {
	queue := aoc.NewPriorityQueue()
	goal := geometry.NewPoint(len(grid[0])-1, len(grid)-1)
	queue.Push(&node{
		vector: *geometry.NewVector(0, 0, geometry.Right),
		cost:   0,
	})
	queue.Push(&node{
		vector: *geometry.NewVector(0, 0, geometry.Down),
		cost:   0,
	})
	seen := make(map[geometry.Vector]bool)
	cost := make(map[geometry.Vector]int)

	for queue.Length() > 0 {
		current := queue.Pop().(*node)

		if current.vector.Point.Equals(goal) {
			return current.cost
		}

		if seen[current.vector] {
			continue
		}
		seen[current.vector] = true

		directions := []geometry.Direction{
			current.vector.Direction.Rotate(-90),
			current.vector.Direction.Rotate(90),
		}
		for _, direction := range directions {
			costIncrease := 0

			for distance := 1; distance <= maxDistance; distance++ {
				delta := geometry.AllDirectionsModifiers[direction-1]
				x := current.vector.Point.X + delta.X*distance
				y := current.vector.Point.Y + delta.Y*distance

				if !myslice.InBounds(grid, y, x) {
					break
				}

				costIncrease += grid[y][x]
				vector := *geometry.NewVector(x, y, direction)
				nextCost := current.cost + costIncrease

				if c, ok := cost[vector]; (ok && c <= nextCost) || distance < minDistance {
					continue
				}
				cost[vector] = nextCost

				queue.Push(&node{
					vector: vector,
					cost:   nextCost,
				})
			}
		}

	}

	return 0
}

type node struct {
	vector geometry.Vector
	cost   int
}

func (n *node) Less(item aoc.PriorityQueueItem) bool {
	return n.cost < item.(*node).cost
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 17, Part1: part1, Part2: part2}
}
