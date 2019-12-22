package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/ints"
	"github.com/wrporter/advent-of-code-2019/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	//puzzles := [][][]int{
	//	{
	//		{0, 0, 4, 3, 0, 0, 0, 0, 0},
	//		{0, 0, 1, 2, 0, 0, 0, 7, 0},
	//		{5, 7, 3, 9, 1, 0, 0, 0, 0},
	//		{6, 3, 5, 0, 0, 2, 0, 0, 0},
	//		{0, 0, 8, 0, 0, 0, 6, 0, 0},
	//		{0, 0, 0, 6, 0, 0, 4, 8, 5},
	//		{0, 0, 0, 0, 9, 8, 5, 2, 3},
	//		{0, 4, 0, 0, 0, 3, 9, 0, 0},
	//		{0, 0, 0, 0, 0, 7, 8, 0, 0},
	//	},
	//}
	puzzles := [][][]int{
		{
			{8, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 3, 6, 0, 0, 0, 0, 0},
			{0, 7, 0, 0, 9, 0, 2, 0, 0},
			{0, 5, 0, 0, 0, 7, 0, 0, 0},
			{0, 0, 0, 0, 4, 5, 7, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 3, 0},
			{0, 0, 1, 0, 0, 0, 0, 6, 8},
			{0, 0, 8, 5, 0, 0, 0, 1, 0},
			{0, 9, 0, 0, 0, 0, 4, 0, 0},
		},
	}
	Solve(puzzles[0])
}

func PrintBoard(board [][]int, original [][]int) {
	out := &strings.Builder{}
	out.WriteString("\033[2J")
	out.WriteString("\033[H")
	for y, row := range board {
		for x, value := range row {
			out.WriteString("\033[32m")
			if original[y][x] != 0 {
				out.WriteString("\033[31m")
			}
			out.WriteString(fmt.Sprintf("%d ", value))
		}
		out.WriteString("\n")
	}
	out.WriteString("\033[0m")
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 100)
}

func Solve(board [][]int) ([][]int, bool) {
	defer timeit.Track(time.Now(), "Sudoku")
	return backtrack(board, ints.Copy2D(board))
}

func backtrack(board [][]int, original [][]int) ([][]int, bool) {
	row, col, exists := getNextEmptyCell(board)
	if !exists {
		return board, true
	}

	for _, number := range getAvailableNumbers(board, row, col) {
		board[row][col] = number
		PrintBoard(board, original)

		if solution, solved := backtrack(board, original); solved {
			return solution, solved
		} else {
			board[row][col] = 0
		}
	}

	return board, false
}

func getNextEmptyCell(board [][]int) (int, int, bool) {
	for y, row := range board {
		for x, value := range row {
			if value == 0 {
				return y, x, true
			}
		}
	}
	return 0, 0, false
}

var Numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func getAvailableNumbers(board [][]int, row int, col int) []int {
	taken := make(map[int]bool)

	// Filter out numbers already in the row
	for _, number := range board[row] {
		taken[number] = true
	}

	// Filter out numbers already in the column
	for y := 0; y < len(board); y++ {
		taken[board[y][col]] = true
	}

	// Filter out numbers already in the 3x3 square
	size := ints.Sqrt(len(board))
	rowStart := row - (row % size)
	colStart := col - (col % size)
	for r := rowStart; r < rowStart+3; r++ {
		for c := colStart; c < colStart+3; c++ {
			taken[board[r][c]] = true
		}
	}

	return Filter(Numbers, func(i int) bool {
		return !taken[i]
	})
}

func Filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
