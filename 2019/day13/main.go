package main

import (
	breakout2 "github.com/wrporter/advent-of-code/2019/day13/internal/breakout"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./2019/day13/input.txt")
	code, _ := convert.ToInts(strings.Split(lines[0], ","))
	b := breakout2.New(code)
	b.InsertQuarters(2)
	b.Play()
}
