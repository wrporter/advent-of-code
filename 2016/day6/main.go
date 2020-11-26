package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
)

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

func part1(input []string) string {
	grid := toGrid(input)
	message := make([]rune, len(grid[0]))

	row := 0
	for col := 0; row < len(grid) && col < len(grid[row]); col++ {
		charCounts := make(map[rune]int)
		mostFrequentCount := 0
		mostFrequentChar := '-'

		for _, line := range grid {
			char := line[col]
			if _, ok := charCounts[char]; !ok {
				charCounts[char] = 1
			} else {
				charCounts[char]++
			}

			if charCounts[char] > mostFrequentCount {
				mostFrequentCount = charCounts[char]
				mostFrequentChar = char
			}
		}

		message[col] = mostFrequentChar
		row++
	}

	return string(message)
}

func part2(input []string) string {
	grid := toGrid(input)
	message := make([]rune, len(grid[0]))

	row := 0
	for col := 0; row < len(grid) && col < len(grid[row]); col++ {
		charCounts := make(map[rune]int)

		for _, line := range grid {
			char := line[col]
			if _, ok := charCounts[char]; !ok {
				charCounts[char] = 1
			} else {
				charCounts[char]++
			}
		}

		leastFrequentCount := len(grid)
		leastFrequentChar := '-'
		for char, count := range charCounts {
			if count < leastFrequentCount {
				leastFrequentCount = count
				leastFrequentChar = char
			}
		}
		message[col] = leastFrequentChar

		row++
	}

	return string(message)
}

func main() {
	input, _ := file.ReadFile("./2016/day6/input.txt")
	answer1 := part1(input)
	answer2 := part2(input)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
