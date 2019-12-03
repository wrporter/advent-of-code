package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Vector struct {
	Direction string
	Magnitude int
}

var DeltaX = map[string]int{"R": 1, "L": -1, "D": 0, "U": 0}
var DeltaY = map[string]int{"R": 0, "L": 0, "D": -1, "U": 1}

func main() {
	wireLines, _ := file.ReadFile("./day3/input.txt")
	fmt.Println(Run(wireLines))
}

func Run(wireLines []string) int {
	var wires [][]string
	for _, wireLine := range wireLines {
		wires = append(wires, strings.Split(wireLine, ","))
	}
	return getMinimumDistanceIntersection(wires)
}

func NewVector(value string) Vector {
	direction := value[0:1]
	magnitudeInt64, _ := strconv.ParseInt(value[1:], 10, 64)
	magnitude := int(magnitudeInt64)
	return Vector{direction, magnitude}
}

func getMinimumDistanceIntersection(wires [][]string) int {
	intersections := getIntersections(wires)
	minDistance := int(^uint(0) >> 1)
	for _, intersection := range intersections {
		minDistance = Min(minDistance, getManhattanDistance(intersection))
	}
	return minDistance
}

func getManhattanDistance(point Point) int {
	return Abs(point.X) + Abs(point.Y)
}

func getIntersections(wires [][]string) (intersections []Point) {
	grid := make(map[Point]int)

	for wireId, wire := range wires {
		x := 0
		y := 0
		for _, vectorString := range wire {
			vector := NewVector(vectorString)
			for i := 0; i < vector.Magnitude; i++ {
				x += DeltaX[vector.Direction]
				y += DeltaY[vector.Direction]
				point := Point{x, y}

				if _, ok := grid[point]; ok && grid[point] != wireId {
					intersections = append(intersections, point)
				} else {
					grid[point] = wireId
				}
			}
		}
	}

	return intersections
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
