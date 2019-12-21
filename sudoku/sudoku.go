package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/ints"
	"github.com/wrporter/advent-of-code-2019/internal/common/timeit"
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
	board, solved := Solve(puzzles[0])
	PrintBoard(board)
	fmt.Println(solved)
}

func PrintBoard(board [][]int) {
	for _, row := range board {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

func Solve(board [][]int) ([][]int, bool) {
	defer timeit.Track(time.Now(), "Sudoku")
	return backtrack(board)
}

func backtrack(board [][]int) ([][]int, bool) {
	row, col, exists := getNextEmptyCell(board)
	if !exists {
		return board, true
	}
	available := getAvailable(board, row, col)
	for _, number := range available {
		board[row][col] = number
		if solution, solved := backtrack(board); solved {
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

func getAvailable(board [][]int, row int, col int) []int {
	taken := make(map[int]bool)

	// Filter out numbers already in the row
	for _, value := range board[row] {
		taken[value] = true
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
