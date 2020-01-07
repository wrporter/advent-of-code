package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"regexp"
)

var regex = regexp.MustCompile(`-?\d+`)

func sum(json string) int {
	values := regex.FindAllString(json, -1)
	numbers, _ := conversion.ToInts(values)
	return ints.Sum(numbers)
}

func main() {
	lines, _ := file.ReadFile("./2015/day12/input.txt")
	fmt.Println(sum(lines[0]))
}
