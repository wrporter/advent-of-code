package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
)

type instructionCommand string

const (
	on     instructionCommand = "turn on"
	toggle instructionCommand = "toggle"
	off    instructionCommand = "turn off"
)

type instruction struct {
	command  instructionCommand
	startCol int
	startRow int
	endCol   int
	endRow   int
}

var regex = regexp.MustCompile(`^(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)$`)

func setupLights(instructions []string) int {
	grid := [1000][1000]bool{}

	for _, instructionString := range instructions {
		ins := toInstruction(instructionString)

		for row := ins.startRow; row <= ins.endRow; row++ {
			for col := ins.startCol; col <= ins.endCol; col++ {
				if ins.command == on {
					grid[row][col] = true
				} else if ins.command == off {
					grid[row][col] = false
				} else if ins.command == toggle {
					grid[row][col] = !grid[row][col]
				}
			}
		}
	}

	return countOnLights(grid)
}

func countOnLights(grid [1000][1000]bool) int {
	numLightsOn := 0
	for row := range grid {
		for _, light := range grid[row] {
			if light {
				numLightsOn++
			}
		}
	}
	return numLightsOn
}

func toInstruction(instructionString string) instruction {
	match := regex.FindStringSubmatch(instructionString)
	ins := instruction{
		command:  instructionCommand(match[1]),
		startCol: conversion.StringToInt(match[2]),
		startRow: conversion.StringToInt(match[3]),
		endCol:   conversion.StringToInt(match[4]),
		endRow:   conversion.StringToInt(match[5]),
	}
	return ins
}

func main() {
	instructions, _ := file.ReadFile("./2015/day6/input.txt")
	fmt.Println(setupLights(instructions))
}
