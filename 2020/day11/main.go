package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/runes"
)

func main() {
	year, day := 2020, 11
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	current := runes.Copy2D(conversion.ToRuneGrid(input))

	for {
		next := runes.Copy2D(current)

		for y, line := range next {
			for x, char := range line {
				switch char {
				case 'L':
					if numAdjacentOccupied(current, y, x) == 0 {
						next[y][x] = '#'
					}
				case '#':
					if numAdjacentOccupied(current, y, x) >= 4 {
						next[y][x] = 'L'
					}
				}
			}
		}

		if nextOccupied := numOccupied(next); nextOccupied == numOccupied(current) {
			return nextOccupied
		}
		current = next
	}
}

func part2(input []string) interface{} {
	current := runes.Copy2D(conversion.ToRuneGrid(input))

	for {
		next := runes.Copy2D(current)

		for y, line := range next {
			for x, char := range line {
				switch char {
				case 'L':
					if numSeenOccupied(current, y, x) == 0 {
						next[y][x] = '#'
					}
				case '#':
					if numSeenOccupied(current, y, x) >= 5 {
						next[y][x] = 'L'
					}
				}
			}
		}

		if nextOccupied := numOccupied(next); nextOccupied == numOccupied(current) {
			return nextOccupied
		}
		current = next
	}
}

func numSeenOccupied(grid [][]rune, row int, col int) int {
	count := 0
	for _, direction := range geometry.AllDirections {
		y := row + direction.Y
		x := col + direction.X

		for y >= 0 && y < len(grid) &&
			x >= 0 && x < len(grid[y]) {

			if grid[y][x] == 'L' {
				break
			}

			if grid[y][x] == '#' {
				count++
				break
			}

			y += direction.Y
			x += direction.X
		}
	}
	return count
}

func numAdjacentOccupied(grid [][]rune, row int, col int) int {
	count := 0
	for _, direction := range geometry.AllDirections {
		y := row + direction.Y
		x := col + direction.X

		if y >= 0 && y < len(grid) &&
			x >= 0 && x < len(grid[y]) &&
			grid[y][x] == '#' {
			count++
		}
	}
	return count
}

func numOccupied(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, spot := range row {
			if spot == '#' {
				count++
			}
		}
	}
	return count
}
