package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
)

func countPermutations(values []int, target int) (int, int) {
	count := 0
	var permute func([]int, int, int)
	sizes := make(map[int]int)
	min := ints.MaxInt

	permute = func(current []int, target int, index int) {
		if target < 0 {
			return
		}

		if target == 0 {
			size := len(current)
			sizes[size]++
			min = ints.Min(min, size)

			count++
			return
		}

		for i := index; i < len(values); i++ {
			next := ints.Copy(current)
			next = append(next, values[i])
			permute(next, target-values[i], i+1)
		}
	}
	permute(nil, target, 0)

	return count, sizes[min]
}

func main() {
	lines, _ := file.ReadFile("./2015/day17/input.txt")
	containers, _ := conversion.ToInts(lines)
	fmt.Println(countPermutations([]int{20, 15, 10, 5, 5}, 25))
	fmt.Println(countPermutations(containers, 150))
}
