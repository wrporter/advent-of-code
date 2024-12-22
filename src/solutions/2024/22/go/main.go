package main

import (
	"aoc/src/lib/go/aoc"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	sum := 0
	for _, numStr := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(numStr)
		for i := 1; i <= 2000; i++ {
			num = next(num)
		}
		sum += num
	}
	return sum
}

func part2(input string, _ ...interface{}) interface{} {
	sequences := make(map[[4]int]int)

	for _, numStr := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(numStr)

		nums := make([]int, 2001)
		nums[0] = num

		changes := make([]int, len(nums)-1)
		seen := make(map[[4]int]bool)

		for i := 1; i <= 2000; i++ {
			nums[i] = next(nums[i-1])
			changes[i-1] = nums[i]%10 - nums[i-1]%10

			if i >= 4 && i < 2000-3 {
				sequence := [4]int{changes[i-4], changes[i-3], changes[i-2], changes[i-1]}
				if !seen[sequence] {
					sequences[sequence] += nums[i] % 10
					seen[sequence] = true
				}
			}
		}
	}

	best := 0
	for _, value := range sequences {
		best = max(best, value)
	}
	return best
}

func next(num int) int {
	num ^= num * 64 % 16777216
	num ^= num / 32 % 16777216
	num ^= num * 2048 % 16777216
	return num
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 22, Part1: part1, Part2: part2}
}
