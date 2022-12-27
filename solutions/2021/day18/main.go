package main

import (
	"encoding/json"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/probability"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 18
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numbers := parseInput(input)
	pair := reduce(numbers)
	return magnitude(pair)
}

func part2(input []string) interface{} {
	max := math.MinInt
	probability.PermuteSpots(input, 2, 2, func(sample []string) {
		numbers := parseInput(sample)
		pair := reduce(numbers)
		max = ints.Max(max, magnitude(pair))
	})
	return max
}

func reduce(numbers []Pair) Pair {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = result.add(number)
	}
	return result
}

func parseInput(input []string) []Pair {
	numbers := make([]Pair, len(input))
	for i, line := range input {
		var pair []interface{}
		_ = json.Unmarshal([]byte(line), &pair)
		numbers[i] = parsePair(pair)
	}
	return numbers
}

func parsePair(pair []interface{}) Pair {
	return Pair{
		left:  parseValue(pair[0]),
		right: parseValue(pair[1]),
	}
}

func parseValue(v interface{}) interface{} {
	switch t := v.(type) {
	case float64:
		return int(t)
	case []interface{}:
		return parsePair(t)
	}
	return nil
}

type Pair struct {
	left  interface{}
	right interface{}
}

func magnitude(value interface{}) int {
	if v, ok := value.(int); ok {
		return v
	} else {
		p := value.(Pair)
		return 3*magnitude(p.left) + 2*magnitude(p.right)
	}
}

func (p Pair) add(p2 Pair) Pair {
	var result interface{}
	result = Pair{
		left:  p,
		right: p2,
	}

	var changed bool
	for {
		changed, _, result, _ = explode(result, 0)
		if changed {
			continue
		}

		changed, result = split(result)
		if !changed {
			break
		}
	}

	return result.(Pair)
}

func split(value interface{}) (bool, interface{}) {
	switch v := value.(type) {
	case int:
		if v >= 10 {
			return true, Pair{left: v / 2, right: int(math.Ceil(float64(v) / 2))}
		}
		return false, v
	case Pair:
		changed, nextLeft := split(v.left)
		if changed {
			return true, Pair{left: nextLeft, right: v.right}
		}

		changed, nextRight := split(v.right)
		return changed, Pair{left: nextLeft, right: nextRight}
	}

	return false, nil
}

func explode(value interface{}, depth int) (bool, interface{}, interface{}, interface{}) {
	switch v := value.(type) {
	case int:
		return false, nil, v, nil
	case Pair:
		if depth == 4 {
			return true, v.left, 0, v.right
		}

		changed, left, mid, right := explode(v.left, depth+1)
		if changed {
			return true, left, Pair{left: mid, right: addLeft(v.right, right)}, nil
		}

		changed, left, mid, right = explode(v.right, depth+1)
		if changed {
			return true, nil, Pair{left: addRight(v.left, left), right: mid}, right
		}
	}
	return false, nil, value, nil
}

func addLeft(value interface{}, right interface{}) interface{} {
	if right == nil {
		return value
	}

	switch v := value.(type) {
	case int:
		return v + right.(int)
	case Pair:
		return Pair{left: addLeft(v.left, right), right: v.right}
	}

	return nil
}

func addRight(value interface{}, left interface{}) interface{} {
	if left == nil {
		return value
	}

	switch v := value.(type) {
	case int:
		return v + left.(int)
	case Pair:
		return Pair{left: v.left, right: addRight(v.right, left)}
	}

	return nil
}
