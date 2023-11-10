package solution

import (
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	rucksacks := parseInput(input)
	priority := 0

	for _, rucksack := range rucksacks {
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]
		found := false

		for i := 0; i < len(compartment1) && !found; i++ {
			for j := 0; j < len(compartment2) && !found; j++ {
				if compartment1[i] == compartment2[j] {
					priority += getPriority(rune(compartment1[i]))
					found = true
				}
			}
		}
	}

	return priority
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	rucksacks := parseInput(input)
	priority := 0

	for i := 0; i < len(rucksacks); i += 3 {
		rucksack1 := rucksacks[i]
		rucksack2 := rucksacks[i+1]
		rucksack3 := rucksacks[i+2]
		found := false

		for j := 0; j < len(rucksack1) && !found; j++ {
			for k := 0; k < len(rucksack2) && !found; k++ {
				for l := 0; l < len(rucksack3) && !found; l++ {
					if rucksack1[j] == rucksack2[k] && rucksack1[j] == rucksack3[l] {
						priority += getPriority(rune(rucksack1[j]))
						found = true
					}
				}
			}
		}
	}

	return priority
}

func getPriority(item rune) int {
	if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	}
	return int(item-'a') + 1
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
