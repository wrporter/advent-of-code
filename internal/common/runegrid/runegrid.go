package runegrid

import "github.com/wrporter/advent-of-code/internal/common/runes"

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
