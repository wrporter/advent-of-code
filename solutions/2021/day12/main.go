package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 12
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	connections := parseInput(input)
	paths := findPaths(connections, make(map[string]bool), "start", "start")
	return len(paths)
}

func part2(input []string) interface{} {
	connections := parseInput(input)
	paths := make(map[string]bool)
	findPaths2(connections, make(map[string]int), "start", "start", false, paths)
	return len(paths)
}

func findPaths(connections map[string][]string, visited map[string]bool, current string, path string) []string {
	var paths []string

	visited[current] = true
	for _, next := range connections[current] {
		if next == "end" {
			paths = append(paths, path+"-end")
		} else if !isSmall(next) || !visited[next] {
			nextPaths := findPaths(connections, visited, next, path+"-"+next)
			paths = append(paths, nextPaths...)
		}
	}
	visited[current] = false

	return paths
}

func findPaths2(connections map[string][]string, visited map[string]int, current string, path string, twice bool, paths map[string]bool) {
	visited[current]++
	for _, next := range connections[current] {
		if next == "start" {
			continue
		} else if next == "end" {
			paths[path+"-end"] = true
		} else if !isSmall(next) || visited[next] < 1 || (visited[next] == 1 && !twice) {
			nextTwice := twice || (isSmall(next) && visited[next] == 1 && !twice)
			findPaths2(connections, visited, next, path+"-"+next, nextTwice, paths)
		}
	}
	visited[current]--
}

func parseInput(input []string) (connections map[string][]string) {
	connections = make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, "-")
		connections[parts[0]] = append(connections[parts[0]], parts[1])
		connections[parts[1]] = append(connections[parts[1]], parts[0])
	}

	return connections
}

func isSmall(cave string) bool {
	return mystrings.IsLower(cave)
}
