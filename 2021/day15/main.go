package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 15
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	cave := parseInput(input)
	lowestRiskLevel := findLowestRiskLevel(cave)
	return lowestRiskLevel
}

func findLowestRiskLevel(cave [][]int) int {
	visited := make(map[geometry.Point]bool)
	dist := make(map[geometry.Point]int)
	prev := make(map[geometry.Point]geometry.Point)
	start := geometry.NewPoint(0, 0)
	end := geometry.NewPoint(len(cave[0])-1, len(cave)-1)
	queue := []node{{Point: start, risk: 0}}

	for y, row := range cave {
		for x := range row {
			vertex := geometry.NewPoint(x, y)
			dist[vertex] = math.MaxInt
		}
	}
	dist[start] = 0

	var cur node
	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		if visited[cur.Point] {
			continue
		}

		for _, direction := range geometry.Directions {
			p := cur.Point.Move(direction)
			x, y := p.X, p.Y
			if visited[p] || y < 0 || y >= len(cave) || x < 0 || x >= len(cave[y]) {
				continue
			}

			risk := cave[p.Y][p.X]
			if dist[cur.Point]+risk < dist[p] {
				next := node{
					Point: p,
					risk:  dist[cur.Point] + risk,
				}
				dist[p] = dist[cur.Point] + risk
				prev[p] = cur.Point
				queue = append(queue, next)
			}
		}
	}

	return dist[end]
}

func part2(input []string) interface{} {
	return 0
}

type node struct {
	geometry.Point
	risk int
}

func parseInput(input []string) [][]int {
	cave := make([][]int, len(input))
	for i, line := range input {
		cave[i], _ = convert.ToInts(strings.Split(line, ""))
	}
	return cave
}
