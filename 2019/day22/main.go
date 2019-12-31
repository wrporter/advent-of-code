package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"math/big"
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
	Size  int
}

func (deck *Deck) Shuffle(techniques []string) {
	for _, technique := range techniques {
		if matched := stackRegex.MatchString(technique); matched {
			deck.ShuffleStack()
		} else if matched := cutRegex.MatchString(technique); matched {
			match := cutRegex.FindStringSubmatch(technique)
			amount := conversion.StringToInt(match[1])
			deck.ShuffleCut(amount)
		} else if matched := incrementRegex.MatchString(technique); matched {
			match := incrementRegex.FindStringSubmatch(technique)
			increment := conversion.StringToInt(match[1])
			deck.ShuffleIncrement(increment)
		}
	}
}

func GetCard(techniques []string, size *big.Int, times *big.Int, position *big.Int) *big.Int {
	incrementMul := big.NewInt(1)
	offsetDiff := big.NewInt(0)

	for _, technique := range techniques {
		if matched := stackRegex.MatchString(technique); matched {
			incrementMul.Mul(incrementMul, big.NewInt(-1))
			incrementMul = WrapMod(incrementMul, size)
			offsetDiff.Add(offsetDiff, incrementMul)
			offsetDiff = WrapMod(offsetDiff, size)
		} else if matched := cutRegex.MatchString(technique); matched {
			match := cutRegex.FindStringSubmatch(technique)
			amount, _ := new(big.Int).SetString(match[1], 10)
			offsetDiff.Add(offsetDiff, amount.Mul(amount, incrementMul))
			offsetDiff = WrapMod(offsetDiff, size)
		} else if matched := incrementRegex.MatchString(technique); matched {
			match := incrementRegex.FindStringSubmatch(technique)
			amount, _ := new(big.Int).SetString(match[1], 10)
			incrementMul.Mul(incrementMul, inv(amount, size))
			incrementMul = WrapMod(incrementMul, size)
		}
	}

	increment := new(big.Int).Exp(incrementMul, times, size)

	offset := new(big.Int).Mul(offsetDiff, new(big.Int).Sub(big.NewInt(1), increment))
	inverse := new(big.Int).Sub(big.NewInt(1), incrementMul)
	offset.Mul(offset, inv(inverse.Mod(inverse, size), size))
	offset.Mod(offset, size)

	card := new(big.Int).Mul(position, increment)
	card.Add(card, offset)
	card.Mod(card, size)
	return card
}

func inv(x *big.Int, y *big.Int) *big.Int {
	return x.Exp(x, new(big.Int).Sub(y, big.NewInt(2)), y)
}

func WrapMod(d, m *big.Int) *big.Int {
	var res = new(big.Int).Mod(d, m)
	if (res.Cmp(big.NewInt(0)) == -1 && m.Cmp(big.NewInt(0)) == 1) ||
		(res.Cmp(big.NewInt(0)) == 1 && m.Cmp(big.NewInt(0)) == -1) {
		return res.Add(res, m)
	}
	return res
}

func (deck *Deck) ShuffleStack() {
	var cards []int
	for _, card := range deck.Cards {
		cards = ints.Prepend(cards, card)
	}
	deck.Cards = cards
}

func (deck *Deck) ShuffleCut(amount int) {
	cards := make([]int, len(deck.Cards))
	to := len(cards) - amount
	from := amount
	if amount < 0 {
		to = ints.Abs(amount)
		from = ints.WrapMod(amount, len(cards))
	}
	copy(cards[to:], deck.Cards[:from])
	copy(cards[:to], deck.Cards[from:])
	deck.Cards = cards
}

func (deck *Deck) ShuffleIncrement(increment int) {
	cards := make([]int, len(deck.Cards))
	var card int
	for i := 0; len(deck.Cards) > 0; i = (i + increment) % len(cards) {
		card, deck.Cards = ints.Poll(deck.Cards)
		cards[i] = card
	}
	deck.Cards = cards
}

func (deck *Deck) Find(card int) int {
	for index, c := range deck.Cards {
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
	return &Deck{cards, size}
}

func main() {
	techniques, _ := file.ReadFile("./2019/day22/input.txt")
	deck := New(SpaceDeckSize)
	deck.Shuffle(techniques)
	fmt.Println(deck.Find(2019))
	// Answer: 7545

	size, _ := new(big.Int).SetString("119315717514047", 10)
	times, _ := new(big.Int).SetString("101741582076661", 10)
	position, _ := new(big.Int).SetString("2020", 10)
	fmt.Println(GetCard(techniques, size, times, position))
	// Answer: 12706692375144
}
