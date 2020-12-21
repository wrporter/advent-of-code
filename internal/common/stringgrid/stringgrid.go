package stringgrid

import (
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"strings"
)

func Flip(array []string) []string {
	result := make([]string, len(array))
	for i, str := range array {
		result[i] = mystrings.Reverse(str)
	}
	return result
}

func Rotate90Clockwise(array []string) []string {
	grid := conversion.ToRuneGrid(array)
	M := len(grid)
	N := len(grid[0])
	ret := make([][]rune, N)
	for r := 0; r < N; r++ {
		ret[r] = make([]rune, M)
	}

	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			ret[c][M-1-r] = grid[r][c]
		}
	}

	result := make([]string, len(ret))
	for i, row := range ret {
		result[i] = string(row)
	}

	return result
}

func GetCol(grid []string, col int) string {
	column := make([]uint8, len(grid))
	for row := 0; row < len(grid); row++ {
		column[row] = grid[row][col]
	}
	return string(column)
}

func Contains(grid []string, row string) bool {
	for _, r := range grid {
		if r == row {
			return true
		}
	}
	return false
}

func IndexOf(grid []string, row string) int {
	for i, r := range grid {
		if r == row {
			return i
		}
	}
	return -1
}

func Copy(array []string) []string {
	cpy := make([]string, len(array))
	copy(cpy, array)
	return cpy
}

func String(array []string) string {
	var sb strings.Builder
	for _, value := range array {
		sb.WriteString(value)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func Top(grid []string) string {
	return grid[0]
}

func Bottom(grid []string) string {
	return grid[len(grid)-1]
}

func Left(grid []string) string {
	return GetCol(grid, 0)
}

func Right(grid []string) string {
	return GetCol(grid, len(grid[0])-1)
}
