package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"regexp"
)

var regex = regexp.MustCompile(`^(\d+)x(\d+)x(\d+)$`)

func smallest(values []int) int {
	smallest := ints.MaxInt
	for _, value := range values {
		if value < smallest {
			smallest = value
		}
	}
	return smallest
}

func sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func calculate(presents []string) (paper int, ribbon int) {
	for _, present := range presents {
		match := regex.FindStringSubmatch(present)
		length, width, height := convert.StringToInt(match[1]), convert.StringToInt(match[2]), convert.StringToInt(match[3])

		areas := []int{length * width, width * height, height * length}
		slack := smallest(areas)
		paper += (2 * sum(areas)) + slack

		perimeters := []int{2 * (length + width), 2 * (width + height), 2 * (height + length)}
		bow := length * width * height
		ribbon += smallest(perimeters) + bow
	}

	return paper, ribbon
}

func main() {
	lines, _ := file.ReadFile("./2015/day2/input.txt")
	fmt.Println(calculate(lines))
}
