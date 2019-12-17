package fft

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/math"
	"strings"
)

var pattern = []int{0, 1, 0, -1}

func Apply(signal string, phases int) string {
	output := parse(signal)
	next := make([]int, len(output))

	for phase := 0; phase < phases; phase++ {
		for row := 0; row < len(output); row++ {
			sum := 0
			for col := 0; col < len(output); col++ {
				inputDigit := output[col]
				coef := coefficient(row+1, col+1)
				sum += inputDigit * coef
			}
			next[row] = math.Abs(sum) % 10
		}

		output, next = next, output
	}

	return toString(output)
}

func coefficient(row, col int) int {
	return pattern[(col/row)%4]
}

func parse(signal string) []int {
	digits := make([]int, len(signal))
	digitStrings := strings.Split(signal, "")
	for i, digitString := range digitStrings {
		digits[i] = conversion.StringToInt(digitString)
	}
	return digits
}

func toString(signal []int) string {
	return strings.Trim(
		strings.Join(
			strings.Split(fmt.Sprint(signal), " "),
			"",
		),
		"[]",
	)
}
