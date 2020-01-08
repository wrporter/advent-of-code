package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/probability"
	"regexp"
)

var regex = regexp.MustCompile(`([a-zA-Z]+) would (lose|gain) (\d+) happiness units by sitting next to ([a-zA-Z]+).`)

func getOptimalRating(ratingStrings []string) (int, []string) {
	ratings, people := parse(ratingStrings)
	var maxHappiness int
	var bestSeatingArrangement []string

	probability.PermuteStrings(people, func(table []string) {
		happiness := 0
		for i, person := range table {
			if i == 0 {
				happiness += ratings[person][table[len(table)-1]]
			} else {
				happiness += ratings[person][table[i-1]]
			}
			if i == len(table)-1 {
				happiness += ratings[person][table[0]]
			} else {
				happiness += ratings[person][table[i+1]]
			}
		}
		if happiness > maxHappiness {
			maxHappiness = happiness
			bestSeatingArrangement = table
		}
	})

	return maxHappiness, bestSeatingArrangement
}

func parse(ratingStrings []string) (map[string]map[string]int, []string) {
	ratings := make(map[string]map[string]int)
	var people []string

	for _, rating := range ratingStrings {
		match := regex.FindStringSubmatch(rating)
		value := conversion.StringToInt(match[3])
		if match[2] == "lose" {
			value = -value
		}
		if _, ok := ratings[match[1]]; ok {
			ratings[match[1]][match[4]] = value
		} else {
			people = append(people, match[1])
			ratings[match[1]] = map[string]int{match[4]: value}
		}
	}
	return ratings, people
}

func main() {
	lines, _ := file.ReadFile("./2015/day13/input.txt")
	//lines := []string{
	//	"Alice would gain 54 happiness units by sitting next to Bob.",
	//	"Alice would lose 79 happiness units by sitting next to Carol.",
	//	"Alice would lose 2 happiness units by sitting next to David.",
	//	"Bob would gain 83 happiness units by sitting next to Alice.",
	//	"Bob would lose 7 happiness units by sitting next to Carol.",
	//	"Bob would lose 63 happiness units by sitting next to David.",
	//	"Carol would lose 62 happiness units by sitting next to Alice.",
	//	"Carol would gain 60 happiness units by sitting next to Bob.",
	//	"Carol would gain 55 happiness units by sitting next to David.",
	//	"David would gain 46 happiness units by sitting next to Alice.",
	//	"David would lose 7 happiness units by sitting next to Bob.",
	//	"David would gain 41 happiness units by sitting next to Carol.",
	//}
	fmt.Println(getOptimalRating(lines))
}
