package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/ints"
)

func getPresents(house int) int {
	presents := 0

	for elf := 1; elf <= ints.Sqrt(house); elf++ {
		if house%elf == 0 {
			presents += elf
			if house/elf != elf {
				presents += house / elf
			}
		}
	}

	return presents * 10
}

func getHouseWith(presents int) int {
	for house := 1; ; house++ {
		delivered := getPresents(house)

		if delivered >= presents {
			return house
		}
	}
}

func main() {
	fmt.Println(getHouseWith(33100000))
}
