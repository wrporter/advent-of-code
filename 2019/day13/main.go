package main

import (
	breakout2 "github.com/wrporter/advent-of-code-2019/2019/day13/internal/breakout"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./day13/input.txt")
	code, _ := conversion.ToInts(strings.Split(lines[0], ","))
	b := breakout2.New(code)
	b.InsertQuarters(2)
	b.Play()
}
