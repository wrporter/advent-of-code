package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	numContainingPairs := 0

	for _, line := range strings.Split(input, "\n") {
		min1, max1, min2, max2 := extractPairs(line)

		if (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1) {
			numContainingPairs += 1
		}
	}

	return numContainingPairs
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	numOverlappingPairs := 0

	for _, line := range strings.Split(input, "\n") {
		min1, max1, min2, max2 := extractPairs(line)

		if (min1 >= min2 && min1 <= max2) || (max1 >= min2 && max1 <= max2) ||
			(min2 >= min1 && min2 <= max1) || (max2 >= min1 && max2 <= max1) {
			numOverlappingPairs += 1
		}
	}

	return numOverlappingPairs
}

func extractPairs(line string) (int, int, int, int) {
	pairs := strings.Split(line, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")
	min1, max1 := convert.StringToInt(pair1[0]), convert.StringToInt(pair1[1])
	min2, max2 := convert.StringToInt(pair2[0]), convert.StringToInt(pair2[1])
	return min1, max1, min2, max2
}
