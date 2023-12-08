package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/v2/mymath"
	"regexp"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	instructions, network := parseInput(input)
	location := "AAA"
	steps := 0

	if _, ok := network["ZZZ"]; !ok {
		return -1
	}

	for location != "ZZZ" {
		instruction := instructions[steps%len(instructions)]
		steps++

		if instruction == "L" {
			location = network[location].left
		} else {
			location = network[location].right
		}
	}

	return steps
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	instructions, network := parseInput(input)
	steps := 0
	end := false

	locations := make(map[string]string)
	for location := range network {
		if location[len(location)-1] == 'A' {
			locations[location] = location
		}
	}

	first := make(map[string]int)

	for !end {
		next := make(map[string]string)
		end = true
		instruction := instructions[steps%len(instructions)]
		steps++

		for from, location := range locations {
			if instruction == "L" {
				next[from] = network[location].left
			} else {
				next[from] = network[location].right
			}

			if _, ok := first[from]; !ok && next[from][2] == 'Z' {
				first[from] = steps
			}

			if len(first) == len(locations) {
				var values []int
				for _, value := range first {
					values = append(values, value)
				}
				return mymath.LCM(values...)
			}

			end = end && next[from][2] == 'Z'
		}

		locations = next
	}

	return steps
}

var networkRegex = regexp.MustCompile(`[0-9A-Z]{3}`)

func parseInput(input string) ([]string, map[string]node) {
	chunks := strings.Split(input, "\n\n")
	instructions := strings.Split(chunks[0], "")
	network := make(map[string]node)

	for _, line := range strings.Split(chunks[1], "\n") {
		parts := networkRegex.FindAllString(line, 3)
		network[parts[0]] = node{parts[1], parts[2]}
	}

	return instructions, network
}

type node struct {
	left  string
	right string
}
