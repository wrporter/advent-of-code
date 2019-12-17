package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day17/internal/scaffold"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./day17/input.txt")
	code, _ := conversion.ToInts(strings.Split(lines[0], ","))
	s := scaffold.New(code)
	fmt.Printf("Alignment parameter sum: %d\n", s.SumAlignmentIntersections())
}
