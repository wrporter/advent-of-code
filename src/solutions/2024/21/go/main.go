package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	return sumComplexities(input, 4)
}

func part2(input string, _ ...interface{}) interface{} {
	return sumComplexities(input, 27)
}

var numericKeypad = []string{
	"789",
	"456",
	"123",
	" 0A",
}

var directionalKeypad = []string{
	" ^A",
	"<v>",
}

type cacheKey struct {
	code string
	area int
}

func sumComplexities(input string, numKeypads int) interface{} {
	total := 0
	cache := make(map[cacheKey]int)
	for _, code := range strings.Split(input, "\n") {
		value, _ := strconv.Atoi(strings.ReplaceAll(code, "A", ""))
		total += getMinNumberKeyPresses(cache, code, 1, numKeypads) * value
	}
	return total
}

func getMinNumberKeyPresses(cache map[cacheKey]int, code string, area int, numKeypads int) int {
	if area == numKeypads {
		return len(code)
	}

	key := cacheKey{code, area}
	if sum, seen := cache[key]; seen {
		return sum
	}

	keypad := directionalKeypad
	if area == 1 {
		keypad = numericKeypad
	}

	sum := 0
	code = "A" + code
	for i := 0; i < len(code)-1; i++ {
		sequence := getShortestSequence(keypad, rune(code[i]), rune(code[i+1]))
		sum += getMinNumberKeyPresses(cache, sequence, area+1, numKeypads)
	}

	cache[key] = sum
	return sum
}

func getShortestSequence(keypad []string, fromKey rune, toKey rune) string {
	from := findPosition(keypad, fromKey)
	to := findPosition(keypad, toKey)

	var dfs func(x, y int, sequence string) []string
	dfs = func(x, y int, sequence string) []string {
		if x == to.X && y == to.Y {
			return []string{sequence + "A"}
		}

		var results []string
		if to.X < x && keypad[y][x-1] != ' ' {
			results = append(results, dfs(x-1, y, sequence+"<")...)
		}
		if to.Y < y && keypad[y-1][x] != ' ' {
			results = append(results, dfs(x, y-1, sequence+"^")...)
		}
		if to.Y > y && keypad[y+1][x] != ' ' {
			results = append(results, dfs(x, y+1, sequence+"v")...)
		}
		if to.X > x && keypad[y][x+1] != ' ' {
			results = append(results, dfs(x+1, y, sequence+">")...)
		}
		return results
	}

	sequences := dfs(from.X, from.Y, "")
	best := sequences[0]
	minChanges := countChanges(best)

	for _, sequence := range sequences {
		if changes := countChanges(sequence); changes < minChanges {
			best = sequence
			minChanges = changes
		}
	}

	return best
}

func findPosition(keypad []string, value rune) geometry.Point {
	for y, row := range keypad {
		for x, char := range row {
			if char == value {
				return geometry.Point{x, y}
			}
		}
	}
	return geometry.Point{-1, -1}
}

func countChanges(sequence string) int {
	changes := 0
	for i := 1; i < len(sequence); i++ {
		if sequence[i] != sequence[i-1] {
			changes++
		}
	}
	return changes
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 21, Part1: part1, Part2: part2}
}
