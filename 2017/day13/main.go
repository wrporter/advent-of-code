package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 13
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	maxDepth, layers, scanners := parseInput(input)

	return getSeverityOfTravel(maxDepth, scanners, layers)
}

func part2(input []string) interface{} {
	_, layers, _ := parseInput(input)

	for delay := 0; delay <= 1_000_000_000; delay++ {
		caught := false
		for depth, layerRange := range layers {
			if isCaught(depth, delay, layerRange) {
				caught = true
				break
			}
		}

		if !caught {
			return delay
		}
	}

	return -1
}

func isCaught(depth int, delay int, layerRange int) bool {
	return (depth+delay)%(2*(layerRange-1)) == 0
}

func getSeverityOfTravel(maxDepth int, scanners map[int]int, layers map[int]int) (severity int) {
	for depth := 0; depth <= maxDepth; depth++ {
		if scannerPosition, ok := scanners[depth]; ok && scannerPosition == 0 {
			severity += depth * layers[depth]
		}

		setScannerPositions(scanners, layers, depth+1)
	}

	return severity
}

func setScannerPositions(scanners map[int]int, layers map[int]int, time int) {
	for depth := range scanners {
		layerRange := layers[depth]
		offset := time % ((layerRange - 1) * 2)

		position := offset
		if offset > layerRange-1 {
			position = 2*(layerRange-1) - offset
		}

		scanners[depth] = position
	}
}

func parseInput(input []string) (int, map[int]int, map[int]int) {
	maxDepth := 0
	layers := make(map[int]int)
	scanners := make(map[int]int)
	for _, line := range input {
		chunks := strings.Split(line, ": ")
		depth := conversion.StringToInt(chunks[0])
		layerRange := conversion.StringToInt(chunks[1])

		layers[depth] = layerRange
		scanners[depth] = 0

		if depth > maxDepth {
			maxDepth = depth
		}
	}
	return maxDepth, layers, scanners
}
