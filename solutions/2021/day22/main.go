package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/sample-input2.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	steps := parseInput(input)
	on := make(map[[3]int]struct{})

	for _, s := range steps {
		for x := ints.Max(-50, s.x.from); x <= ints.Min(50, s.x.to); x++ {
			for y := ints.Max(-50, s.y.from); y <= ints.Min(50, s.y.to); y++ {
				for z := ints.Max(-50, s.z.from); z <= ints.Min(50, s.z.to); z++ {
					point := [3]int{x, y, z}
					if s.state > 0 {
						on[point] = struct{}{}
					} else {
						delete(on, point)
					}
				}
			}
		}
	}

	return len(on)
}

func part2(input []string) interface{} {
	steps := parseInput(input)

	cubes := make(map[cuboid]int)
	for _, s := range steps {

		update := make(map[cuboid]int)
		for c, state := range cubes {
			intersection := s.intersect(c)
			if intersection != nil {
				update[*intersection] -= state
			}
		}

		if s.state > 0 {
			update[s.cuboid] += s.state
		}

		for c, state := range update {
			cubes[c] += state
		}
	}

	return count(cubes)
}

func count(cubes map[cuboid]int) interface{} {
	sum := 0
	for c, state := range cubes {
		sum += c.volume() * state
	}
	return sum
}

var regex = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

func parseInput(input []string) []step {
	steps := make([]step, len(input))
	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		ranges, _ := convert.ToInts(match[2:])
		state := 1
		if match[1] == "off" {
			state = -1
		}
		steps[i] = step{
			state: state,
			cuboid: cuboid{
				x: newCRange(ranges[0], ranges[1]),
				y: newCRange(ranges[2], ranges[3]),
				z: newCRange(ranges[4], ranges[5]),
			},
		}
	}
	return steps
}

type step struct {
	state int
	cuboid
}

type crange struct {
	from, to int
}

func newCRange(a, b int) crange {
	return crange{
		from: ints.Min(a, b),
		to:   ints.Max(a, b),
	}
}

type cuboid struct {
	x, y, z crange
}

func (c cuboid) volume() int {
	return (c.x.to - c.x.from + 1) * (c.y.to - c.y.from + 1) * (c.z.to - c.z.from + 1)
}

func (c cuboid) intersect(o cuboid) *cuboid {
	x0 := ints.Max(c.x.from, o.x.from)
	x1 := ints.Min(c.x.to, o.x.to)
	y0 := ints.Max(c.y.from, o.y.from)
	y1 := ints.Min(c.y.to, o.y.to)
	z0 := ints.Max(c.z.from, o.z.from)
	z1 := ints.Min(c.z.to, o.z.to)

	if x0 <= x1 && y0 <= y1 && z0 <= z1 {
		return &cuboid{
			x: newCRange(x0, x1),
			y: newCRange(y0, y1),
			z: newCRange(z0, z1),
		}
	}

	return nil
}
