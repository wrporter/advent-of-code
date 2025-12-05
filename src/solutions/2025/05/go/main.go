package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	parts := strings.Split(input, "\n\n")
	ids, _ := convert.ToInts(strings.Split(parts[1], "\n"))
	numFresh := 0

	for _, id := range ids {
		for _, idRange := range strings.Split(parts[0], "\n") {
			idRangeParts := strings.Split(idRange, "-")
			start := convert.StringToInt(idRangeParts[0])
			end := convert.StringToInt(idRangeParts[1])

			if id >= start && id <= end {
				numFresh++
				break
			}
		}
	}

	return numFresh
}

type idRange struct {
	start int
	end   int
}

// Sort ranges by the start and trim ranges afterward
func part2(input string, _ ...interface{}) interface{} {
	parts := strings.Split(input, "\n\n")
	numFresh := 0
	idRangesStr := strings.Split(parts[0], "\n")
	ranges := make([]idRange, len(idRangesStr))

	for i, idRangeStr := range strings.Split(parts[0], "\n") {
		idRangeParts := strings.Split(idRangeStr, "-")
		start := convert.StringToInt(idRangeParts[0])
		end := convert.StringToInt(idRangeParts[1])
		ranges[i] = idRange{
			start: start,
			end:   end,
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	for i, r := range ranges {
		start := r.start
		end := r.end
		overlaps := false

		for j := 0; j < i; j++ {
			prev := ranges[j]

			if start >= prev.start && end <= prev.end {
				overlaps = true
				break
			}

			if start >= prev.start && start <= prev.end && end > prev.end {
				start = prev.end + 1
			}
		}

		if !overlaps {
			numFresh += end - start + 1
		}

	}

	return numFresh
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 5, Part1: part1, Part2: part2}
}
