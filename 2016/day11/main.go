package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
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

	queue := []State

	return 0
}

func part2(input []string) interface{} {
	return 0
}

type (
	State struct {
		NumMoves      int
		Floors        []map[Item]bool
		ElevatorFloor int
	}
	Item struct {
		Element string
		Type    string
	}
)

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
