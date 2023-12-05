package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"math"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	seeds, _, maps := parseInput(input)

	for _, map0 := range maps {
		for i, seed := range seeds {
			for _, c := range map0.convert {
				if seed >= c.source.start && seed <= c.source.end {
					delta := seed - c.source.start
					seeds[i] = c.destination.start + delta
				}
			}
		}
	}

	low := math.MaxInt
	for _, seed := range seeds {
		low = min(low, seed)
	}

	return low
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	_, seeds, maps := parseInput(input)

	for _, map0 := range maps {
		for i, seed := range seeds {
			for _, c := range map0.convert {
				if seed >= c.source.start && seed <= c.source.end {
					delta := seed - c.source.start
					seeds[i] = c.destination.start + delta
				}
			}
		}
	}

	low := math.MaxInt
	for _, seed := range seeds {
		low = min(low, seed)
	}

	return low
}

type aRange struct {
	start int
	end   int
}

type conversion struct {
	source      aRange
	destination aRange
}

type mapping struct {
	from    string
	to      string
	convert []conversion
}

func parseInput(input string) ([]int, []int, []*mapping) {
	chunks := strings.Split(input, "\n\n")
	seeds, _ := convert.ToInts(strings.Fields(strings.Split(chunks[0], "seeds: ")[1]))
	mapsStrings := chunks[1:]
	maps := make([]*mapping, len(mapsStrings))

	for i, mapString := range mapsStrings {
		parts := strings.Split(mapString, "\n")
		categoryParts := strings.Split(strings.Fields(parts[0])[0], "-")
		conversions := make([]conversion, len(parts)-1)

		for j, line := range parts[1:] {
			numbers, _ := convert.ToInts(strings.Fields(line))

			conversions[j] = conversion{
				source: aRange{
					start: numbers[1],
					end:   numbers[1] + numbers[2] - 1,
				},
				destination: aRange{
					start: numbers[0],
					end:   numbers[0] + numbers[2] - 1,
				},
			}
		}

		maps[i] = &mapping{
			convert: conversions,
			from:    categoryParts[0],
			to:      categoryParts[2],
		}
	}

	var seeds2 []int
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		for seed := start; seed < start+length-1; seed++ {
			seeds2 = append(seeds2, seed)
		}
	}

	return seeds, seeds2, maps
}
