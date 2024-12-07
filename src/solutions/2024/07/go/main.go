package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"fmt"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	operators := []operatorFunc{add, multiply}
	return getTotalCalibrationResult(input, operators)
}

func part2(input string, _ ...interface{}) interface{} {
	operators := []operatorFunc{add, multiply, concatenate}
	return getTotalCalibrationResult(input, operators)
}

func getTotalCalibrationResult(input string, operators []operatorFunc) interface{} {
	equations := strings.Split(input, "\n")
	result := 0

	for _, equation := range equations {
		parts := strings.Split(equation, ": ")
		testValue := convert.StringToInt(parts[0])
		numbers := convert.ToIntsV2(strings.Fields(parts[1]))

		if isTrueEquation(operators, testValue, numbers[1:], numbers[0]) {
			result += testValue
		}
	}

	return result
}

type operatorFunc func(a, b int) int

func add(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }
func concatenate(a, b int) int {
	return convert.StringToInt(fmt.Sprintf("%d%d", a, b))
}

func isTrueEquation(operators []operatorFunc, target int, numbers []int, value int) bool {
	if len(numbers) == 0 {
		return value == target
	}

	for _, operation := range operators {
		if isTrueEquation(operators, target, numbers[1:], operation(value, numbers[0])) {
			return true
		}
	}
	return false
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 7, Part1: part1, Part2: part2}
}
