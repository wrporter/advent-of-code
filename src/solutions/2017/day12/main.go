package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 12
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^(\d+) <-> (.*)$`)

func part1(input []string) interface{} {
	pipes := parseInput(input)

	numProgramsInGroup0 := 0
	targetGroup := 0

	for programId := range pipes {
		if isInGroup(pipes, make(map[int]bool), programId, targetGroup) {
			numProgramsInGroup0++
		}
	}

	return numProgramsInGroup0
}

func part2(input []string) interface{} {
	pipes := parseInput(input)

	grouped := make(map[int]bool)
	numGroups := 0
	var currentProgramId int

	for groupId := range pipes {
		if grouped[groupId] {
			continue
		}

		numGroups++
		queue := []int{groupId}

		for len(queue) > 0 {
			currentProgramId, queue = ints.Pop(queue)

			for nextProgramId := range pipes[currentProgramId] {
				if !grouped[nextProgramId] {
					grouped[nextProgramId] = true
					queue = append(queue, nextProgramId)
				}
			}
		}
	}

	return numGroups
}

func parseInput(input []string) map[int]map[int]bool {
	pipes := make(map[int]map[int]bool)

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		fromId := convert.StringToInt(match[1])
		toIds, _ := convert.ToInts(strings.Split(match[2], ", "))

		toIdsMap := make(map[int]bool)
		for _, id := range toIds {
			toIdsMap[id] = true
		}
		pipes[fromId] = toIdsMap
	}
	return pipes
}

func isInGroup(pipes map[int]map[int]bool, visited map[int]bool, programId int, group int) bool {
	visited[programId] = true

	if programId == group || pipes[programId][group] {
		return true
	}

	for nextProgramId := range pipes[programId] {
		if !visited[nextProgramId] && isInGroup(pipes, visited, nextProgramId, group) {
			return true
		}
	}

	return false
}
