package main

import (
	"fmt"
	droid2 "github.com/wrporter/advent-of-code/2019/day15/internal/droid"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./2019/day15/input.txt")
	code, _ := convert.ToInts(strings.Split(lines[0], ","))
	d := droid2.New(code)
	m, start, oxygen := d.ScanShip()
	fmt.Printf("Shortest path: %d\n", d.FindShortestPath(m, start, oxygen))
	fmt.Printf("Time to fill oxygen: %d\n", d.TimeToFillOxygen(m, oxygen))
}
