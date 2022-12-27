package main

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	breakout2 "github.com/wrporter/advent-of-code/solutions/2019/day13/internal/breakout"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./2019/day13/input.txt")
	code, _ := convert.ToInts(strings.Split(lines[0], ","))
	b := breakout2.New(code)
	b.InsertQuarters(2)
	b.Play()
}
