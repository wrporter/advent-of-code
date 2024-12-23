package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"fmt"
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

func setupLightBrightness(instructions []string) int {
	grid := [1000][1000]int{}

	for _, instructionString := range instructions {
		ins := toInstruction(instructionString)

		for row := ins.startRow; row <= ins.endRow; row++ {
			for col := ins.startCol; col <= ins.endCol; col++ {
				if ins.command == on {
					grid[row][col] += 1
				} else if ins.command == off {
					grid[row][col] = ints.Max(0, grid[row][col]-1)
				} else if ins.command == toggle {
					grid[row][col] += 2
				}
			}
		}
	}

	return getTotalBrightness(grid)
}

func getTotalBrightness(grid [1000][1000]int) int {
	total := 0
	for row := range grid {
		for _, brightness := range grid[row] {
			total += brightness
		}
	}
	return total
}

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
		startCol: convert.StringToInt(match[2]),
		startRow: convert.StringToInt(match[3]),
		endCol:   convert.StringToInt(match[4]),
		endRow:   convert.StringToInt(match[5]),
	}
	return ins
}

func main() {
	instructions, _ := file.ReadFile("./2015/day6/input.txt")
	fmt.Println(setupLights(instructions))
	fmt.Println(setupLightBrightness(instructions))
}
