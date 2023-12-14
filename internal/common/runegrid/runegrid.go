package runegrid

import (
	"github.com/wrporter/advent-of-code/internal/common/runes"
	"strings"
)

// Rotate rotates the matrix by 90 degrees. For example, the following grid
//
// ```
// 1 2 3
// 4 5 6
// 7 8 9
// ```
//
// Would result in the following:
//
// ```
// 7 4 1
// 8 5 2
// 9 6 3
// ```
func Rotate(matrix [][]rune) [][]rune {
	grid := runes.Copy2D(matrix)
	size := len(grid)
	layerCount := size / 2

	for layer := 0; layer < layerCount; layer++ {

		first := layer
		last := size - first - 1

		for element := first; element < last; element++ {
			offset := element - first

			top := grid[first][element]
			right := grid[element][last]
			bottom := grid[last][last-offset]
			left := grid[last-offset][first]

			grid[first][element] = left
			grid[element][last] = top
			grid[last][last-offset] = right
			grid[last-offset][first] = bottom
		}
	}

	return grid
}

// RotateSwap does the same thing as Rotate, but rotates the grid in place rather than copying it
// to a new one.
func RotateSwap(grid [][]rune) {
	size := len(grid)

	for i := 0; i < size/2; i++ {
		for j := i; j < size-i-1; j++ {
			swap(&grid[i][j], &grid[size-j-1][i])
			swap(&grid[size-j-1][i], &grid[size-i-1][size-j-1])
			swap(&grid[size-i-1][size-j-1], &grid[j][size-i-1])
		}
	}
}

func swap[T interface{}](a, b *T) {
	*a, *b = *b, *a
}

func GetCol(grid [][]rune, col int) []rune {
	column := make([]rune, len(grid))
	for row := 0; row < len(grid); row++ {
		column[row] = grid[row][col]
	}
	return column
}

func Contains(grid [][]rune, row []rune) bool {
	for _, r := range grid {
		if runes.Equal(r, row) {
			return true
		}
	}
	return false
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

func String(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	return sb.String()
}
