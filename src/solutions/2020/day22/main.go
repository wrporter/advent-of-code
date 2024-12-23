package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	player1, player2 := parse(input)
	winner := play(player1, player2)
	return score(winner)
}

func part2(input []string) interface{} {
	player1, player2 := parse(input)
	_, deck := playRecursive(player1, player2)
	return score(deck)
}

func playRecursive(deck1 []int, deck2 []int) (int, []int) {
	seen := make(map[int]bool)
	var card1, card2, winner int

	for round := 1; ; round++ {
		if len(deck1) == 0 {
			return 2, deck2
		}
		if len(deck2) == 0 {
			return 1, deck1
		}

		key1 := ints.HashCode(deck1)
		key2 := ints.HashCode(deck2)
		if seen[key1] && seen[key2] {
			return 1, deck1
		}

		seen[key1] = true
		seen[key2] = true
		card1, deck1 = deck1[0], deck1[1:]
		card2, deck2 = deck2[0], deck2[1:]

		if card1 <= len(deck1) && card2 <= len(deck2) {
			player1Copy := copySize(deck1, card1)
			player2Copy := copySize(deck2, card2)
			winner, _ = playRecursive(player1Copy, player2Copy)
		} else if card1 > card2 {
			winner = 1
		} else {
			winner = 2
		}

		if winner == 1 {
			deck1 = append(deck1, card1)
			deck1 = append(deck1, card2)
		} else {
			deck2 = append(deck2, card2)
			deck2 = append(deck2, card1)
		}
	}
}

func copySize(array []int, size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = array[i]
	}
	return result
}

func score(deck []int) int {
	result := 0

	multiplier := len(deck)
	for _, card := range deck {
		result += card * multiplier
		multiplier--
	}

	return result
}

func play(deck1 []int, deck2 []int) []int {
	var card1, card2 int
	for {
		if len(deck1) == 0 {
			return deck2
		} else if len(deck2) == 0 {
			return deck1
		}

		card1, deck1 = deck1[0], deck1[1:]
		card2, deck2 = deck2[0], deck2[1:]

		if card1 > card2 {
			deck1 = append(deck1, card1)
			deck1 = append(deck1, card2)
		} else {
			deck2 = append(deck2, card2)
			deck2 = append(deck2, card1)
		}
	}
}

func parse(input []string) ([]int, []int) {
	var deck1, deck2 []int

	section := 0
	for _, line := range input {
		if line == "" {
			section++
			continue
		}

		if strings.HasPrefix(line, "Player") {
			continue
		}

		if section == 0 {
			deck1 = append(deck1, convert.StringToInt(line))
		} else {
			deck2 = append(deck2, convert.StringToInt(line))
		}
	}
	return deck1, deck2
}
