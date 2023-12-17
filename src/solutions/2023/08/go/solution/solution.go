package solution

import (
	"aoc/src/lib/go/v2/mymath"
	"regexp"
	"strings"
	"sync"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	instructions, network := parseInput(input)

	if _, ok := network["ZZZ"]; !ok {
		return -1
	}

	return calculateSteps(instructions, network, "AAA", func(location string) bool {
		return location == "ZZZ"
	})
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	instructions, network := parseInput(input)

	var starts []string
	for location := range network {
		if location[len(location)-1] == 'A' {
			starts = append(starts, location)
		}
	}

	done := func(location string) bool {
		return location[2] == 'Z'
	}

	var first []int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, start := range starts {
		wg.Add(1)

		go func(start string) {
			steps := calculateSteps(instructions, network, start, done)

			mu.Lock()
			first = append(first, steps)
			mu.Unlock()

			wg.Done()
		}(start)
	}

	wg.Wait()

	return mymath.LCM(first...)
}

func calculateSteps(instructions []string, network map[string]node, start string, done func(location string) bool) int {
	location := start
	steps := 0

	for !done(location) {
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
