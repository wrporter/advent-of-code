package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"sort"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 10
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	adapters, _ := convert.ToInts(input)
	sort.Ints(adapters)

	builtInAdapter := getBuiltInAdapter(adapters)
	result, _ := jolt(adapters, builtInAdapter)

	return result[1] * result[3]
}

func part2(input []string) interface{} {
	adapters, _ := convert.ToInts(input)
	sort.Ints(adapters)
	builtInAdapter := getBuiltInAdapter(adapters)
	adapters = append(adapters, builtInAdapter)

	// Dynamic Programming memoization
	combinations := map[int]int{0: 1}
	for _, adapter := range adapters {
		combinations[adapter] = combinations[adapter-1] + combinations[adapter-2] + combinations[adapter-3]
	}

	// - Can we solve this by calculating the number of combinations in a range?
	// - Does not work due to neglecting combinations based on the positioning of each element.
	// - Maybe we could do this by somehow accounting for those other combinations?
	//min := []int{0}
	//prev := 0
	//for _, adapter := range adapters {
	//	if adapter-min[len(min)-1] > 3 {
	//		min = append(min, prev)
	//	}
	//	prev = adapter
	//}
	//numElementRemoved := len(adapters) - len(min)
	//fmt.Println(nCrRange(numElementRemoved, 0, numElementRemoved))

	// Brute force way that takes too long for large data sets
	//count := 0
	//probability.ComboSize(adapters, 1, len(adapters), func(i []int) {
	//	if _, passed := jolt(i, builtInAdapter); passed {
	//		count++
	//	}
	//})
	//return count

	return combinations[adapters[len(adapters)-1]]
}

func getBuiltInAdapter(adapter []int) int {
	builtInAdapter := 0
	for _, rating := range adapter {
		builtInAdapter = ints.Max(builtInAdapter, rating)
	}
	return builtInAdapter + 3
}

func jolt(adapters []int, builtInAdapter int) (map[int]int, bool) {
	adapters = append(adapters, builtInAdapter)

	result := make(map[int]int)
	j := 0
	for _, rating := range adapters {
		diff := ints.Abs(j - rating)
		if diff <= 3 {
			result[diff]++
			j = rating
		} else {
			return nil, false
		}
	}
	return result, true
}

func nCrRange(n, start, end int) int {
	result := 0
	for i := start; i <= end; i++ {
		result += nCr(n, i)
	}
	return result
}

func nPrRange(n, start, end int) int {
	result := 0
	for i := start; i <= end; i++ {
		result += nPr(n, i)
	}
	return result
}

func nPr(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func nCr(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func factorial(n int) int {
	fact := 1
	i := 1
	for i <= n {
		fact *= i
		i++
	}
	return fact
}
