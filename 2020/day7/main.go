package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 7
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	bags := parseBags(input)
	numBagsContainingShinyGold := 0

	queue := []string{"shiny gold"}
	visited := make(map[string]bool)
	var desiredColor string

	for len(queue) > 0 {
		desiredColor, queue = queue[len(queue)-1], queue[:len(queue)-1]

		for nextColor := range bags {
			if _, ok := bags[nextColor][desiredColor]; ok && !visited[nextColor] {
				numBagsContainingShinyGold++
				visited[nextColor] = true
				queue = append(queue, nextColor)
			}
		}
	}

	return numBagsContainingShinyGold
}

func part2(input []string) interface{} {
	bags := parseBags(input)
	var recurse func(desiredColor string) int

	recurse = func(desiredColor string) int {
		if bags[desiredColor] == nil {
			return 1
		}

		numBags := 1
		for nextColor, amount := range bags[desiredColor] {
			numBags += amount * recurse(nextColor)
		}
		return numBags
	}

	return recurse("shiny gold") - 1
}

var regex = regexp.MustCompile(`^(.+) bags contain (no other bags|.* bags?,? ?)*\.$`)
var contentsRegex = regexp.MustCompile(`^(\d+) (.*) bags?$`)

func parseBags(input []string) map[string]map[string]int {
	bags := make(map[string]map[string]int)

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		var contents map[string]int

		if match[2] != "no other bags" {
			contentString := strings.Split(match[2], ", ")
			contents = make(map[string]int)
			for _, content := range contentString {
				contentMatch := contentsRegex.FindStringSubmatch(content)
				contents[contentMatch[2]] = convert.StringToInt(contentMatch[1])
			}
		}

		bags[match[1]] = contents
	}

	return bags
}
