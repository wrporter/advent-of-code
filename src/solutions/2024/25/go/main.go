package main

import (
	"aoc/src/lib/go/aoc"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	var locks, keys []Schematic
	for _, schematicStr := range strings.Split(input, "\n\n") {
		schematic := strings.Split(schematicStr, "\n")
		if schematic[0][0] == '#' {
			locks = append(locks, newSchematic(schematic))
		} else {
			keys = append(keys, newSchematic(schematic))
		}
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if lock.fits(key) {
				count++
			}
		}
	}
	return count
}

func part2(_ string, _ ...interface{}) interface{} {
	return "Merry Christmas! 🎄"
}

func (s Schematic) fits(b Schematic) bool {
	for i := range len(b.heights) {
		if s.heights[i]+b.heights[i] > 5 {
			return false
		}
	}
	return true
}

func newSchematic(diagram []string) Schematic {
	heights := make([]int, len(diagram[0]))
	for _, row := range diagram {
		for x, cell := range row {
			if cell == '#' {
				heights[x] += 1
			}
		}
	}
	for i := range heights {
		heights[i]--
	}

	return Schematic{
		diagram: diagram,
		heights: heights,
	}
}

type Schematic struct {
	diagram []string
	heights []int
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 25, Part1: part1, Part2: part2}
}
