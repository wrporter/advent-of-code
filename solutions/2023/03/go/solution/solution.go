package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	sum := 0

	for row, line := range lines {
		start := -1
		var number string

		for col, char := range line {
			if isDigit(char) && start == -1 {
				start = col
			}

			if (col == len(line)-1 || !isDigit(rune(line[col+1]))) && start >= 0 {
				number = line[start : col+1]
				isPart := false

				for y := row - 1; y <= row+1 && !isPart; y++ {
					for x := start - 1; x <= col+1 && !isPart; x++ {
						if y >= 0 && y < len(lines) &&
							x >= 0 && x < len(lines[y]) &&
							isSymbol(rune(lines[y][x])) {
							value, _ := strconv.Atoi(number)
							sum += value
							isPart = true
						}
					}
				}

				debugNumber(isPart, start, col, line, row, lines)

				start = -1
			}
		}
	}

	return sum
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	sum := 0
	numbers := make(map[point]string)
	stars := make(map[point]bool)

	for row, line := range lines {
		start := -1
		var number string

		for col, char := range line {
			if char == '*' {
				stars[point{row, col}] = true
			}

			if isDigit(char) && start == -1 {
				start = col
			}

			if (col == len(line)-1 || !isDigit(rune(line[col+1]))) && start >= 0 {
				number = line[start : col+1]
				numbers[point{row, start}] = number
				start = -1
			}
		}
	}

	for star := range stars {
		var adjacent []string

		for num, number := range numbers {
			if star.row >= num.row-1 && star.row <= num.row+1 &&
				star.col >= num.col-1 && star.col <= num.col+len(number) {
				adjacent = append(adjacent, number)
			}
		}

		if len(adjacent) == 2 {
			part1, _ := strconv.Atoi(adjacent[0])
			part2, _ := strconv.Atoi(adjacent[1])
			sum += part1 * part2
		}
	}

	return sum
}

type point struct {
	row, col int
}

func isDigit(char rune) bool {
	return char <= '9' && char >= '0'
}

func isSymbol(char rune) bool {
	return !isDigit(char) && char != '.'
}

func debugNumber(isPart bool, start int, col int, line string, row int, lines []string) {
	if !isPart {
		chunk := ""
		colStart := start - 1
		if colStart < 0 {
			colStart = 0
		}
		colEnd := col + 2
		if colEnd > len(line) {
			colEnd = len(line)
		}

		if row > 0 {
			chunk += lines[row-1][colStart:colEnd] + "\n"
		}

		chunk += lines[row][colStart:colEnd] + "\n"

		if row+1 < len(line)-1 {
			chunk += lines[row+1][colStart:colEnd] + "\n"
		}

		fmt.Println(chunk)
	}
}
