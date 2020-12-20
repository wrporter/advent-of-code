package stringgrid

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
