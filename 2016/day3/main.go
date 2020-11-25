package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
)

var regex = regexp.MustCompile(`\s*(\d+)\s+(\d+)\s+(\d+)`)

func isValidTriangle(a, b, c int) bool {
	return a+b > c &&
		b+c > a &&
		a+c > b
}

func part1(input []string) int {
	numPossibleTriangles := 0

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		a, b, c := conversion.StringToInt(match[1]), conversion.StringToInt(match[2]), conversion.StringToInt(match[3])
		if isValidTriangle(a, b, c) {
			numPossibleTriangles++
		}
	}

	return numPossibleTriangles
}

func part2(input []string) int {
	numPossibleTriangles := 0
	triangles := make([][]int, 3)

	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		a, b, c := conversion.StringToInt(match[1]), conversion.StringToInt(match[2]), conversion.StringToInt(match[3])
		triangles[i%3] = []int{a, b, c}

		if (i+1)%3 == 0 && i != 0 {
			for triangleIndex := range triangles {
				a, b, c = triangles[0][triangleIndex], triangles[1][triangleIndex], triangles[2][triangleIndex]
				if isValidTriangle(a, b, c) {
					numPossibleTriangles++
				}
			}
		}
	}

	return numPossibleTriangles
}

func main() {
	input, _ := file.ReadFile("./2016/day3/input.txt")
	answer1 := part1(input)
	answer2 := part2(input)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
