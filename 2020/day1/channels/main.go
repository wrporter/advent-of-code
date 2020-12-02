package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/probability"
)

func main() {
	probability.ComboSize([]int{1, 2, 3}, 2, 2, func(ints []int) {
		fmt.Println(ints)
	})
	fmt.Println("Channels")

	for values := range ComboSize([]int{1, 2, 3}, 2, 2) {
		fmt.Println(values)
	}
}

func ComboSize(values []int, startSize int, endSize int) <-chan []int {
	var permuteSize func([]int, int, int)
	emit := make(chan []int, 1)

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

		if index == 0 && size == startSize {
			close(emit)
		}
	}

	for size := startSize; size <= endSize; size++ {
		go permuteSize(nil, 0, size)
	}

	return emit
}
