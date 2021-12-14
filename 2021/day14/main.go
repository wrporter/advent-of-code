package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 14
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	polymer := parseInput(input)

	for s := 1; s <= 10; s++ {
		polymer.step()
	}

	min, max := polymer.getMinMax()
	return max - min
}

func part2(input []string) interface{} {
	polymer := parseInput(input)

	for s := 1; s <= 40; s++ {
		polymer.step()
		min, max := polymer.getMinMax()
		fmt.Println(s, min, max, max-min)
	}

	min, max := polymer.getMinMax()
	return max - min
}

func (p *Polymer) getMinMax() (int, int) {
	max := 0
	min := math.MaxInt
	for _, count := range p.counts {
		max = ints.Max(max, count)
		min = ints.Min(min, count)
	}
	return min, max
}

func (p *Polymer) step() {
	numInsertions := 0
	next := p.template
	for i := 0; i < len(p.template)-1; i++ {
		part := p.template[i : i+2]

		if result, ok := p.rules[part]; ok {
			index := i + numInsertions + 1
			next = next[:index] + result + next[index:]

			numInsertions++
			p.counts[result]++
		}
	}
	p.template = next
}

func parseInput(input []string) *Polymer {
	template := input[0]
	rules := make(map[string]string)
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	counts := make(map[string]int)
	for _, char := range template {
		counts[string(char)]++
	}

	return &Polymer{
		template: template,
		rules:    rules,
		counts:   counts,
	}
}

type Polymer struct {
	template string
	rules    map[string]string
	counts   map[string]int
}
