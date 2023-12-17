package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 20
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	blocklist := parseInput(input)
	ip := int64(-1)
	max := int64(-1)

	for _, block := range blocklist {
		numValid := getMax(0, block.start-max-1)
		if numValid > 0 {
			ip = max + 1
			break
		}
		max = getMax(max, block.end)
	}

	return ip
}

func part2(input []string) interface{} {
	blocklist := parseInput(input)
	var count int64
	max := int64(-1)

	for _, block := range blocklist {
		count += getMax(0, block.start-max-1)
		max = getMax(max, block.end)
	}

	return count
}

func parseInput(input []string) []ipRange {
	ranges := make([]ipRange, len(input))
	for i, line := range input {
		parts := strings.Split(line, "-")
		start := int64(convert.StringToInt(parts[0]))
		end := int64(convert.StringToInt(parts[1]))
		ranges[i] = ipRange{start, end}
	}
	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	return ranges
}

func getMax(values ...int64) int64 {
	var result int64
	for _, value := range values {
		if value > result {
			result = value
		}
	}
	return result
}

var maxInt64 = int64(^uint32(0))

type ipRange struct {
	start int64
	end   int64
}
