package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	breakout2 "aoc/src/solutions/2019/day13/lib/breakout"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./2019/day13/input.txt")
	code, _ := convert.ToInts(strings.Split(lines[0], ","))
	b := breakout2.New(code)
	b.InsertQuarters(2)
	b.Play()
}
