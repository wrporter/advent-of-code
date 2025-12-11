package main

import (
	"aoc/src/lib/go/aoc"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	devices := parse(input)
	return countPaths(devices, "you", "out")
}

func part2(input string, _ ...interface{}) interface{} {
	devices := parse(input)

	memo := make(map[cacheKey]int)
	var count func(current, target string) int

	count = func(current, target string) int {
		key := cacheKey{current, target}
		if _, ok := memo[key]; ok {
			return memo[key]
		}

		if current == target {
			memo[key] = 1
			return 1
		}

		sum := 0
		for _, next := range devices[current] {
			sum += count(next, target)
		}

		memo[key] = sum
		return sum
	}

	return count("svr", "dac")*count("dac", "fft")*count("fft", "out") +
		count("svr", "fft")*count("fft", "dac")*count("dac", "out")
}

type cacheKey struct {
	current string
	target  string
}

func countPaths(devices map[string][]string, start string, end string) interface{} {
	paths := 0
	queue := []string{start}
	var current string

	for len(queue) > 0 {
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]

		if current == end {
			paths++
			continue
		}

		for _, output := range devices[current] {
			queue = append(queue, output)
		}
	}

	return paths
}

func parse(input string) map[string][]string {
	devices := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		id := parts[0]
		outputs := strings.Fields(parts[1])
		devices[id] = outputs
	}
	return devices
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 11, Part1: part1, Part2: part2}
}
