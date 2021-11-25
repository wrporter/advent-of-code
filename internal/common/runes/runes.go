package runes

import "github.com/wrporter/advent-of-code/internal/common/ints"

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

func StringToGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for row, line := range lines {
		grid[row] = make([]rune, len(line))

		for col, char := range line {
			grid[row][col] = char
		}
	}

	return grid
}

func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Reverse(arr []rune) []rune {
	result := Copy(arr)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func Rotate(values []rune, amount int) []rune {
	if len(values) == 0 {
		return values
	}

	rotation := len(values) - ints.WrapMod(amount, len(values))
	values = append(values[rotation:], values[:rotation]...)

	return values
}

func Concat(slices [][]rune) []rune {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]rune, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func Remove(slice []rune, index int) []rune {
	return append(slice[:index], slice[index+1:]...)
}

func Insert(a []rune, index int, value rune) []rune {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
