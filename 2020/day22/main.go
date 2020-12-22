package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

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
	winner := playRecursive(player1, player2, 1)
	return score(winner)
}

type Player struct {
	ID   int
	Deck []int
}

func NewPlayer(id int) Player {
	return Player{
		ID: id,
	}
}

func playRecursive(player1 Player, player2 Player, gameDepth int) Player {
	seen := make(map[string]bool)
	var card1, card2, winnerID int

	for round := 1; ; round++ {
		if len(player1.Deck) == 0 {
			return player2
		}
		if len(player2.Deck) == 0 {
			return player1
		}

		alreadySeen := seen[player1.String()] || seen[player2.String()]
		seen[player1.String()] = true
		seen[player2.String()] = true

		card1, player1.Deck = player1.Deck[0], player1.Deck[1:]
		card2, player2.Deck = player2.Deck[0], player2.Deck[1:]

		if alreadySeen {
			winnerID = player1.ID
		} else if card1 <= len(player1.Deck) && card2 <= len(player2.Deck) {
			player1Copy := player1.Copy(card1)
			player2Copy := player2.Copy(card2)
			winnerID = playRecursive(player1Copy, player2Copy, gameDepth+1).ID
		} else if card1 > card2 {
			winnerID = player1.ID
		} else {
			winnerID = player2.ID
		}

		if winnerID == player1.ID {
			player1.Deck = append(player1.Deck, card1)
			player1.Deck = append(player1.Deck, card2)
		} else {
			player2.Deck = append(player2.Deck, card2)
			player2.Deck = append(player2.Deck, card1)
		}
	}
}

func (p Player) String() string {
	var sb strings.Builder
	delimiter := ','
	sb.WriteString(fmt.Sprintf("%d: ", p.ID))

	for i, card := range p.Deck {
		sb.WriteString(fmt.Sprintf("%d", card))
		if i < len(p.Deck)-1 {
			sb.WriteRune(delimiter)
		}
	}

	return sb.String()
}

func (p Player) Copy(size int) Player {
	result := NewPlayer(p.ID)
	result.Deck = make([]int, size)
	for i := 0; i < size; i++ {
		result.Deck[i] = p.Deck[i]
	}
	return result
}

func score(winner Player) int {
	result := 0

	multiplier := len(winner.Deck)
	for _, card := range winner.Deck {
		result += card * multiplier
		multiplier--
	}

	return result
}

func play(player1 Player, player2 Player) Player {
	var card1, card2 int
	for {
		if len(player1.Deck) == 0 {
			return player2
		} else if len(player2.Deck) == 0 {
			return player1
		}

		card1, player1.Deck = player1.Deck[0], player1.Deck[1:]
		card2, player2.Deck = player2.Deck[0], player2.Deck[1:]

		if card1 > card2 {
			player1.Deck = append(player1.Deck, card1)
			player1.Deck = append(player1.Deck, card2)
		} else {
			player2.Deck = append(player2.Deck, card2)
			player2.Deck = append(player2.Deck, card1)
		}
	}
}

func parse(input []string) (Player, Player) {
	player1 := NewPlayer(1)
	player2 := NewPlayer(2)

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
			player1.Deck = append(player1.Deck, conversion.StringToInt(line))
		} else {
			player2.Deck = append(player2.Deck, conversion.StringToInt(line))
		}
	}
	return player1, player2
}
