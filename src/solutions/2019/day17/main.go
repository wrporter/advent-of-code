package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	scaffold2 "aoc/src/solutions/2019/day17/lib/scaffold"
	"fmt"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./2019/day17/input.txt")
	code, _ := convert.ToInts(strings.Split(lines[0], ","))

	s := scaffold2.New(code)
	grid, robot := s.Scan()

	fmt.Printf("Alignment parameter sum: %d\n", s.SumAlignmentIntersections(grid))
	fmt.Printf("Dust collected: %d\n", s.MoveRobot(grid, robot))
}
