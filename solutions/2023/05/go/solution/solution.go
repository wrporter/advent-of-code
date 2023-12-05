package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/v2/mymath"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	seeds, maps := parseInput(input)

	for _, conversions := range maps {
		for i, seed := range seeds {
			for _, c := range conversions {
				if seed >= c.source && seed < c.source+c.length {
					seeds[i] = seed + c.destination - c.source
				}
			}
		}
	}

	return mymath.Min(seeds...)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	seeds, maps := parseInput(input)
	var seeds2 []int
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		for seed := start; seed < start+length-1; seed++ {
			seeds2 = append(seeds2, seed)
		}
	}

	for _, conversions := range maps {
		for i, seed := range seeds2 {
			for _, c := range conversions {
				if seed >= c.source && seed < c.source+c.length {
					delta := seed - c.source
					seeds2[i] = c.destination + delta
				}
			}
		}
	}

	return mymath.Min(seeds2...)
}

type conversion struct {
	destination int
	source      int
	length      int
}

func parseInput(input string) ([]int, [][]conversion) {
	chunks := strings.Split(input, "\n\n")
	seeds, _ := convert.ToInts(strings.Fields(chunks[0])[1:])
	maps := make([][]conversion, len(chunks)-1)

	for i, chunk := range chunks[1:] {
		parts := strings.Split(chunk, "\n")
		conversions := make([]conversion, len(parts)-1)

		for j, line := range parts[1:] {
			numbers, _ := convert.ToInts(strings.Fields(line))

			conversions[j] = conversion{
				destination: numbers[0],
				source:      numbers[1],
				length:      numbers[2],
			}
		}

		maps[i] = conversions
	}

	return seeds, maps
}
