package main

import (
	"aoc/src/lib/go/convert"
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

	year, day := 2017, 24
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	connections := parseInput(input)

	port := 0
	maxStrength := 0
	for _, component := range connections[port] {
		used := map[Component]bool{component: true}
		strength := findMaxStrength(connections, used, component, port, component.Strength())
		maxStrength = ints.Max(maxStrength, strength)
	}

	return maxStrength
}

func part2(input []string) interface{} {
	connections := parseInput(input)

	port := 0
	maxStrength := 0
	maxLength := 0
	for _, component := range connections[port] {
		used := map[Component]bool{component: true}
		strength, length := findMaxStrengthOfLongestBridge(connections, used, component, port, component.Strength(), 1)
		if length > maxLength || (length == maxLength && strength > maxStrength) {
			maxStrength = ints.Max(maxStrength, strength)
			maxLength = ints.Max(maxLength, length)
		}
	}

	return maxStrength
}

func findMaxStrengthOfLongestBridge(connections map[int][]Component, used map[Component]bool, component Component, portA int, strength int, length int) (int, int) {
	portB := component.OtherPort(portA)
	maxStrength := strength
	maxLength := length

	for _, next := range connections[portB] {
		if !used[next] {
			used[next] = true
			nextStrength, nextLength := findMaxStrengthOfLongestBridge(connections, used, next, portB, strength+next.Strength(), length+1)
			if nextLength > maxLength || (nextLength == maxLength && nextStrength > maxStrength) {
				maxStrength = ints.Max(maxStrength, nextStrength)
				maxLength = ints.Max(maxLength, nextLength)
			}
			used[next] = false
		}
	}

	return maxStrength, maxLength
}

func findMaxStrength(connections map[int][]Component, used map[Component]bool, component Component, portA int, strength int) int {
	portB := component.OtherPort(portA)
	maxStrength := strength

	for _, next := range connections[portB] {
		if !used[next] {
			used[next] = true
			nextStrength := findMaxStrength(connections, used, next, portB, strength+next.Strength())
			maxStrength = ints.Max(maxStrength, nextStrength)
			used[next] = false
		}
	}

	return maxStrength
}

func parseInput(input []string) map[int][]Component {
	connections := make(map[int][]Component)
	for i, line := range input {
		parts := strings.Split(line, "/")
		portA := convert.StringToInt(parts[0])
		portB := convert.StringToInt(parts[1])
		component := Component{
			ID:    i,
			PortA: portA,
			PortB: portB,
		}
		connections[component.PortA] = append(connections[component.PortA], component)
		connections[component.PortB] = append(connections[component.PortB], component)
	}
	return connections
}

type (
	Component struct {
		ID    int
		PortA int
		PortB int
	}
)

func (c Component) OtherPort(port int) int {
	if port == c.PortA {
		return c.PortB
	}
	return c.PortA
}

func (c Component) Strength() int {
	return c.PortA + c.PortB
}
