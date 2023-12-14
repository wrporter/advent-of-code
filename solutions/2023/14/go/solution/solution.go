package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/runegrid"
	"github.com/wrporter/advent-of-code/internal/common/v2/myslice"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	rollRocks(grid)
	return calculateNorthLoad(grid)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	seen := make(map[string]int)
	cycles := 1_000_000_000
	cycle := 1

	for ; cycle <= cycles; cycle++ {
		for tilt := 1; tilt <= 4; tilt++ {
			rollRocks(grid)
			myslice.Rotate90Degrees(grid)
		}

		str := runegrid.String(grid)
		if start, ok := seen[str]; ok {
			period := cycle - start
			remaining := cycles - cycle
			jump := (remaining / period) * period
			cycle += jump
		}

		seen[str] = cycle
	}

	return calculateNorthLoad(grid)
}

func rollRocks(grid [][]rune) {
	for x := 0; x < len(grid[0]); x++ {
		empty := 0

		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 'O' {
				grid[y][x] = '.'
				grid[empty][x] = 'O'
				empty += 1
			} else if grid[y][x] == '#' {
				empty = y + 1
			}
		}
	}
}

func calculateNorthLoad(grid [][]rune) interface{} {
	sum := 0
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'O' {
				load := len(grid) - y
				sum += load
			}
		}
	}
	return sum
}
