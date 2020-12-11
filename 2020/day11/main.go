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
	phases := 0
	current := runes.Copy2D(conversion.ToRunes(input))

	for ; ; phases++ {
		next := runes.Copy2D(current)
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
	current := runes.Copy2D(conversion.ToRunes(input))

	for ; ; phases++ {
		next := runes.Copy2D(current)
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

		//fmt.Println(runes.GridToString(next))

		if countOccupied(current) == countOccupied(next) {
			return countOccupied(next)
		}
		current = next
	}
}

func numAdjacentOccupiedSeen(input [][]rune, row int, col int) int {
	count := 0
	for _, direction := range geometry.AllDirections {
		y := row + direction.Y
		x := col + direction.X

		for y >= 0 && y < len(input) &&
			x >= 0 && x < len(input[y]) &&
			input[y][x] != 'L' {

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
	for _, direction := range geometry.AllDirections {
		y := row + direction.Y
		x := col + direction.X

		if y >= 0 && y < len(input) &&
			x >= 0 && x < len(input[y]) &&
			input[y][x] == '#' {
			count++
		}
	}
	return count
}

func countOccupied(grid [][]rune) int {
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
