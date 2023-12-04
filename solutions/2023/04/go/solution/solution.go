package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"strconv"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	cards := parseInput(input)
	sum := 0

	for _, card := range cards {
		matches := 0
		for num := range card.have {
			if card.winning[num] {
				matches += 1
			}
		}

		sum += ints.Pow(2, matches-1)
	}

	return sum
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	cards := parseInput(input)
	numCards := 0

	for id, card := range cards {
		matches := 0
		for num := range card.have {
			if card.winning[num] {
				matches += 1
				if id+matches < len(cards) {
					cards[id+matches].copies += card.copies
				}
			}
		}

		numCards += card.copies
	}

	return numCards
}

type cardType struct {
	id      int
	winning map[int]bool
	have    map[int]bool
	copies  int
}

func parseInput(input string) []*cardType {
	cardsStr := strings.Split(input, "\n")
	cards := make([]*cardType, len(cardsStr))

	// The actual ID is not important, just the order. So it's okay if we start
	// at 0.
	for id, card := range cardsStr {
		parts := strings.Split(card, ":")
		numbers := strings.Split(parts[1], "|")
		winning := make(map[int]bool)
		have := make(map[int]bool)

		for _, num := range strings.Fields(numbers[0]) {
			value, _ := strconv.Atoi(num)
			winning[value] = true
		}
		for _, num := range strings.Fields(numbers[1]) {
			value, _ := strconv.Atoi(num)
			have[value] = true
		}

		cards[id] = &cardType{
			id:      id,
			winning: winning,
			have:    have,
			copies:  1,
		}
	}

	return cards
}
