package solution

import (
	"aoc/src/lib/go/convert"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	records := parse(input)
	return sumArrangements(records)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	records := parse(input)

	for i, record := range records {
		var sb strings.Builder
		sb.WriteString(record.springs)
		delimiter := byte('?')

		for copies := 1; copies < 5; copies++ {
			sb.WriteByte(delimiter)
			sb.WriteString(record.springs)
			records[i].sizes = append(records[i].sizes, record.sizes...)
		}

		records[i].springs = sb.String()
	}

	return sumArrangements(records)
}

func sumArrangements(records []Record) int {
	sum := 0
	for _, record := range records {
		sum += countArrangements(record)
	}
	return sum
}

func countArrangements(record Record) int {
	springs := "." + record.springs
	sizes := record.sizes

	cache := make([]int, len(springs)+1)
	cache[0] = 1

	for i, spring := range springs {
		if spring != '#' {
			cache[i+1] = 1
		} else {
			break
		}
	}

	requirementStart := 1

	for _, size := range sizes {
		nextCache := make([]int, len(springs)+1)
		group := 0
		sizeRequirementMet := false

		for i := requirementStart; i < len(springs); i++ {
			spring := springs[i]

			if spring != '.' {
				group += 1
			} else {
				group = 0
			}

			if spring != '#' {
				nextCache[i+1] += nextCache[i]
			}

			if group >= size && springs[i-size] != '#' {
				nextCache[i+1] += cache[i-size]
			}

			if nextCache[i+1] != 0 && !sizeRequirementMet {
				requirementStart = i + 2
				sizeRequirementMet = true
			}
		}

		cache = nextCache
	}

	return cache[len(cache)-1]
}

type CacheKey struct {
	springIndex int
	sizeIndex   int
	groupSize   int
}

type Record struct {
	springs string
	sizes   []int
}

func parse(input string) []Record {
	lines := strings.Split(input, "\n")
	records := make([]Record, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		springs := parts[0]
		sizes, _ := convert.ToInts(strings.Split(parts[1], ","))
		records[i] = Record{springs, sizes}
	}
	return records
}
