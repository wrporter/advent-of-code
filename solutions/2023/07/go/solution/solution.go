package solution

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"sort"
	"strconv"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	hands := parseInput(input)
	sortHands(hands)
	return sumWinnings(hands)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	hands := parseInput(strings.ReplaceAll(input, "J", "W"))
	sortHands(hands)
	return sumWinnings(hands)
}

var strength = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
	'W': -1, // wild joker
}

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type camelHand struct {
	cards string
	bid   int
	type_ int
}

func parseInput(input string) []camelHand {
	lines := strings.Split(input, "\n")
	hands := make([]camelHand, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		cards := parts[0]
		bid := convert.StringToInt(parts[1])

		counts := make(map[rune]int)
		most := 0
		wilds := 0

		for _, card := range cards {
			if card != 'W' {
				if _, ok := counts[card]; ok {
					counts[card]++
				} else {
					counts[card] = 1
				}
				most = max(most, counts[card])
			} else {
				wilds++
			}
		}

		most = most + wilds
		type_ := -1

		if most == 5 {
			type_ = fiveOfAKind
		} else if most == 4 {
			type_ = fourOfAKind
		} else if len(counts) == 2 && most == 3 {
			type_ = fullHouse
		} else if most == 3 {
			type_ = threeOfAKind
		} else if len(counts) == 3 && most == 2 {
			type_ = twoPair
		} else if most == 2 {
			type_ = onePair
		} else {
			type_ = highCard
		}

		hands[i] = camelHand{cards, bid, type_}
	}
	return hands
}

func sortHands(hands []camelHand) {
	sort.Slice(hands, func(i, j int) bool {
		hand1, hand2 := hands[i], hands[j]

		if hand1.type_ == hand2.type_ {
			for c := range hand1.cards {
				strength1 := strength[hand1.cards[c]]
				strength2 := strength[hand2.cards[c]]

				if strength1 < strength2 {
					return true
				} else if strength1 > strength2 {
					return false
				}
			}
			return false
		}

		return hand1.type_ < hand2.type_
	})
}

func sumWinnings(hands []camelHand) int {
	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		totalWinnings += hand.bid * rank
	}
	return totalWinnings
}

func printDebug(hands []camelHand) {
	out := &strings.Builder{}
	for _, hand := range hands {
		out.WriteString(hand.cards)
		out.WriteString(" ")
		out.WriteString(strconv.Itoa(hand.type_))
		out.WriteString("\n")
	}
	fmt.Print(out.String())
}
