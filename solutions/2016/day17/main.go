package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 17
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	passcode := input[0]
	return getShortestPath(passcode)
}

func part2(input []string) interface{} {
	passcode := input[0]
	return getLongestPathLength(passcode)
}

func getLongestPathLength(passcode string) int {
	current := node{passcode, "", geometry.NewPoint(0, 0)}
	queue := []node{current}
	vault := geometry.NewPoint(3, 3)
	max := 0

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, next := range getNeighbors(current) {
			if next.position == vault {
				max = ints.Max(max, len(next.path))
				continue
			}
			queue = append(queue, next)
		}
	}

	return max
}

func getShortestPath(passcode string) string {
	current := node{passcode, "", geometry.NewPoint(0, 0)}
	queue := []node{current}
	vault := geometry.NewPoint(3, 3)

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, next := range getNeighbors(current) {
			if next.position == vault {
				return next.path
			}
			queue = append(queue, next)
		}
	}

	return ""
}

var (
	directions = []rune{'U', 'D', 'L', 'R'}
	modifiers  = []geometry.Point{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
)

func getNeighbors(n node) []node {
	var neighbors []node
	sum := md5.Sum([]byte(n.passcode))
	hash := hex.EncodeToString(sum[:])
	for i, direction := range directions {
		position := n.position.Add(modifiers[i])
		if isOpen(hash[i]) && position.X >= 0 && position.X < 4 && position.Y >= 0 && position.Y < 4 {
			neighbors = append(neighbors, node{
				passcode: n.passcode + string(direction),
				path:     n.path + string(direction),
				position: position,
			})
		}
	}
	return neighbors
}

func isOpen(b byte) bool {
	return b >= 'b'
}

type node struct {
	passcode string
	path     string
	position geometry.Point
}
