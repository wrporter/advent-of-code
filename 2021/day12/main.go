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
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	connections := parseInput(input)
	paths := findPaths(connections, make(map[string]bool), "start", "", "")
	return len(paths)
}

func part2(input []string) interface{} {
	connections := parseInput(input)
	paths := make(map[string]bool)

	for cave := range connections {
		if isSmall(cave) && cave != "start" && cave != "end" {
			findPaths2(connections, make(map[string]int), "start", cave, "", "", paths)
		}
	}

	return len(paths)
}

func findPaths(connections map[string][]string, visited map[string]bool, current string, path string, delimiter string) []string {
	path += delimiter + current
	delimiter = "-"
	var paths []string

	if current == "end" {
		paths = append(paths, path)
		return paths
	}

	visited[current] = true
	for _, next := range connections[current] {
		if !isSmall(next) || !visited[next] {
			nextPaths := findPaths(connections, visited, next, path, delimiter)
			paths = append(paths, nextPaths...)
		}
	}
	visited[current] = false

	return paths
}

func findPaths2(connections map[string][]string, visited map[string]int, current string, twice string, path string, delim string, paths map[string]bool) {
	path += delim + current
	delim = "-"

	if current == "end" {
		paths[path] = true
		return
	}

	visited[current]++
	for _, next := range connections[current] {
		if !isSmall(next) ||
			(next == twice && visited[next] < 2) ||
			(next != twice && visited[next] < 1) {
			findPaths2(connections, visited, next, twice, path, delim, paths)
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
