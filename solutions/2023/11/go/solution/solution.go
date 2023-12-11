package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	return getTotalDistance(input, 2)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	return getTotalDistance(input, 1_000_000)
}

func getTotalDistance(input string, gapSize int) interface{} {
	galaxies := parseInput(input, gapSize-1)
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].ManhattanDistance(galaxies[j])
		}
	}
	return sum
}

func parseInput(input string, addGap int) []*geometry.Point {
	image := strings.Split(input, "\n")

	yGaps := make([]int, len(image))
	yGap := 0
	for y := 0; y < len(image); y++ {
		gap := addGap
		for x := 0; x < len(image[0]); x++ {
			if image[y][x] == '#' {
				gap = 0
			}
		}

		yGap += gap
		yGaps[y] = yGap
	}

	xGaps := make([]int, len(image[0]))
	xGap := 0
	for x := 0; x < len(image[0]); x++ {
		gap := addGap
		for y := 0; y < len(image); y++ {
			if image[y][x] == '#' {
				gap = 0
			}
		}

		xGap += gap
		xGaps[x] = xGap
	}

	var galaxies []*geometry.Point

	for y, line := range image {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, geometry.NewPoint(x+xGaps[x], y+yGaps[y]))
			}
		}
	}

	return galaxies
}
