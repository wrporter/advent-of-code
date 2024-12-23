package main

import (
	"aoc/src/lib/go/aoc"
	"github.com/samber/lo"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	computers, connections := parse(input)
	count := 0

	for i := 0; i < len(computers); i++ {
		for j := i + 1; j < len(computers); j++ {
			for k := j + 1; k < len(computers); k++ {
				a, b, c := computers[i], computers[j], computers[k]

				if (a[0] == 't' || b[0] == 't' || c[0] == 't') &&
					connections[[2]string{a, b}] &&
					connections[[2]string{b, c}] &&
					connections[[2]string{c, a}] {
					count++
				}
			}
		}
	}

	return count
}

func part2(input string, _ ...interface{}) interface{} {
	computers, connections := parse(input)

	networks := make([]map[string]bool, 0)
	for _, computer := range computers {
		networks = append(networks, map[string]bool{computer: true})
	}

	var largest map[string]bool
	for _, network := range networks {
		for _, a := range computers {
			isConnected := true

			for b := range network {
				if !connections[[2]string{a, b}] {
					isConnected = false
					break
				}
			}

			if isConnected {
				network[a] = true
			}
		}

		if len(network) > len(largest) {
			largest = network
		}
	}

	computerNames := lo.Keys(largest)
	sort.Strings(computerNames)
	password := strings.Join(computerNames, ",")

	return password
}

func parse(input string) ([]string, map[[2]string]bool) {
	computers := make(map[string]bool)
	connections := make(map[[2]string]bool)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		computers[a] = true
		computers[b] = true
		connections[[2]string{a, b}] = true
		connections[[2]string{b, a}] = true
	}

	return lo.Keys(computers), connections
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 23, Part1: part1, Part2: part2}
}
