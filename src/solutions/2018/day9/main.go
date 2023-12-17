package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"container/ring"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 9
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numPlayers, lastMarble := parseInput(input)
	return playMarbles(numPlayers, lastMarble)
}

func part2(input []string) interface{} {
	numPlayers, lastMarble := parseInput(input)
	return playMarbles(numPlayers, lastMarble*100)
}

func parseInput(input []string) (int, int) {
	parts := strings.Fields(input[0])
	numPlayers := convert.StringToInt(parts[0])
	lastMarble := convert.StringToInt(parts[6])
	return numPlayers, lastMarble
}

func playMarbles(numPlayers int, lastMarble int) int {
	scores := make(map[int]int)
	circle := newRing(0)
	maxScore := 0

	for marble := 1; marble <= lastMarble; marble++ {
		if marble%23 == 0 {
			circle = circle.Move(-8)
			removedMarble := circle.Unlink(1).Value.(int)
			circle = circle.Next()
			player := marble % numPlayers
			scores[player] += marble + removedMarble
			maxScore = ints.Max(maxScore, scores[player])
		} else {
			circle = circle.Next()
			circle.Link(newRing(marble))
			circle = circle.Next()
		}
	}

	return maxScore
}

func newRing(v int) *ring.Ring {
	r := ring.New(1)
	r.Value = v
	return r
}
