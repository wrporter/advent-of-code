package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Intersection struct {
	Point Point
	Steps int
}

type Vector struct {
	Direction string
	Magnitude int
}

var DeltaX = map[string]int{"R": 1, "L": -1, "D": 0, "U": 0}
var DeltaY = map[string]int{"R": 0, "L": 0, "D": -1, "U": 1}

func main() {
	wireLines, _ := file.ReadFile("./2019/day3/input.txt")
	minIntersectionDistance, minStepDistance := Run(wireLines)
	fmt.Printf("Intersection Distance: %d\nStep Distance: %d", minIntersectionDistance, minStepDistance)
}

func Run(wireLines []string) (minIntersectionDistance int, minStepDistance int) {
	var wires [][]string
	for _, wireLine := range wireLines {
		wires = append(wires, strings.Split(wireLine, ","))
	}
	intersections := getIntersections(wires)
	return getMinimumDistance(intersections)
}

func getMinimumDistance(intersections []Intersection) (int, int) {
	minIntersectionDistance := int(^uint(0) >> 1)
	minStepDistance := int(^uint(0) >> 1)
	for _, intersection := range intersections {
		minIntersectionDistance = Min(minIntersectionDistance, getManhattanDistance(intersection.Point))
		minStepDistance = Min(minStepDistance, intersection.Steps)
	}
	return minIntersectionDistance, minStepDistance
}

func getIntersections(wires [][]string) (intersections []Intersection) {
	grid := make(map[Point]map[int]int)

	for wireID, wire := range wires {
		x := 0
		y := 0
		steps := 0
		for _, vectorString := range wire {
			vector := NewVector(vectorString)
			for i := 0; i < vector.Magnitude; i++ {
				x += DeltaX[vector.Direction]
				y += DeltaY[vector.Direction]
				steps++
				point := Point{x, y}

				if pointMap, ok := grid[point]; ok {
					for otherWireID, otherWireSteps := range pointMap {
						if otherWireID != wireID {
							intersections = append(intersections, Intersection{point, otherWireSteps + steps})
						}
					}
				}

				grid[point] = make(map[int]int)
				grid[point][wireID] = steps
			}
		}
	}

	return intersections
}

func NewVector(value string) Vector {
	direction := value[0:1]
	magnitudeInt64, _ := strconv.ParseInt(value[1:], 10, 64)
	magnitude := int(magnitudeInt64)
	return Vector{direction, magnitude}
}

func getManhattanDistance(point Point) int {
	return Abs(point.X) + Abs(point.Y)
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
