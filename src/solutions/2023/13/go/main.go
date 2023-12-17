package main

import (
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	return summarize(input, false)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	return summarize(input, true)
}

func summarize(input string, withSmudge bool) interface{} {
	sum := 0
	parse(input, func(rows, cols []int) {
		sum += 100 * indexOfMirror(rows, withSmudge)
		sum += indexOfMirror(cols, withSmudge)
	})
	return sum
}

func indexOfMirror(pattern []int, withSmudge bool) int {
	for i := 0; i < len(pattern)-1; i++ {
		smudged := false
		reflected := true
		prev := i
		next := i + 1

		for prev >= 0 && next < len(pattern) && reflected {
			if withSmudge && !smudged && oneBitDiffers(pattern[prev], pattern[next]) {
				smudged = true
			} else {
				reflected = pattern[prev] == pattern[next]
			}
			prev--
			next++
		}

		if reflected && (!withSmudge || (withSmudge && smudged)) {
			return i + 1
		}
	}

	return 0
}

func oneBitDiffers(a int, b int) bool {
	return isPowerOfTwo(a ^ b)
}

func isPowerOfTwo(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

func parse(input string, process func(rows, cols []int)) {
	for _, chunk := range strings.Split(input, "\n\n") {
		grid := strings.Split(chunk, "\n")
		height := len(grid)
		width := len(grid[0])
		rows := make([]int, height)
		cols := make([]int, width)

		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				if grid[row][col] == '#' {
					rows[row] |= 1 << col
				}
			}
		}

		for col := 0; col < width; col++ {
			for row := 0; row < height; row++ {
				if grid[row][col] == '#' {
					cols[col] |= 1 << row
				}
			}
		}

		process(rows, cols)
	}
}

func main() {
	Run()
}
