package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strconv"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 20
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	algorithm, image := parseInput(input)
	image = enhance(image, algorithm, 2)
	numLit := countLit(image)
	return numLit
}

func part2(input []string) interface{} {
	algorithm, image := parseInput(input)
	image = enhance(image, algorithm, 50)
	numLit := countLit(image)
	return numLit
}

func enhance(image []string, algorithm string, numSteps int) []string {
	lit0 := algorithm[0] == '#'
	lit512 := algorithm[511] == '#'

	for step := 0; step < numSteps; step++ {
		lit := lit0
		if lit0 && !lit512 && step%2 == 0 {
			lit = lit512
		}
		image = enhanceStep(image, algorithm, lit)
	}

	return image
}

func enhanceStep(image []string, algorithm string, lit bool) []string {
	enhanced := make([]string, len(image)+2)

	for y := 0; y < len(enhanced); y++ {
		for x := 0; x < len(image[0])+2; x++ {
			pixels := ""

			for dy := y - 1; dy <= y+1; dy++ {
				for dx := x - 1; dx <= x+1; dx++ {
					if dy < 1 || dy > len(image) || dx < 1 || dx > len(image[0]) {
						if lit {
							pixels += "1"
						} else {
							pixels += "0"
						}
					} else if image[dy-1][dx-1] == '.' {
						pixels += "0"
					} else {
						pixels += "1"
					}
				}
			}

			index := binaryToInt(pixels)
			enhanced[y] += string(algorithm[index])
		}
	}

	return enhanced
}

func countLit(image []string) int {
	result := 0
	for _, row := range image {
		for _, char := range row {
			if char == '#' {
				result++
			}
		}
	}
	return result
}

func parseInput(input []string) (string, []string) {
	algorithm := input[0]
	image := input[2:]
	return algorithm, image
}

func binaryToInt(binary string) int {
	value, _ := strconv.ParseInt(binary, 2, 32)
	return int(value)
}
