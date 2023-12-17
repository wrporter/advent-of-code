package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 15
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	discs := parseInput(input)
	return findTimeForCapsuleToComeOut(discs)
}

func part2(input []string) interface{} {
	discs := parseInput(input)
	discs = append(discs, &Disc{
		id:        len(discs) + 1,
		positions: 11,
		time:      0,
		position:  0,
	})
	residues, moduli := make([]int, len(discs)), make([]int, len(discs))
	for i, d := range discs {
		residues[i] = (d.positions - d.position - d.id) % d.positions
		moduli[i] = d.positions
	}
	return ints.ChineseRemainderTheorem(residues, moduli)
}

func findTimeForCapsuleToComeOut(discs []*Disc) int {
	for t := 0; t < 10_000_000; t++ {
		capsuleComesOut := true
		for step, disc := range discs {
			if !disc.isAtSlot0(t + step + 1) {
				capsuleComesOut = false
				break
			}
		}
		if capsuleComesOut {
			return t
		}
	}
	return -1
}

var regex = regexp.MustCompile(`^Disc #(\d+) has (\d+) positions; at time=(\d+), it is at position (\d+)\.$`)

func parseInput(input []string) []*Disc {
	discs := make([]*Disc, len(input))
	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		discs[i] = &Disc{
			id:        convert.StringToInt(match[1]),
			positions: convert.StringToInt(match[2]),
			time:      convert.StringToInt(match[3]),
			position:  convert.StringToInt(match[4]),
		}
	}
	return discs
}

type Disc struct {
	id        int
	positions int
	time      int
	position  int
}

func (d *Disc) isAtSlot0(time int) bool {
	return (d.position+time)%d.positions == 0
}
