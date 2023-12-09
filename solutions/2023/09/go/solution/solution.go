package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	sequences := parseInput(input)
	return sumExtrapolatedValues(sequences, func(sequence []int, value int) int {
		return sequence[len(sequence)-1] + value
	})
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	sequences := parseInput(input)
	return sumExtrapolatedValues(sequences, func(sequence []int, value int) int {
		return sequence[0] - value
	})

}

func sumExtrapolatedValues(sequences [][]int, compute func(sequence []int, value int) int) int {
	sum := 0
	for _, sequence := range sequences {
		_, value := extrapolate(sequence, compute)
		sum += value
	}
	return sum
}

func extrapolate(sequence []int, compute func(sequence []int, value int) int) ([]int, int) {
	if isAllZero(sequence) {
		return sequence, 0
	}

	next := make([]int, len(sequence)-1)
	for i := 1; i < len(sequence); i++ {
		next[i-1] = sequence[i] - sequence[i-1]
	}

	_, value := extrapolate(next, compute)
	return next, compute(sequence, value)
}

func isAllZero(sequence []int) bool {
	for _, value := range sequence {
		if value != 0 {
			return false
		}
	}
	return true
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	sequences := make([][]int, len(lines))

	for i, line := range lines {
		sequences[i], _ = convert.ToInts(strings.Fields(line))
	}

	return sequences
}
