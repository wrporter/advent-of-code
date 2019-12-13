package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day13/internal/breakout"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./day13/input.txt")
	code, _ := conversion.ToInts(strings.Split(lines[0], ","))
	b := breakout.New(code)
	fmt.Println(b.Play())
}
