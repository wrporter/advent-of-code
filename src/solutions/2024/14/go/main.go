package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/v2/geometry"
	"github.com/samber/lo"
	"math"
	"regexp"
	"strings"
)

var robotRegex = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func part1(input string, args1 ...interface{}) interface{} {
	robots, width, height := parse(input, args1)
	return getSafetyFactor(robots, 100, width, height)
}

func part2(input string, args1 ...interface{}) interface{} {
	robots, width, height := parse(input, args1)
	minSafetyFactor := math.MaxInt
	minTime := math.MaxInt

	for t := 1; t <= 10_000; t++ {
		safetyFactor := getSafetyFactor(robots, 1, width, height)
		if safetyFactor < minSafetyFactor {
			minSafetyFactor = safetyFactor
			minTime = t
		}
		//fmt.Printf("\u001B[H\u001B[2J\n\nAfter %d seconds\n%s\n", t, render(robots, width, height))
		//time.Sleep(100 * time.Millisecond)
	}

	//r2, _, _ := parse(input, args1)
	//getSafetyFactor(r2, minTime, width, height)
	//fmt.Println(render(r2, width, height))
	return minTime
}

func getSafetyFactor(robots []*Robot, time, width, height int) int {
	quadrants := make([]int, 4)

	for _, r := range robots {
		r.X = ints.WrapMod(r.X+r.velocity.X*time, width)
		r.Y = ints.WrapMod(r.Y+r.velocity.Y*time, height)

		if r.Y < height/2 {
			if r.X < width/2 {
				quadrants[0]++
			} else if r.X > width/2 {
				quadrants[1]++
			}
		} else if r.Y > height/2 {
			if r.X < width/2 {
				quadrants[2]++
			} else if r.X > width/2 {
				quadrants[3]++
			}
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func parse(input string, args1 []interface{}) ([]*Robot, int, int) {
	width, height := args1[0].(int), args1[1].(int)
	return lo.Map(strings.Split(input, "\n"), func(line string, _ int) *Robot {
		values := convert.ToIntsV2(robotRegex.FindStringSubmatch(line)[1:])
		return &Robot{
			Point:    geometry.NewPoint(values[0], values[1]),
			velocity: geometry.NewPoint(values[2], values[3]),
		}
	}), width, height
}

func render(robots []*Robot, width, height int) string {
	grid := make([][]rune, height)
	for y := range height {
		grid[y] = make([]rune, width)
		for x := range width {
			grid[y][x] = '.'
		}
	}

	for _, r := range robots {
		x, y := r.X, r.Y
		cell := grid[r.Y][r.X]
		if cell == '.' {
			grid[y][x] = '1'
		} else {
			grid[y][x] += 1
		}
	}

	builder := strings.Builder{}
	delimiter := ""
	for _, row := range grid {
		builder.WriteString(delimiter)
		builder.WriteString(string(row))
		delimiter = "\n"
	}
	return builder.String()
}

type Robot struct {
	*geometry.Point
	velocity *geometry.Point
}

func main() {
	args := []interface{}{101, 103}
	New().Run(args, args)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 14, Part1: part1, Part2: part2}
}
