package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day15/internal/droid"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	lines, _ := file.ReadFile("./day15/input.txt")
	code, _ := conversion.ToInts(strings.Split(lines[0], ","))
	d := droid.New(code)
	m, start, oxygen := d.ScanShip()
	fmt.Printf("Shortest path: %d\n", d.FindShortestPath(m, start, oxygen))
	fmt.Printf("Time to fill oxygen: %d\n", d.TimeToFillOxygen(m, oxygen))
}
