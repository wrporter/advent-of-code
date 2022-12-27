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
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	p := parseInput(input)

	for s := 1; s <= 10; s++ {
		p.step()
	}

	min, max := p.getMinMax()
	return max - min
}

func part2(input []string) interface{} {
	p := parseInput(input)

	for s := 1; s <= 40; s++ {
		p.step()
	}

	min, max := p.getMinMax()
	return max - min
}

func (p *polymer) step() {
	for pair, count := range copyMap(p.pairs) {
		insert := p.rules[pair]
		p.pairs[pair] -= count
		p.pairs[pair[:1]+insert] += count
		p.pairs[insert+pair[1:]] += count
		p.elements[insert] += count
	}
}

func (p *polymer) getMinMax() (int, int) {
	min := math.MaxInt
	max := 0
	for _, count := range p.elements {
		min = ints.Min(min, count)
		max = ints.Max(max, count)
	}
	return min, max
}

func parseInput(input []string) *polymer {
	template := input[0]
	rules := make(map[string]string)
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	elements := make(map[string]int)
	pairs := make(map[string]int)
	for i, char := range template {
		elements[string(char)]++
		if i < len(template)-1 {
			pairs[template[i:i+2]]++
		}
	}

	return &polymer{
		rules:    rules,
		elements: elements,
		pairs:    pairs,
	}
}

type polymer struct {
	rules    map[string]string
	elements map[string]int
	pairs    map[string]int
}

func copyMap(m map[string]int) map[string]int {
	result := make(map[string]int)
	for key, value := range m {
		result[key] = value
	}
	return result
}
