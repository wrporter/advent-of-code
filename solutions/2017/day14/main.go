package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/knot"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strconv"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 14
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	key := input[0]
	numUsed := 0

	grid := toKnotHashBinaryGrid(key)

	for _, row := range grid {
		for _, value := range row {
			if value == '1' {
				numUsed++
			}
		}
	}

	return numUsed
}

func part2(input []string) interface{} {
	key := input[0]
	grid := toKnotHashBinaryGrid(key)
	return countRegions(grid)
}

func countRegions(grid []string) interface{} {
	visited := make(map[geometry.Point]bool)
	numRegions := 0
	var node geometry.Point

	for y, row := range grid {
		for x, value := range row {
			position := geometry.Point{X: x, Y: y}

			if !visited[position] && value == '1' {
				numRegions++
				queue := []geometry.Point{position}

				for len(queue) > 0 {
					node, queue = queue[0], queue[1:]
					if visited[node] {
						continue
					}
					visited[node] = true

					for _, direction := range geometry.Directions {
						next := node.Move(direction)
						if !visited[next] &&
							next.Y >= 0 && next.Y < len(grid) &&
							next.X >= 0 && next.X < len(grid[next.Y]) &&
							grid[next.Y][next.X] == '1' {
							queue = append(queue, next)
						}
					}
				}
			}
		}
	}

	return numRegions
}

func toKnotHashBinaryGrid(key string) []string {
	var grid []string

	for row := 0; row < 128; row++ {
		hex := knot.Hash(fmt.Sprintf("%s-%d", key, row))
		binary := longHexToBinary(hex)
		grid = append(grid, binary)
	}

	return grid
}

func longHexToBinary(hex string) string {
	result := ""
	for _, char := range hex {
		result += hexToBinary(string(char))
	}
	return result
}

func hexToBinary(hex string) string {
	ui, _ := strconv.ParseUint(hex, 16, 4)
	return fmt.Sprintf("%04b", ui)
}
