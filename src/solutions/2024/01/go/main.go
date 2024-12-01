package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/mymath"
	"slices"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	list1 := make([]int, 0, len(lines))
	list2 := make([]int, 0, len(lines))

	for _, pairStr := range lines {
		pair := convert.ToIntsV2(strings.Fields(pairStr))
		list1 = insertSorted(list1, pair[0])
		list2 = insertSorted(list2, pair[1])
	}

	totalDistance := 0
	for i := range list1 {
		totalDistance += mymath.Abs(list1[i] - list2[i])
	}
	return totalDistance
}

func part2(input string, _ ...interface{}) interface{} {
	lines := strings.Split(input, "\n")
	list1 := make([]int, 0, len(lines))
	counts2 := make(map[int]int)

	for _, pairStr := range lines {
		pair := convert.ToIntsV2(strings.Fields(pairStr))
		num1, num2 := pair[0], pair[1]
		list1 = append(list1, num1)
		counts2[num2] += 1
	}

	similarityScore := 0
	for _, num := range list1 {
		similarityScore += num * counts2[num]
	}
	return similarityScore
}

func insertSorted(list []int, value int) []int {
	index := sort.SearchInts(list, value)
	return slices.Insert(list, index, value)
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 1, Part1: part1, Part2: part2}
}
