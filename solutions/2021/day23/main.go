package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

//#############
//#...........#
//###B#C#B#D###
//  #A#D#C#A#
//  #########

func part1(input []string) interface{} {
	burrow := make(map[point]bool)
	amphipods := make(map[point]amphipod)
	for y, row := range input {
		for x, char := range row {
			p := point{x: x, y: y}
			if char == '.' {
				burrow[p] = true
			}

			if char >= 'A' && char <= 'D' {
				burrow[p] = true
				amphipods[p] = newAmphipod(p, char)
			}
		}
	}

	energy, _ := findLeastEnergy(burrow, amphipods, make(map[int]bool), 0)
	return energy
}

func findLeastEnergy(burrow map[point]bool, amphipods map[point]amphipod, seen map[int]bool, energy int) (int, bool) {
	if organized(amphipods) {
		return energy, true
	}

	min := math.MaxInt

	for prev, a := range amphipods {
		for _, d := range directions {
			p := point{x: a.x + d.x, y: a.y + d.y}
			_, unavailable := amphipods[p]
			if burrow[p] && !unavailable { // TODO check burrow not seen
				delete(amphipods, a.point)
				a.point = p
				amphipods[a.point] = a

				e, done := findLeastEnergy(burrow, amphipods, seen, energy+a.energy)
				if done {
					min = ints.Min(min, e)
				}

				delete(amphipods, a.point)
				a.point = prev
				amphipods[a.point] = a
			}
		}
	}

	return min, true
}

func organized(amphipods map[point]amphipod) bool {
	for _, a := range amphipods {
		if !a.isInOwnRoom() {
			return false
		}
	}
	return true
	//return amphipods[point{x: 3, y: 2}].value == 'A' && amphipods[point{x: 3, y: 3}].value == 'A' &&
	//	amphipods[point{x: 5, y: 2}].value == 'B' && amphipods[point{x: 5, y: 3}].value == 'B' &&
	//	amphipods[point{x: 7, y: 2}].value == 'C' && amphipods[point{x: 7, y: 3}].value == 'C' &&
	//	amphipods[point{x: 9, y: 2}].value == 'D' && amphipods[point{x: 9, y: 3}].value == 'D'
}

func part2(input []string) interface{} {
	return 0
}

var directions = []point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

var eligibleHallwaySpots = []int{1, 2, 4, 6, 8, 10, 11}

type move struct {
	point
	energy int
}

type point struct {
	x, y int
}

type amphipod struct {
	point
	value  rune
	energy int
	room   int
}

func newAmphipod(p point, value rune) amphipod {
	return amphipod{
		point:  p,
		value:  value,
		energy: ints.Pow(10, int(value-'A')),
		room:   ((int(value-'A') + 1) * 2) + 1,
	}
}

func (a amphipod) getPossibleMoves(burrow map[point]bool, amphipods map[point]amphipod) []move {
	if !a.canMove(burrow, amphipods) {
		return nil
	}

	var moves []move
	//for _, hall := range eligibleHallwaySpots {
	//
	//	for x := a.x; x
	//}
	return moves
}

func (a amphipod) canMove(burrow map[point]bool, amphipods map[point]amphipod) bool {
	for _, diff := range directions {
		spot := point{x: a.x + diff.x, y: a.y + diff.y}
		if _, taken := amphipods[spot]; !taken && burrow[spot] {
			return true
		}
	}
	return false
}

func (a amphipod) isInHallway() bool {
	return a.y == 1
}

func (a amphipod) isInOwnRoom() bool {
	return a.y >= 2 && a.x == a.room
}

func (a amphipod) isInDifferentRoom() bool {
	return a.y >= 2 && a.x != a.room
}
