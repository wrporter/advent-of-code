package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
)

func toDir(arrow rune) geometry.Direction {
	switch arrow {
	case '^':
		return geometry.Up
	case 'v':
		return geometry.Down
	case '<':
		return geometry.Left
	case '>':
		return geometry.Right
	default:
		return geometry.Up
	}
}

func deliver(directions string) int {
	location := geometry.NewPoint(0, 0)
	houses := map[geometry.Point]int{location: 1}

	for _, arrow := range directions {
		location = location.Add(toDir(arrow))
		houses[location]++
	}

	return len(houses)
}

func main() {
	lines, _ := file.ReadFile("./2015/day3/input.txt")
	input := lines[0]
	fmt.Println(deliver(input))
}
