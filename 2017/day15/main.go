package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
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
const modulo = 2147483647
const first16Bits = 0xffff

func part1(input []string) interface{} {
	generatorA, generatorB := parseInput(input)

	iterations := 40_000_000
	numEqual := 0

	for i := 0; i < iterations; i++ {
		generatorA = (generatorA * factorA) % modulo
		generatorB = (generatorB * factorB) % modulo

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

	abortA := make(chan int)
	generatorA := &Generator{
		Name:        "A",
		StartValue:  startValueA,
		Factor:      factorA,
		Denominator: 4,
		Abort:       abortA,
	}
	abortB := make(chan int)
	generatorB := &Generator{
		Name:        "B",
		StartValue:  startValueB,
		Factor:      factorB,
		Denominator: 8,
		Abort:       abortB,
	}

	channelA := generatorA.Generate()
	channelB := generatorB.Generate()

	for i := 0; i < iterations; i++ {
		generatorAValue := <-channelA
		generatorBValue := <-channelB

		if first16BitsAreEqual(generatorAValue, generatorBValue) {
			numEqual++
		}
	}

	close(abortA)
	close(abortB)

	return numEqual
}

func first16BitsAreEqual(generatorA int, generatorB int) bool {
	valueA := generatorA & first16Bits
	valueB := generatorB & first16Bits
	return valueA == valueB
}

type Generator struct {
	Name        string
	StartValue  int
	Factor      int
	Denominator int
	Abort       <-chan int
}

func (g *Generator) Generate() <-chan int {
	channel := make(chan int)
	value := g.StartValue

	go func() {
		defer close(channel)
		for i := 0; ; i++ {
			select {
			case <-g.Abort:
				return
			default:
				value = (value * g.Factor) % modulo
				if (value % g.Denominator) == 0 {
					channel <- value
				}
			}
		}
	}()

	return channel
}

func parseInput(input []string) (int, int) {
	lineA := strings.Split(input[0], " ")
	generatorA := conversion.StringToInt(lineA[len(lineA)-1])
	lineB := strings.Split(input[1], " ")
	generatorB := conversion.StringToInt(lineB[len(lineB)-1])
	return generatorA, generatorB
}
