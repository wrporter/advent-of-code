package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	hands := parseInput(input)
	sortHands(hands)
	return sumWinnings(hands)
}

func part2(input string, _ ...interface{}) interface{} {
	hands := parseInput(strings.ReplaceAll(input, "J", "W"))
	sortHands(hands)
	return sumWinnings(hands)
}

var strength = make(map[byte]int)

func init() {
	for i, card := range []byte("W23456789TJQKA") {
		strength[card] = i
	}
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

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 7, Part1: part1, Part2: part2}
}
