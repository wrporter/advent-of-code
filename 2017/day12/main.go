package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 12
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

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
	groups := make(map[int]map[int]bool)

	for targetGroupId := range pipes {
		for programId := range pipes {

			if !isAlreadyInAGroup(groups, programId) && isInGroup(pipes, make(map[int]bool), programId, targetGroupId) {
				if groups[targetGroupId] == nil {
					groups[targetGroupId] = make(map[int]bool)
				}
				groups[targetGroupId][programId] = true
			}
		}
	}

	return len(groups)
}

func isAlreadyInAGroup(groups map[int]map[int]bool, programId int) bool {
	isAlreadyInAGroup := false
	for _, group := range groups {
		if group[programId] {
			isAlreadyInAGroup = true
		}
	}
	return isAlreadyInAGroup
}

func parseInput(input []string) map[int]map[int]bool {
	pipes := make(map[int]map[int]bool)

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		fromId := conversion.StringToInt(match[1])
		toIds, _ := conversion.ToInts(strings.Split(match[2], ", "))

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
