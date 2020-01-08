package main

import (
	"encoding/json"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"regexp"
)

var regex = regexp.MustCompile(`-?\d+`)

func sum(json string) int {
	values := regex.FindAllString(json, -1)
	numbers, _ := conversion.ToInts(values)
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
