package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
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
	phases := 0
	current := Copy2D(toGrid(input))

	for ; ; phases++ {
		next := Copy2D(current)
		for row, line := range next {
			for col, char := range line {
				switch char {
				case '.':
				case 'L':
					if numAdjacentOccupied(current, row, col) == 0 {
						next[row][col] = '#'
					}
				case '#':
					if numAdjacentOccupied(current, row, col) >= 4 {
						next[row][col] = 'L'
					}
				}
			}
		}

		if countOccupied(current) == countOccupied(next) {
			return countOccupied(next)
		}
		current = next
	}
}

func part2(input []string) interface{} {
	phases := 0
	current := Copy2D(toGrid(input))

	for ; ; phases++ {
		next := Copy2D(current)
		for row, line := range next {
			for col, char := range line {
				switch char {
				case '.':
				case 'L':
					if numAdjacentOccupiedSeen(current, row, col) == 0 {
						next[row][col] = '#'
					}
				case '#':
					if numAdjacentOccupiedSeen(current, row, col) >= 5 {
						next[row][col] = 'L'
					}
				}
			}
		}

		//fmt.Println(renderGrid(next))

		if countOccupied(current) == countOccupied(next) {
			return countOccupied(next)
		}
		current = next
	}
}

var directions = []geometry.Point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func renderGrid(grid [][]rune) string {
	result := ""
	for _, row := range grid {
		for _, spot := range row {
			result += string(spot)
		}
		result += "\n"
	}
	return result
}

func numAdjacentOccupiedSeen(input [][]rune, row int, col int) int {
	count := 0
	for _, direction := range directions {
		y, x := row+direction.Y, col+direction.X
		for y >= 0 && y < len(input) && x >= 0 && x < len(input[y]) && input[y][x] != 'L' {
			if input[y][x] == '#' {
				count++
				break
			}
			y += direction.Y
			x += direction.X
		}
	}
	return count
}

func numAdjacentOccupied(input [][]rune, row int, col int) int {
	count := 0
	for _, direction := range directions {
		y := row + direction.Y
		x := col + direction.X
		if y >= 0 && y < len(input) && x >= 0 && x < len(input[y]) && input[y][x] == '#' {
			count++
		}
	}
	return count
}

func countOccupied(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col == '#' {
				count++
			}
		}
	}
	return count
}

func toGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for row, line := range lines {
		grid[row] = make([]rune, len(line))

		for col, char := range line {
			grid[row][col] = char
		}
	}

	return grid
}

func Copy(array []rune) []rune {
	cpy := make([]rune, len(array))
	copy(cpy, array)
	return cpy
}

func Copy2D(grid [][]rune) [][]rune {
	cpy := make([][]rune, len(grid))
	for i := range grid {
		cpy[i] = Copy(grid[i])
	}
	return cpy
}
