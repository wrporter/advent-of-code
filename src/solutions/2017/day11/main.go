package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 11
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	return getDistance(input[0])
}

func part2(input []string) interface{} {
	return getFurthestDistance(input[0])
}

func getDistance(stepsStr string) int {
	steps := strings.Split(stepsStr, ",")
	start := HexCube{}
	position := start

	for _, step := range steps {
		direction := toDirection(step)
		position = position.Move(direction)
	}

	distance := start.Distance(position)
	return distance
}

func getFurthestDistance(stepsStr string) int {
	steps := strings.Split(stepsStr, ",")
	start := HexCube{}
	position := start
	maxDistance := 0

	for _, step := range steps {
		direction := toDirection(step)
		position = position.Move(direction)
		distance := start.Distance(position)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func toDirection(step string) Direction {
	switch step {
	case "n":
		return North
	case "ne":
		return NorthEast
	case "se":
		return SouthEast
	case "s":
		return South
	case "sw":
		return SouthWest
	case "nw":
		return NorthWest
	default:
		return North
	}
}

type (
	// HexCube comes from the cube coordinate system described in https://www.redblobgames.com/grids/hexagons/.
	HexCube struct {
		Q int
		S int
		R int
	}

	Direction int
)

const (
	North Direction = iota
	NorthEast
	SouthEast
	South
	SouthWest
	NorthWest
)

var Directions = []Direction{
	North,
	NorthEast,
	SouthEast,
	South,
	SouthWest,
	NorthWest,
}

func (h HexCube) Move(direction Direction) HexCube {
	q, s, r := h.Q, h.S, h.R

	switch direction {
	case North:
		s += 1
		r -= 1
	case NorthEast:
		q += 1
		r -= 1
	case SouthEast:
		q += 1
		s -= 1
	case South:
		s -= 1
		r += 1
	case SouthWest:
		q -= 1
		r += 1
	case NorthWest:
		q -= 1
		s += 1
	}

	return HexCube{Q: q, S: s, R: r}
}

func (h HexCube) Subtract(b HexCube) HexCube {
	return HexCube{
		Q: h.Q - b.Q,
		S: h.S - b.S,
		R: h.R - b.R,
	}
}

func (h HexCube) Distance(b HexCube) int {
	vector := h.Subtract(b)
	return ints.Max(
		ints.Abs(vector.Q),
		ints.Abs(vector.S),
		ints.Abs(vector.R),
	)
}
