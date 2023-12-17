package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"fmt"
)

func main() {
	year, day := 2015, 25
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	r := 2947
	c := 3029
	target := (r+c-1)*(r+c-2)/2 + c - 1

	return (20151125 * modExp(252533, target, 33554393)) % 33554393
}

func part2(input []string) interface{} {
	return "We did it! Merry Christmas!"
}

func modExp(b, e, m int) int {
	r, x := 1, 1
	for e != 0 {
		r = e % 2
		e /= 2
		if r == 1 {
			x = (x * b) % m
		}
		b = (b * b) % m
	}
	return x
}
