package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"fmt"
)

func main() {
	year, day := 2015, 24
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	weights, _ := convert.ToInts(input)
	target := ints.Sum(weights) / 3

	smallest := ints.Copy(weights)
	ComboSize(weights, 1, len(weights), func(i []int) bool {
		if len(i) < len(smallest) && ints.Sum(i) == target {
			smallest = i
			return true
		}
		return false
	})

	return ints.Product(smallest)
}

func part2(input []string) interface{} {
	weights, _ := convert.ToInts(input)
	target := ints.Sum(weights) / 4

	smallest := ints.Copy(weights)
	minQuantumEntanglement := ints.MaxInt
	ComboSize(weights, 1, len(weights), func(i []int) bool {
		if len(i) > len(smallest) {
			return true
		}
		if ints.Sum(i) == target && ints.Product(i) < minQuantumEntanglement {
			smallest = i
			minQuantumEntanglement = ints.Product(i)
			return true
		}
		return false
	})

	return ints.Product(smallest)
}

func ComboSize(values []int, startSize int, endSize int, emit func([]int) bool) {
	var permuteSize func([]int, int, int)

	stop := false
	permuteSize = func(current []int, index int, size int) {
		if len(current) == size {
			stop = emit(current)
			return
		}

		for i := index; i < len(values); i++ {
			if stop {
				return
			}

			current = append(current, values[i])
			permuteSize(current, i+1, size)
			current = current[:len(current)-1]
		}
	}

	for size := startSize; size <= endSize; size++ {
		if stop {
			return
		}
		permuteSize(nil, 0, size)
	}
}
