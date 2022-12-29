package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"sort"
	"strings"
)

func (s Solution) Part1(input string, args ...interface{}) interface{} {
	elves := strings.Split(input, "\n\n")
	max := 0

	for _, elf := range elves {
		total := getTotal(elf)
		max = ints.Max(max, total)
	}

	return max
}

func (s Solution) Part2(input string, args ...interface{}) interface{} {
	elves := strings.Split(input, "\n\n")
	totals := make([]int, len(elves))

	for i, elf := range elves {
		total := getTotal(elf)
		totals[i] = total
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	return totals[0] + totals[1] + totals[2]
}

func getTotal(elf string) int {
	food := strings.Split(elf, "\n")
	total := 0
	for _, item := range food {
		total += convert.StringToInt(item)
	}
	return total
}
