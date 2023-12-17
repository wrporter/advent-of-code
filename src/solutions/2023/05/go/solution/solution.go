package solution

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/mymath"
	"math"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	seeds, maps := parseInput(input)

	for _, conversions := range maps {
		for i, seed := range seeds {
			for _, c := range conversions {
				if seed >= c.source && seed < c.sourceEnd {
					seeds[i] = seed + c.destination - c.source
				}
			}
		}
	}

	return mymath.Min(seeds...)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	seeds, maps := parseInput(input)
	low := math.MaxInt

	for i := 0; i < len(seeds); i += 2 {
		current := []interval{{seeds[i], seeds[i] + seeds[i+1]}}

		for _, conversions := range maps {
			// Seed ranges that are fully contained within a category conversion
			var contained []interval

			for _, c := range conversions {
				// Ranges that are split that still need to be considered for conversion by other
				// ranges
				var split []interval

				for _, node := range current {
					// Ranges relative to the category conversion
					before := interval{node.start, min(node.end, c.source)}
					within := interval{max(node.start, c.source), min(c.sourceEnd, node.end)}
					after := interval{max(c.sourceEnd, node.start), node.end}

					// Only process valid ranges
					if before.end > before.start {
						split = append(split, before)
					}

					if within.end > within.start {
						contained = append(contained, interval{
							start: within.start - c.source + c.destination,
							end:   within.end - c.source + c.destination,
						})
					}

					if after.end > after.start {
						split = append(split, after)
					}
				}

				current = split
			}

			current = append(current, contained...)
		}

		for _, r := range current {
			low = min(low, r.start)
		}
	}

	return low
}

// range is a keyword, so we use interval instead
type interval struct {
	start int
	end   int
}

type conversion struct {
	destination int
	source      int
	sourceEnd   int
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
				sourceEnd:   numbers[1] + numbers[2],
			}
		}

		maps[i] = conversions
	}

	return seeds, maps
}
