package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/probability"
)

func main() {
	probability.PermuteInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 14, func(ints []int) {
		fmt.Println(ints)
	})
}
