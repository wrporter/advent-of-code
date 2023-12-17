package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"fmt"
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
		location = location.Move(toDir(arrow))
		houses[location]++
	}

	return len(houses)
}

func deliverWithRobo(directions string) int {
	santa := geometry.NewPoint(0, 0)
	roboSanta := geometry.NewPoint(0, 0)
	houses := map[geometry.Point]int{santa: 2}

	santaMove := true
	for _, arrow := range directions {
		if santaMove {
			santa = santa.Move(toDir(arrow))
			houses[santa]++
		} else {
			roboSanta = roboSanta.Move(toDir(arrow))
			houses[roboSanta]++
		}
		santaMove = !santaMove
	}

	return len(houses)
}

func main() {
	lines, _ := file.ReadFile("./2015/day3/input.txt")
	input := lines[0]
	fmt.Println(deliver(input))
	fmt.Println(deliverWithRobo(input))
}
