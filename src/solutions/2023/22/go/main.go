package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"regexp"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	bricks := parse(input)
	drop(bricks)
	stable := 0

	for i := range bricks {
		next := clone(bricks)
		next = append(next[:i], next[i+1:]...)
		fallen := drop(next)

		if fallen == 0 {
			stable++
		}
	}

	return stable
}

func part2(input string, _ ...interface{}) interface{} {
	bricks := parse(input)
	drop(bricks)
	fallen := 0

	for i := range bricks {
		next := clone(bricks)
		next = append(next[:i], next[i+1:]...)
		fallen += drop(next)
	}

	return fallen
}

func drop(bricks []*brick) int {
	peaks := make(map[point]int)
	fallen := 0

	for _, b := range bricks {
		area := make(map[point]bool)
		peak := 0

		for x := b.sx; x <= b.ex; x++ {
			for y := b.sy; y <= b.ey; y++ {
				p := point{x, y}
				area[p] = true
				peak = max(peak, peaks[p])
			}
		}

		peak += 1

		for p := range area {
			peaks[p] = peak + b.ez - b.sz
		}

		fall := b.sz - peak
		b.sz -= fall
		b.ez -= fall

		if fall > 0 {
			fallen++
		}
	}

	return fallen
}

func clone(s []*brick) []*brick {
	result := make([]*brick, len(s))
	for i, e := range s {
		result[i] = &brick{e.sx, e.sy, e.sz, e.ex, e.ey, e.ez}
	}
	return result
}

func main() {
	New().Run(nil, nil)
}

var brickRegex = regexp.MustCompile(`\d+`)

func parse(input string) []*brick {
	lines := strings.Split(input, "\n")
	bricks := make([]*brick, len(lines))

	for i, line := range lines {
		nums := convert.ToIntsV2(brickRegex.FindAllString(line, 6))
		bricks[i] = &brick{
			sx: nums[0], sy: nums[1], sz: nums[2],
			ex: nums[3], ey: nums[4], ez: nums[5],
		}
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i].sz < bricks[j].sz
	})

	return bricks
}

type point struct{ x, y int }
type brick struct{ sx, sy, sz, ex, ey, ez int }

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 22, Part1: part1, Part2: part2}
}
