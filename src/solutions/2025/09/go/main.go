package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
	"aoc/src/lib/go/v2/mymath"
	"fmt"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	points := parse(input)

	largest := 0
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			width := mymath.Abs(point.X-points[j].X) + 1
			height := mymath.Abs(point.Y-points[j].Y) + 1
			area := width * height
			largest = max(largest, area)
		}
	}

	return largest
}

func part2(input string, _ ...interface{}) interface{} {
	points := parse(input)

	// Compress coordinates
	xset := make(map[int]bool)
	yset := make(map[int]bool)
	for _, pt := range points {
		xset[pt.X] = true
		xset[pt.X+1] = true
		yset[pt.Y] = true
		yset[pt.Y+1] = true
	}

	var xlst []int
	for x := range xset {
		xlst = append(xlst, x)
	}
	sort.Ints(xlst)

	var ylst []int
	for y := range yset {
		ylst = append(ylst, y)
	}
	sort.Ints(ylst)

	xmap := make(map[int]int)
	ymap := make(map[int]int)
	for i, x := range xlst {
		xmap[x] = i
	}
	for i, y := range ylst {
		ymap[y] = i
	}

	nX := len(xlst)
	nY := len(ylst)

	// grid initialization
	grid := make([][]int, nX)
	for i := range grid {
		grid[i] = make([]int, nY)
	}

	// Process line segments
	for i := 0; i < len(points); i++ {
		pt1 := points[i]
		pt2 := points[(i+1)%len(points)]

		x1, y1 := pt1.X, pt1.Y
		x2, y2 := pt2.X, pt2.Y

		mapX1, mapX2 := xmap[x1], xmap[x2]
		mapY1, mapY2 := ymap[y1], ymap[y2]

		if mapX1 != mapX2 { // Horizontal segment
			if mapY1 != mapY2 {
				// We expect only axis-aligned segments
				panic(fmt.Errorf("assertion failed: y1 (%d) != y2 (%d) for horizontal segment", mapY1, mapY2))
			}

			if mapX1 > mapX2 {
				mapX1, mapX2 = mapX2, mapX1
			}

			grid[mapX1][mapY1] |= 1
			grid[mapX2][mapY1] |= 2
			for x := mapX1 + 1; x < mapX2; x++ {
				grid[x][mapY1] |= 3
			}
		}
		// Vertical segments are ignored
	}

	// Parity check
	realGrid := make([][]bool, nX)
	for i := range realGrid {
		realGrid[i] = make([]bool, nY)
		next := 0
		for j, cell := range grid[i] {
			realGrid[i][j] = next > 0 || cell > 0
			next ^= cell
		}
	}

	// 2D prefix sum array
	realGridSums := make([][]int, nX+1)
	for i := range realGridSums {
		realGridSums[i] = make([]int, nY+1)
	}

	for i := 0; i < nX; i++ {
		for j := 0; j < nY; j++ {
			cellValue := 0
			if realGrid[i][j] {
				cellValue = 1
			}
			realGridSums[i+1][j+1] = cellValue + realGridSums[i][j+1] + realGridSums[i+1][j] - realGridSums[i][j]
		}
	}

	// Find largest fully covered rectangle
	largest := 0
	for i := 0; i < len(points); i++ {
		pt1 := points[i]
		for j := 0; j < i; j++ {
			pt2 := points[j]

			x1, y1 := pt1.X, pt1.Y
			x2, y2 := pt2.X, pt2.Y

			// Sort coordinates for bounding box
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			// Map to compressed indices
			xr1, yr1 := xmap[x1], ymap[y1]
			xr2, yr2 := xmap[x2], ymap[y2]

			// Calculate area of the compressed rectangle
			numGoods := realGridSums[xr2+1][yr2+1] - realGridSums[xr2+1][yr1] - realGridSums[xr1][yr2+1] + realGridSums[xr1][yr1]

			// Expected area
			expectedGoods := (xr2 - xr1 + 1) * (yr2 - yr1 + 1)

			// Check for full coverage
			if numGoods == expectedGoods {
				area := (x2 - x1 + 1) * (y2 - y1 + 1)
				largest = max(largest, area)
			}
		}
	}

	return largest
}

func parse(input string) []*geometry.Point {
	lines := strings.Split(input, "\n")
	points := make([]*geometry.Point, len(lines))
	for i, line := range lines {
		parts, _ := convert.ToInts(strings.Split(line, ","))
		points[i] = geometry.NewPoint(parts[0], parts[1])
	}
	return points
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 9, Part1: part1, Part2: part2}
}
