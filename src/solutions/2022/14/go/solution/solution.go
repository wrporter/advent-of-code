package solution

import (
	"aoc/src/lib/go/geometry"
	"aoc/src/lib/go/ints"
	"math"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	source, scan, bottom := parseRockScan(input)
	hasEnteredVoid := func(unit geometry.Point) bool { return unit.Y > bottom }
	shouldFallTo := func(position geometry.Point) bool {
		_, exists := scan[position]
		return !exists
	}
	return sumRestedSand(scan, source, shouldFallTo, hasEnteredVoid)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	source, scan, bottom := parseRockScan(input)
	bottom += 2
	hasReachedSource := func(unit geometry.Point) bool { return unit == source }
	shouldFallTo := func(position geometry.Point) bool {
		_, exists := scan[position]
		return !exists && position.Y != bottom
	}
	return sumRestedSand(scan, source, shouldFallTo, hasReachedSource)
}

func parseRockScan(input string) (geometry.Point, map[geometry.Point]string, int) {
	rockPaths := strings.Split(input, "\n")
	source := geometry.Point{X: 500, Y: 0}
	scan := make(map[geometry.Point]string)
	bottom := math.MinInt

	for _, path := range rockPaths {
		points := geometry.ToPoints(strings.Split(path, " -> "))
		for i := 1; i < len(points); i++ {
			a := points[i-1]
			b := points[i]

			if a.X == b.X {
				for y := ints.Min(a.Y, b.Y); y <= ints.Max(a.Y, b.Y); y++ {
					scan[geometry.Point{X: a.X, Y: y}] = "#"
				}
			} else if a.Y == b.Y {
				for x := ints.Min(a.X, b.X); x <= ints.Max(a.X, b.X); x++ {
					scan[geometry.Point{X: x, Y: a.Y}] = "#"
				}
			}

			bottom = ints.Max(bottom, a.Y, b.Y)
		}
	}
	return source, scan, bottom
}

func sumRestedSand(
	scan map[geometry.Point]string,
	source geometry.Point,
	shouldFallTo func(position geometry.Point) bool,
	shouldExit func(unit geometry.Point) bool,
) int {
	exitCondition := false
	sumSandComeToRest := 0
	var next geometry.Point

	for !exitCondition {
		unit := source

		hasComeToRest := false
		for !hasComeToRest && !exitCondition {
			if next = unit.Down(); shouldFallTo(next) {
				unit = next
			} else if next = unit.DownLeft(); shouldFallTo(next) {
				unit = next
			} else if next = unit.DownRight(); shouldFallTo(next) {
				unit = next
			} else {
				scan[unit] = "o"
				sumSandComeToRest++
				hasComeToRest = true
			}

			if shouldExit(unit) {
				exitCondition = true
			}
		}
	}

	return sumSandComeToRest
}
