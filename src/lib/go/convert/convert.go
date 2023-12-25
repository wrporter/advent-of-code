package convert

import (
	"strconv"
	"strings"
)

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

func ToIntGrid(lines string) [][]int {
	rows := strings.Split(lines, "\n")
	grid := make([][]int, len(rows))
	for i, line := range rows {
		grid[i], _ = ToInts(strings.Split(line, ""))
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

func ToIntsV2(values []string) (result []int) {
	for _, stringValue := range values {
		value := StringToInt(stringValue)
		result = append(result, value)
	}
	return result
}

func ToFloats(values []string) []float64 {
	result := make([]float64, len(values))
	for i, stringValue := range values {
		value, _ := strconv.ParseFloat(stringValue, 64)
		result[i] = value
	}
	return result
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
