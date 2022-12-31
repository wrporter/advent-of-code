package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 3
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	location := convert.StringToInt(input[0])

	x := 0
	y := 0
	bottomRight := 1
	size := 3
	var topRight, topLeft, bottomLeft int

	for memory := 1; memory < location; memory++ {
		if memory == bottomRight {
			x++

			topRight = bottomRight + size - 1
			topLeft = topRight + size - 1
			bottomLeft = topLeft + size - 1
			bottomRight = bottomLeft + size - 1
			size += 2
		} else if memory < topRight {
			y++
		} else if memory < topLeft {
			x--
		} else if memory < bottomLeft {
			y--
		} else if memory < bottomRight {
			x++
		}
	}

	return ints.Abs(x) + ints.Abs(y)
}

func part2(input []string) interface{} {
	destination := convert.StringToInt(input[0])
	value := 1
	grid := map[geometry.Point]int{
		geometry.NewPoint(0, 0): 1,
	}

	x := 0
	y := 0
	bottomRight := 1
	size := 3
	var topRight, topLeft, bottomLeft int

	for memory := 1; value <= destination; memory++ {
		if memory == bottomRight {
			x++

			topRight = bottomRight + size - 1
			topLeft = topRight + size - 1
			bottomLeft = topLeft + size - 1
			bottomRight = bottomLeft + size - 1
			size += 2
		} else if memory < topRight {
			y++
		} else if memory < topLeft {
			x--
		} else if memory < bottomLeft {
			y--
		} else if memory < bottomRight {
			x++
		}

		value = 0
		for _, direction := range geometry.AllDirectionsModifiers {
			location := geometry.NewPoint(x+direction.X, y+direction.Y)
			if locationValue, ok := grid[location]; ok {
				value += locationValue
			}
		}

		grid[geometry.NewPoint(x, y)] = value
	}

	return value
}
