package main

import (
	"aoc/src/lib/go/file"
	"fmt"
)

const (
	Up   = '('
	Down = ')'
)

func elevator(instructions string) int {
	floor := 0
	for _, instruction := range instructions {
		if instruction == Up {
			floor++
		} else if instruction == Down {
			floor--
		}
	}
	return floor
}

func firstTimeAt(instructions string, stop int) int {
	floor := 0
	for characterPosition, instruction := range instructions {
		if instruction == Up {
			floor++
		} else if instruction == Down {
			floor--
		}
		if floor == stop {
			return characterPosition + 1
		}
	}
	return -1
}

func main() {
	lines, _ := file.ReadFile("./2015/day1/input.txt")
	input := lines[0]
	fmt.Println(elevator(input))
	fmt.Println(firstTimeAt(input, -1))
}
