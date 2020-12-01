package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/probability"
)

func main() {
	probability.PermuteSize([]int{1, 2, 3}, 2, 2, func(ints []int) {
		fmt.Println(ints)
	})
	fmt.Println("Channels")

	output := make(chan []int, 1)
	go PermuteSize([]int{1, 2, 3}, 2, 2, output)
	for values := range output {
		fmt.Println(values)
	}
}

func PermuteSize(values []int, startSize int, endSize int, emit chan<- []int) {
	var permuteSize func([]int, int, int)

	permuteSize = func(current []int, index int, size int) {
		if len(current) == size {
			emit <- ints.Copy(current)
			return
		}

		for i := index; i < len(values); i++ {
			current = append(current, values[i])
			permuteSize(current, i+1, size)
			current = current[:len(current)-1]
		}
	}

	for size := startSize; size <= endSize; size++ {
		permuteSize(nil, 0, size)
	}

	close(emit)
}
