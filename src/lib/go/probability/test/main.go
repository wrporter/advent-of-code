package main

import (
	"aoc/src/lib/go/probability"
	"fmt"
)

func main() {
	probability.PermuteInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 14, func(ints []int) {
		fmt.Println(ints)
	})
}
