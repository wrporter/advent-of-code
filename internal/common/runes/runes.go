package runes

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

func GridToString(grid [][]rune) string {
	result := ""
	for _, row := range grid {
		for _, spot := range row {
			result += string(spot)
		}
		result += "\n"
	}
	return result
}
