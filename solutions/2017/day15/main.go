package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 15
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

const factorA = 16807
const factorB = 48271
const modulus = 2147483647
const first16Bits = 0xffff

func part1(input []string) interface{} {
	generatorA, generatorB := parseInput(input)

	iterations := 40_000_000
	numEqual := 0

	for i := 0; i < iterations; i++ {
		generatorA = (generatorA * factorA) % modulus
		generatorB = (generatorB * factorB) % modulus

		if first16BitsAreEqual(generatorA, generatorB) {
			numEqual++
		}
	}

	return numEqual
}

func part2(input []string) interface{} {
	startValueA, startValueB := parseInput(input)

	iterations := 5_000_000
	numEqual := 0

	a, b := startValueA, startValueB

	for i := 0; i < iterations; i++ {
		for {
			a = (a * factorA) % modulus
			if a%4 == 0 {
				break
			}
		}

		for {
			b = (b * factorB) % modulus
			if b%8 == 0 {
				break
			}
		}

		if first16BitsAreEqual(a, b) {
			numEqual++
		}
	}

	return numEqual
}

func first16BitsAreEqual(generatorA int, generatorB int) bool {
	valueA := generatorA & first16Bits
	valueB := generatorB & first16Bits
	return valueA == valueB
}

func parseInput(input []string) (int, int) {
	lineA := strings.Split(input[0], " ")
	generatorA := convert.StringToInt(lineA[len(lineA)-1])
	lineB := strings.Split(input[1], " ")
	generatorB := convert.StringToInt(lineB[len(lineB)-1])
	return generatorA, generatorB
}
