package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 24
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	blackTiles := initialState(input)
	return len(blackTiles)
}

func part2(input []string) interface{} {
	blackTiles := initialState(input)
	blackTiles = run(blackTiles, 100)
	return len(blackTiles)
}

func run(blackTiles map[Point]bool, numDays int) map[Point]bool {
	for day := 1; day <= numDays; day++ {
		blackTiles = cycle(blackTiles)
	}
	return blackTiles
}

func cycle(blackTiles map[Point]bool) map[Point]bool {
	next := make(map[Point]bool)

	for tile := range blackTiles {
		flip(blackTiles, next, tile)

		for _, direction := range Directions {
			neighbor := tile.Move(direction)
			flip(blackTiles, next, neighbor)
		}
	}

	return next
}

func flip(blackTiles map[Point]bool, next map[Point]bool, tile Point) {
	numBlackNeighbors := countBlackNeighbors(blackTiles, tile)

	if blackTiles[tile] {
		if numBlackNeighbors == 0 || numBlackNeighbors > 2 {
			// do nothing
		} else {
			next[tile] = true
		}
	} else if numBlackNeighbors == 2 {
		next[tile] = true
	}
}

func countBlackNeighbors(blackTiles map[Point]bool, tile Point) int {
	count := 0
	for _, direction := range Directions {
		neighbor := tile.Move(direction)
		if blackTiles[neighbor] {
			count++
		}
	}
	return count
}

func initialState(input []string) map[Point]bool {
	blackTiles := make(map[Point]bool)

	for _, line := range input {
		current := Point{0, 0, 0}

		for i := 0; i < len(line); i++ {
			char := line[i]
			next := peek(line, i)

			if char == 'e' {
				current = current.Move(East)
			} else if char == 'w' {
				current = current.Move(West)
			} else if char == 'n' && next == 'e' {
				current = current.Move(NorthEast)
				i++
			} else if char == 'n' && next == 'w' {
				current = current.Move(NorthWest)
				i++
			} else if char == 's' && next == 'e' {
				current = current.Move(SouthEast)
				i++
			} else if char == 's' && next == 'w' {
				current = current.Move(SouthWest)
				i++
			}
		}

		if blackTiles[current] {
			delete(blackTiles, current)
		} else {
			blackTiles[current] = true
		}
	}
	return blackTiles
}

func peek(str string, index int) uint8 {
	if (index + 1) >= len(str) {
		return 0
	}
	return str[index+1]
}

type (
	Point struct {
		A int
		R int
		C int
	}

	Direction int
)

const (
	East Direction = iota
	SouthEast
	SouthWest
	West
	NorthWest
	NorthEast
)

var Directions = []Direction{
	East,
	SouthEast,
	SouthWest,
	West,
	NorthWest,
	NorthEast,
}

func (p Point) Move(direction Direction) Point {
	// https://en.wikipedia.org/wiki/Hexagonal_Efficient_Coordinate_System#/media/File:HECS_Nearest_Neighbors.jpg.png
	a := p.A
	r := p.R
	c := p.C

	switch direction {
	case East:
		c += 1
	case SouthEast:
		r += a
		c += a
		a = 1 - a
	case SouthWest:
		r += a
		a = 1 - a
		c -= a
	case West:
		c -= 1
	case NorthWest:
		a = 1 - a
		r -= a
		c -= a
	case NorthEast:
		c += a
		a = 1 - a
		r -= a
	}

	return Point{a, r, c}
}
