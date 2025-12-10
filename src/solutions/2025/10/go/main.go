package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	machines := parse(input)

	total := 0
	for _, machine := range machines {
		total += findMinNumPresses(machine)
	}

	return total
}

func part2(input string, _ ...interface{}) interface{} {
	return "TBD"
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 10, Part1: part1, Part2: part2}
}

func findMinNumPresses(machine Machine) int {
	queue := []Node{{lights: 0, numPresses: 0}}
	var current Node
	seen := map[int]bool{0: true}

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current.lights == machine.lights {
			return current.numPresses
		}

		for _, button := range machine.buttons {
			next := toggleBits(current.lights, button)

			if !seen[next] {
				seen[next] = true
				queue = append(queue, Node{
					lights:     next,
					numPresses: current.numPresses + 1,
				})
			}
		}
	}

	return -1
}

var regex = regexp.MustCompile(`^\[([.#]+)] ([(\d,) ?]+) \{([\d,]+)}$`)

type Machine struct {
	// Bitmask to represent the expected indicator lights
	lights int
	// Bitmasks of lights that turn on when a button is pressed
	buttons  []int
	joltages []int
}

type Node struct {
	lights     int
	numPresses int
}

func parse(input string) []Machine {
	lines := strings.Split(input, "\n")
	machines := make([]Machine, len(lines))
	for i, line := range lines {
		match := regex.FindStringSubmatch(line)

		lightsStr := match[1]
		lights := 0
		for j, light := range lightsStr {
			if light == '#' {
				lights = setBit(lights, j)
			}
		}

		buttonStrs := strings.Split(match[2], " ")
		buttons := make([]int, len(buttonStrs))
		for j, buttonStr := range buttonStrs {
			buttonStr = strings.TrimPrefix(buttonStr, "(")
			buttonStr = strings.TrimSuffix(buttonStr, ")")
			for _, v := range strings.Split(buttonStr, ",") {
				bit, _ := strconv.Atoi(v)
				buttons[j] = setBit(buttons[j], bit)
			}
		}

		joltages, _ := convert.ToInts(strings.Split(match[3], ","))

		machines[i] = Machine{
			lights:   lights,
			buttons:  buttons,
			joltages: joltages,
		}
	}
	return machines
}

func setBit(n int, pos int) int {
	n |= 1 << pos
	return n
}

func isBitSet(n int, pos uint) bool {
	mask := 1 << pos
	return (n & mask) != 0
}

func clearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func toggleBits(n int, mask int) int {
	return n ^ mask
}
