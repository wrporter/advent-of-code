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

func part2(input []string) interface{} {
	cave := parseInput(input)
	cave = expand(cave)
	lowestRiskLevel := findLowestRiskLevel(cave)
	return lowestRiskLevel
}

func expand(cave [][]int) [][]int {
	size := 5
	result := make([][]int, len(cave)*size)
	for y := range result {
		result[y] = make([]int, len(cave[0])*size)
		for x := range result[y] {
			risk := cave[y%len(cave)][x%len(cave[0])] + (y / len(cave)) + (x / len(cave[0]))
			increase := risk / 10
			risk -= increase * 9
			result[y][x] = risk
		}
	}
	return result
}

func findLowestRiskLevel(cave [][]int) int {
	dist := make(map[geometry.Point]int)
	prev := make(map[geometry.Point]geometry.Point)
	source := geometry.NewPoint(0, 0)
	target := geometry.NewPoint(len(cave[0])-1, len(cave)-1)
	queue := []geometry.Point{source}

	for y, row := range cave {
		for x := range row {
			vertex := geometry.NewPoint(x, y)
			dist[vertex] = math.MaxInt
		}
	}
	dist[source] = 0

	var u geometry.Point
	for len(queue) > 0 {
		u, queue = queue[0], queue[1:]

		for _, direction := range geometry.Directions {
			v := u.Move(direction)
			x, y := v.X, v.Y
			if y < 0 || y >= len(cave) || x < 0 || x >= len(cave[y]) {
				continue
			}

			alt := dist[u] + cave[y][x]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				queue = append(queue, v)
			}
		}
	}

	return dist[target]
}

func parseInput(input []string) [][]int {
	cave := make([][]int, len(input))
	for i, line := range input {
		cave[i], _ = convert.ToInts(strings.Split(line, ""))
	}
	return cave
}
