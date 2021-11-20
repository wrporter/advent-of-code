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

	year, day := 2017, 25
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	stateName, checksum, states := parseInput(input)
	tape := make(map[int]int)
	cursor := 0

	for i := 0; i < checksum; i++ {
		state := states[stateName]
		condition := state.Conditions[tape[cursor]]

		tape[cursor] = condition.WriteValue
		cursor += condition.Move
		stateName = condition.Goto
	}

	return calculateDiagnosticChecksum(tape)
}

func parseInput(input []string) (string, int, map[string]State) {
	allInput := strings.Join(input, "\n")
	stateName := beginRegex.FindStringSubmatch(allInput)[1]
	checksum := conversion.StringToInt(checksumRegex.FindStringSubmatch(allInput)[1])
	stateMatches := stateRegex.FindAllStringSubmatch(allInput, -1)

	states := make(map[string]State)
	for _, stateMatch := range stateMatches {
		state := State{
			Name: stateMatch[1],
		}
		state.Conditions = append(state.Conditions, Condition{
			WriteValue: conversion.StringToInt(stateMatch[3]),
			Move:       getMoveValue(stateMatch[4]),
			Goto:       stateMatch[5],
		})
		state.Conditions = append(state.Conditions, Condition{
			WriteValue: conversion.StringToInt(stateMatch[7]),
			Move:       getMoveValue(stateMatch[8]),
			Goto:       stateMatch[9],
		})
		states[stateMatch[1]] = state
	}
	return stateName, checksum, states
}

func part2(input []string) interface{} {
	return "Merry Christmas!!!"
}

func calculateDiagnosticChecksum(tape map[int]int) int {
	checksum := 0
	for _, value := range tape {
		if value == 1 {
			checksum++
		}
	}
	return checksum
}

func getMoveValue(direction string) int {
	switch direction {
	case "left":
		return -1
	case "right":
		return 1
	}
	return 0
}

var (
	beginRegex    = regexp.MustCompile(`Begin in state ([A-Z])\.`)
	checksumRegex = regexp.MustCompile(`Perform a diagnostic checksum after (\d+) steps\.`)
	stateRegex    = regexp.MustCompile(`In state ([A-Z]):
  If the current value is (\d):
    - Write the value (\d)\.
    - Move one slot to the (left|right)\.
    - Continue with state ([A-Z])\.
  If the current value is (\d):
    - Write the value (\d)\.
    - Move one slot to the (left|right)\.
    - Continue with state ([A-Z])\.`)
)

type (
	State struct {
		Name       string
		Conditions []Condition
	}
	Condition struct {
		WriteValue int
		Move       int
		Goto       string
	}
)
