package main

import (
	"fmt"
	scaffold2 "github.com/wrporter/advent-of-code/2019/day17/internal/scaffold"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
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
