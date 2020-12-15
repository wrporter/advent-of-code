package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"math"
	"regexp"
	"strings"
)

func main() {
	year, day := 2016, 11
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`a ([a-z]+)(?:-compatible)? (generator|microchip)`)

func part1(input []string) interface{} {
	floors := parseFloors(input)
	queue := []State{{}}
	var state State

	maxMoves := 0

	for len(queue) > 0 {
		state, queue = queue[0], queue[1:]

		maxMoves = ints.Max(maxMoves, state.NumMoves)

		bestPositive := 0
		bestNegative := -math.MaxFloat64

		var clearedFloor int
		for i, floor := range state.Floors {
			if len(floor) == 0 {
				clearedFloor = i
			}
		}

		fmt.Println(bestPositive, bestNegative)
	}

	return 0
}

func part2(input []string) interface{} {
	return 0
}

type (
	State struct {
		NumMoves int
		Floors   []map[Item]bool
		Elevator int
	}
	Item struct {
		Element string
		Type    string
	}
)

func (s []State) Moves(elevator int) {
	state := s[elevator]

	var chips []Item
	var generators []Item
	for item := range state. {
		if item.Type == "generator" {

		}
	}


}

func parseFloors(input []string) []map[Item]bool {
	floors := make([]map[Item]bool, len(input))
	for floor, line := range input {
		floors[floor] = make(map[Item]bool)
		if !strings.Contains(line, "nothing relevant") {
			matches := regex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				item := Item{
					Element: match[1],
					Type:    match[2],
				}
				floors[floor][item] = true
			}
		}
	}
	return floors
}
