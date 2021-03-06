package conversion

import "strconv"

func ToRuneGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for row, line := range lines {
		grid[row] = make([]rune, len(line))

		for col, char := range line {
			grid[row][col] = char
		}
	}

	return grid
}

func ToRunes(line string) []rune {
	row := make([]rune, len(line))

	for col, char := range line {
		row[col] = char
	}

	return row
}

func ToInts(values []string) (result []int, err error) {
	for _, stringValue := range values {
		value, err := strconv.ParseInt(stringValue, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, int(value))
	}
	return result, nil
}

func RuneToInt(rune uint8) int {
	return int(rune - '0')
}

func StringToInt(value string) int {
	valueInt64, _ := strconv.ParseInt(value, 10, 64)
	return int(valueInt64)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
