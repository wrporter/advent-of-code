package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"math"
	"strconv"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last byte

		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				first = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				last = line[i]
				break
			}
		}

		value, _ := strconv.Atoi(string(first) + string(last))
		sum += value
	}

	return sum
}

var numbers = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last string
		locations := make(map[int]string)

		for num := range numbers {
			if i := strings.Index(line, num); i >= 0 {
				locations[i] = num
			}
			if i := strings.LastIndex(line, num); i >= 0 {
				locations[i] = num
			}
		}

		low := math.MaxInt
		high := math.MinInt
		for i := range locations {
			low = ints.Min(low, i)
			high = ints.Max(high, i)
		}

		first = numbers[locations[low]]
		last = numbers[locations[high]]

		value, _ := strconv.Atoi(first + last)
		sum += value
	}

	return sum
}
