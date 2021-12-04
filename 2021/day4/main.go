package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 4
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numbers, boards := parseInput(input)
	winningScores := play(numbers, boards)
	return winningScores[0]
}

func part2(input []string) interface{} {
	numbers, boards := parseInput(input)
	winningScores := play(numbers, boards)
	return winningScores[len(winningScores)-1]
}

func play(numbers []int, boards map[int]*board) []int {
	var winningScores []int

	for _, number := range numbers {
		for i, b := range boards {
			b.mark(number)
			if b.won() {
				winningScores = append(winningScores, b.score(number))
				delete(boards, i)
			}
		}
	}

	return winningScores
}

func parseInput(input []string) ([]int, map[int]*board) {
	numbers, _ := conversion.ToInts(strings.Split(input[0], ","))

	boards := make(map[int]*board)
	grid := make([][]int, 0)
	boardInput := input[2:]

	for i, line := range boardInput {
		if line != "" {
			row, _ := conversion.ToInts(strings.Fields(line))
			grid = append(grid, row)
		}
		if line == "" || i == len(boardInput)-1 {
			boards[i] = newBoard(grid)
			grid = nil
		}
	}

	return numbers, boards
}

type board struct {
	numbers [][]int
	markers map[int]bool
}

func newBoard(numbers [][]int) *board {
	return &board{
		numbers: numbers,
		markers: make(map[int]bool),
	}
}

func (b *board) mark(number int) {
	b.markers[number] = true
}

func (b *board) won() bool {
	return b.aRowIsMarked() || b.aColumnIsMarked()
}

func (b *board) score(number int) int {
	sumUnmarkedNumbers := 0
	for _, row := range b.numbers {
		for _, unmarked := range row {
			if !b.markers[unmarked] {
				sumUnmarkedNumbers += unmarked
			}
		}
	}
	return sumUnmarkedNumbers * number
}

func (b *board) aRowIsMarked() bool {
	for _, row := range b.numbers {
		rowIsMarked := true
		for _, number := range row {
			if !b.markers[number] {
				rowIsMarked = false
			}
		}
		if rowIsMarked {
			return true
		}
	}
	return false
}

func (b *board) aColumnIsMarked() bool {
	for x := range b.numbers[0] {
		colIsMarked := true
		for y := range b.numbers {
			number := b.numbers[y][x]
			if !b.markers[number] {
				colIsMarked = false
			}
		}
		if colIsMarked {
			return true
		}
	}
	return false
}
