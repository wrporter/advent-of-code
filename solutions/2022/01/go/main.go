package main

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/solution"
	"sort"
	"strings"
)

func main() {
	s := Solution{AbstractSolution: solution.AbstractSolution{
		Solution: Solution{},
		Year:     2022,
		Day:      1,
	}}

	s.Run([]interface{}{}, []interface{}{})
}

type Solution struct {
	solution.AbstractSolution
}

func (s Solution) Part1(input string, args ...interface{}) interface{} {
	lists := strings.Split(input, "\n\n")
	max := 0

	for _, listString := range lists {
		total := getTotal(listString)
		max = ints.Max(max, total)
	}

	return max
}

func (s Solution) Part2(input string, args ...interface{}) interface{} {
	lists := strings.Split(input, "\n\n")
	var totals []int

	for _, listString := range lists {
		total := getTotal(listString)
		totals = append(totals, total)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	return totals[0] + totals[1] + totals[2]
}

func getTotal(listString string) int {
	calorieStrings := strings.Split(listString, "\n")
	total := 0
	for _, calorieString := range calorieStrings {
		total += convert.StringToInt(calorieString)
	}
	return total
}
