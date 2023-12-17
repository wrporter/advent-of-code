package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"fmt"
	"regexp"
	"strings"
)

var sueRegex = regexp.MustCompile(`(Sue \d+): (.*)`)
var itemRegex = regexp.MustCompile(`([a-z]+): (\d+)`)

func matchAuntSue(auntSueStrings []string, theOne map[string]int, shouldCompare bool) string {
	for _, auntSueString := range auntSueStrings {
		match := sueRegex.FindStringSubmatch(auntSueString)
		sue := make(map[string]int)

		itemStrings := strings.Split(match[2], ", ")
		for _, itemString := range itemStrings {
			itemMatch := itemRegex.FindStringSubmatch(itemString)
			sue[itemMatch[1]] = convert.StringToInt(itemMatch[2])
		}

		found := true
		for item, quantity := range sue {
			if shouldCompare {
				if !((item == "cats" || item == "trees") && quantity >= theOne[item]) &&
					!((item == "pomeranians" || item == "goldfish") && quantity <= theOne[item]) &&
					theOne[item] != quantity {
					found = false
				}
			} else if theOne[item] != quantity {
				found = false
			}
		}
		if found {
			return match[1]
		}
	}

	return "No sue :("
}

func main() {
	input, _ := file.ReadFile("./2015/day16/input.txt")
	theOne := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	fmt.Println(matchAuntSue(input, theOne, false))
	fmt.Println(matchAuntSue(input, theOne, true))
}
