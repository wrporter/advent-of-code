package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"regexp"
	"strings"
)

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

var Directions = []Direction{North, East, South, West}
var DirectionModifiers = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p *Point) Add(direction Direction, amount int) {
	p.X += DirectionModifiers[direction].X * amount
	p.Y += DirectionModifiers[direction].Y * amount
}

var regex = regexp.MustCompile(`^(R|L)(\d+)$`)

func distance(point Point) int {
	return ints.Abs(point.X) + ints.Abs(point.Y)
}

func part1(instructions []string) int {
	position := NewPoint(0, 0)
	direction := North

	for _, instruction := range instructions {
		match := regex.FindStringSubmatch(instruction)
		rotation, steps := match[1], convert.StringToInt(match[2])

		modifier := 0
		if rotation == "R" {
			modifier = 1
		} else if rotation == "L" {
			modifier = -1
		}
		direction = Directions[ints.WrapMod(int(Directions[direction])+modifier, len(Directions))]

		position.Add(direction, steps)
	}

	return ints.Abs(position.X) + ints.Abs(position.Y)
}

func part2(instructions []string) int {
	position := NewPoint(0, 0)
	direction := North
	visited := make(map[Point]bool)
	visited[NewPoint(position.X, position.Y)] = true

	for _, instruction := range instructions {
		match := regex.FindStringSubmatch(instruction)
		rotation, steps := match[1], convert.StringToInt(match[2])

		modifier := 0
		if rotation == "R" {
			modifier = 1
		} else if rotation == "L" {
			modifier = -1
		}
		direction = Directions[ints.WrapMod(int(Directions[direction])+modifier, len(Directions))]

		for numSteps := 0; numSteps < steps; numSteps++ {
			position.Add(direction, 1)
			if visited[position] {
				return distance(position)
			}
			visited[NewPoint(position.X, position.Y)] = true
		}
	}

	return distance(position)
}

func main() {
	instructions, _ := file.ReadFile("./2016/day1/input.txt")
	result1 := part1(strings.Split(instructions[0], ", "))
	result2 := part2(strings.Split(instructions[0], ", "))
	fmt.Println(result1)
	fmt.Println(result2)
}
