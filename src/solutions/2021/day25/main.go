package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"aoc/src/lib/go/v2/myslice"
	"fmt"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 25
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	grid := convert.ToRuneGrid(input)

	for step := 1; step <= 1000; step++ {
		//fmt.Println(runegrid.String(grid))
		next := myslice.Copy2D(grid)
		madeMove := false

		for y, row := range grid {
			for x, spot := range row {
				if spot == '>' {
					nextX := (x + 1) % len(row)
					if row[nextX] == '.' {
						next[y][nextX] = spot
						next[y][x] = '.'
						madeMove = true
					}
				}
			}
		}

		grid = myslice.Copy2D(next)

		for y, row := range grid {
			for x, spot := range row {
				if spot == 'v' {
					nextY := (y + 1) % len(grid)
					if grid[nextY][x] == '.' {
						next[nextY][x] = spot
						next[y][x] = '.'
						madeMove = true
					}
				}
			}
		}

		if !madeMove {
			return step
		}
		grid = next
	}

	return 0
}

func part2(input []string) interface{} {
	return "Merry Christmas!!!"
}
