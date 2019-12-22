package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"github.com/wrporter/advent-of-code-2019/internal/common/ints"
	"regexp"
)

const SpaceDeckSize = 10007

var (
	stackRegex     = regexp.MustCompile(`^deal into new stack$`)
	cutRegex       = regexp.MustCompile(`^cut (-?\d+)$`)
	incrementRegex = regexp.MustCompile(`^deal with increment (\d+)$`)
)

type Deck struct {
	Cards []int
}

func (d *Deck) Shuffle(techniques []string) {
	for _, technique := range techniques {
		if matched := stackRegex.MatchString(technique); matched {
			d.ShuffleStack()
		} else if matched := cutRegex.MatchString(technique); matched {
			match := cutRegex.FindStringSubmatch(technique)
			amount := conversion.StringToInt(match[1])
			d.ShuffleCut(amount)
		} else if matched := incrementRegex.MatchString(technique); matched {
			match := incrementRegex.FindStringSubmatch(technique)
			increment := conversion.StringToInt(match[1])
			d.ShuffleIncrement(increment)
		}
	}
}

func (d *Deck) ShuffleStack() {
	var cards []int
	for _, card := range d.Cards {
		cards = ints.Prepend(cards, card)
	}
	d.Cards = cards
}

func (d *Deck) ShuffleCut(amount int) {
	cards := make([]int, len(d.Cards))
	to := len(cards) - amount
	from := amount
	if amount < 0 {
		to = ints.Abs(amount)
		from = ints.WrapMod(amount, len(cards))
	}
	copy(cards[to:], d.Cards[:from])
	copy(cards[:to], d.Cards[from:])
	d.Cards = cards
}

func (d *Deck) ShuffleIncrement(increment int) {
	cards := make([]int, len(d.Cards))
	var card int
	for i := 0; len(d.Cards) > 0; i = (i + increment) % len(cards) {
		card, d.Cards = ints.Poll(d.Cards)
		cards[i] = card
	}
	d.Cards = cards
}

func (d *Deck) Find(card int) int {
	for index, c := range d.Cards {
		if card == c {
			return index
		}
	}
	return -1
}

func New(size int) *Deck {
	cards := make([]int, size)
	for i := 0; i < size; i++ {
		cards[i] = i
	}
	return &Deck{cards}
}

func main() {
	techniques, _ := file.ReadFile("./day22/input.txt")
	deck := New(SpaceDeckSize)
	deck.Shuffle(techniques)
	fmt.Println(deck.Find(2019))
}
