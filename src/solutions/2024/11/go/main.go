package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"math"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	return countStones(input, 25)
}

func part2(input string, _ ...interface{}) interface{} {
	return countStones(input, 75)
}

func countStones(input string, numBlinks int) int {
	count := 0
	for _, number := range convert.ToIntsV2(strings.Fields(input)) {
		count += countStonesRec(make(map[cacheKey]int), number, numBlinks)
	}
	return count
}

func countStonesRec(cache map[cacheKey]int, number, blinks int) int {
	key := cacheKey{number, blinks}
	if blinks == 0 {
		cache[key] = 1
		return 1
	}

	if count, ok := cache[key]; ok {
		return count
	}

	result := 0
	if number == 0 {
		result += countStonesRec(cache, 1, blinks-1)
	} else if digits := countDigits(number); digits%2 == 0 {
		firstHalf := number / int(math.Pow10(digits/2))
		result += countStonesRec(cache, firstHalf, blinks-1)

		secondHalf := number % int(math.Pow10(digits/2))
		result += countStonesRec(cache, secondHalf, blinks-1)
	} else {
		result += countStonesRec(cache, number*2024, blinks-1)
	}

	cache[key] = result
	return result
}

func countDigits(number int) int {
	return int(math.Floor(math.Log10(float64(number)))) + 1
}

type cacheKey struct {
	number int
	blinks int
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 11, Part1: part1, Part2: part2}
}
