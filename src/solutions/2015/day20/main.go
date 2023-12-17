package main

import (
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/timeit"
	"fmt"
	"time"
)

func part1(minPresents int) int {
	defer timeit.Track(time.Now(), "part1")
	for house := 1; ; house++ {
		presents := 0
		divisors := ints.GetDivisors(house)

		for _, elf := range divisors {
			presents += elf
		}
		presents *= 10

		if presents >= minPresents {
			return house
		}
	}
}

func part2(minPresents int) int {
	defer timeit.Track(time.Now(), "part2")
	for house := 1; ; house++ {
		presents := 0
		divisors := ints.GetDivisors(house)
		divisors = ints.TakeLast(divisors, 50)

		for _, elf := range divisors {
			if elf*50 >= house {
				presents += elf
			}
		}
		presents *= 11

		if presents >= minPresents {
			return house
		}
	}
}

func main() {
	fmt.Println(part1(33100000))
	fmt.Println(part2(33100000))
}
