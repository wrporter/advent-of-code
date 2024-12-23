package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"encoding/json"
	"fmt"
	"regexp"
)

var regex = regexp.MustCompile(`-?\d+`)

func sum(json string) int {
	values := regex.FindAllString(json, -1)
	numbers, _ := convert.ToInts(values)
	return ints.Sum(numbers)
}

func SumRed(data string) int {
	var container interface{}
	_ = json.Unmarshal([]byte(data), &container)
	return sumRed(container)
}

func sumRed(object interface{}) int {
	sum := 0

	switch cur := object.(type) {
	case float64:
		sum += int(cur)
	case []interface{}:
		for i := 0; i < len(cur); i++ {
			sum += sumRed(cur[i])
		}
	case map[string]interface{}:
		potentialSum := 0
		for _, value := range cur {
			if value == "red" {
				potentialSum = 0
				break
			} else {
				potentialSum += sumRed(value)
			}
		}
		sum += potentialSum
	}

	return sum
}

func main() {
	lines, _ := file.ReadFile("./2015/day12/input.txt")
	fmt.Println(sum(lines[0]))
	fmt.Println(SumRed(lines[0]))
}
