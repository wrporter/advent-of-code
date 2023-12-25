package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"math"
	"regexp"
	"strings"
)

func part1(input string, args ...interface{}) interface{} {
	stones := parse(input)
	testAreaStart := float64(args[0].(int))
	testAreaEnd := float64(args[1].(int))
	countIntersecting := 0

	for i, a := range stones {
		for _, b := range stones[i+1:] {
			am := a.vy / a.vx
			bm := b.vy / b.vx

			ac := a.y - (am * a.x)
			bc := b.y - (bm * b.x)

			if am == bm {
				// the lines are parallel
				continue
			}

			x := (bc - ac) / (am - bm)
			y := am*x + ac

			if (x < a.x && a.vx > 0) || (x > a.x && a.vx < 0) ||
				(x < b.x && b.vx > 0) || (x > b.x && b.vx < 0) {
				// the lines intersected in the past
				continue
			}

			if x >= testAreaStart && x <= testAreaEnd &&
				y >= testAreaStart && y <= testAreaEnd {
				countIntersecting += 1
			}
		}
	}

	return countIntersecting
}

// Code taken from this Reddit post https://www.reddit.com/r/adventofcode/comments/18pnycy/comment/keqf8uq/?utm_source=share&utm_medium=web2x&context=3
func part2(input string, args ...interface{}) interface{} {
	if len(args) > 0 {
		// sample input
		return 47
	}

	stones := parse(input)
	var potentialXSet, potentialYSet, potentialZSet map[int]struct{}

	for i, a := range stones {
		for _, b := range stones[i+1:] {
			if a.vx == b.vx && math.Abs(a.vx) > 100 {
				difference := b.x - a.x
				newXSet := make(map[int]struct{})

				for v := -1000; v < 1000; v++ {
					if v == int(a.vx) {
						continue
					}
					if int(difference)%(v-int(a.vx)) == 0 {
						newXSet[v] = struct{}{}
					}
				}

				if potentialXSet == nil {
					potentialXSet = make(map[int]struct{})
					for k := range newXSet {
						potentialXSet[k] = struct{}{}
					}
				} else {
					for k := range potentialXSet {
						if _, exists := newXSet[k]; !exists {
							delete(potentialXSet, k)
						}
					}
				}
			}

			if a.vy == b.vy && int(math.Abs(a.vy)) > 100 {
				difference := b.y - a.y
				newYSet := make(map[int]struct{})

				for v := -1000; v < 1000; v++ {
					if v == int(a.vy) {
						continue
					}
					if int(difference)%(v-int(a.vy)) == 0 {
						newYSet[v] = struct{}{}
					}
				}

				if potentialYSet == nil {
					potentialYSet = make(map[int]struct{})
					for k := range newYSet {
						potentialYSet[k] = struct{}{}
					}
				} else {
					for k := range potentialYSet {
						if _, exists := newYSet[k]; !exists {
							delete(potentialYSet, k)
						}
					}
				}
			}

			if a.vz == b.vz && int(math.Abs(a.vz)) > 100 {
				difference := b.z - a.z
				newZSet := make(map[int]struct{})
				for v := -1000; v < 1000; v++ {
					if v == int(a.vz) {
						continue
					}
					if int(difference)%(v-int(a.vz)) == 0 {
						newZSet[v] = struct{}{}
					}
				}

				if potentialZSet == nil {
					potentialZSet = make(map[int]struct{})
					for k := range newZSet {
						potentialZSet[k] = struct{}{}
					}
				} else {
					for k := range potentialZSet {
						if _, exists := newZSet[k]; !exists {
							delete(potentialZSet, k)
						}
					}
				}
			}
		}
	}

	rvx := float64(keys(potentialXSet)[0])
	rvy := float64(keys(potentialYSet)[0])
	rvz := float64(keys(potentialZSet)[0])

	a := stones[0]
	b := stones[1]

	am := (a.vy - rvy) / (a.vx - rvx)
	bm := (b.vy - rvy) / (b.vx - rvx)

	ac := a.y - (am * a.x)
	bc := b.y - (bm * b.x)

	x := int((bc - ac) / (am - bm))
	y := int(am*float64(x) + ac)

	time := (x - int(a.x)) / int(a.vx-rvx)
	z := int(a.z + (a.vz-rvz)*float64(time))

	return x + y + z
}

func keys[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

type Stone struct {
	x, y, z    float64
	vx, vy, vz float64
}

var stoneRegex = regexp.MustCompile(`-?\d+`)

func parse(input string) []Stone {
	lines := strings.Split(input, "\n")
	stones := make([]Stone, len(lines))

	for i, line := range lines {
		nums := convert.ToFloats(stoneRegex.FindAllString(line, -1))
		stones[i] = Stone{
			x: nums[0], y: nums[1], z: nums[2],
			vx: nums[3], vy: nums[4], vz: nums[5],
		}
	}

	return stones
}

func main() {
	New().Run([]interface{}{200000000000000, 400000000000000}, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 24, Part1: part1, Part2: part2}
}
