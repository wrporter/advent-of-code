package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"strings"
)

var debugFlag = false

func main() {
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

func debug(str string) {
	if debugFlag {
		fmt.Println(str)
	}
}

//- Before either player deals a card, if there was a previous round in this game that had exactly the same cards in the
// same order in the same players' decks, the game instantly ends in a win for player 1. Previous rounds from other games
// are not considered. (This prevents infinite games of Recursive Combat, which everyone agrees is a bad idea.)
//- Otherwise, this round's cards must be in a new configuration; the players begin the round by each drawing the top card
// of their deck as normal.
//- If both players have at least as many cards remaining in their deck as the value of the card they just drew, the winner
// of the round is determined by playing a new game of Recursive Combat (see below).
//- Otherwise, at least one player must not have enough cards left in their deck to recurse; the winner of the round is
// the player with the higher-value card.
func playRecursive(player1 Player, player2 Player, gameDepth int) Player {
	debug(fmt.Sprintf("=== Game %d ===", gameDepth))
	seen := make(map[string]bool)
	var card1, card2 int

	for round := 1; ; round++ {
		debug(fmt.Sprintf("\n-- Round %d (Game %d) --", round, gameDepth))
		debug(player1.Render())
		debug(player2.Render())

		if len(player1.Deck) == 0 {
			debug(fmt.Sprintf("The winner of gameDepth %d is player 2!", gameDepth))
			return player2
		}
		if len(player2.Deck) == 0 {
			debug(fmt.Sprintf("The winner of gameDepth %d is player 1!", gameDepth))
			return player1
		}

		if seen[player1.String()] || seen[player2.String()] {
			card1, player1.Deck = player1.Deck[0], player1.Deck[1:]
			card2, player2.Deck = player2.Deck[0], player2.Deck[1:]
			player1.Deck = append(player1.Deck, card1)
			player1.Deck = append(player1.Deck, card2)
			seen[player1.String()] = true
			seen[player2.String()] = true
			continue
		}

		seen[player1.String()] = true
		seen[player2.String()] = true

		card1, player1.Deck = player1.Deck[0], player1.Deck[1:]
		card2, player2.Deck = player2.Deck[0], player2.Deck[1:]

		debug(fmt.Sprintf("Player 1 plays: %d", card1))
		debug(fmt.Sprintf("Player 2 plays: %d", card2))

		if card1 <= len(player1.Deck) && card2 <= len(player2.Deck) {
			debug("Playing a sub-gameDepth to determine the winner...\n")
			player1Copy := player1.Copy(card1)
			player2Copy := player2.Copy(card2)
			winner := playRecursive(player1Copy, player2Copy, gameDepth+1)
			debug(fmt.Sprintf("\n...anyway, back to gameDepth %d.", gameDepth))
			if winner.ID == player1.ID {
				debug(fmt.Sprintf("Player 1 wins round %d of gameDepth %d!", round, gameDepth))
				player1.Deck = append(player1.Deck, card1)
				player1.Deck = append(player1.Deck, card2)
			} else {
				debug(fmt.Sprintf("Player 2 wins round %d of gameDepth %d!", round, gameDepth))
				player2.Deck = append(player2.Deck, card2)
				player2.Deck = append(player2.Deck, card1)
			}
		} else {
			if card1 > card2 {
				debug(fmt.Sprintf("Player 1 wins round %d of gameDepth %d!", round, gameDepth))
				player1.Deck = append(player1.Deck, card1)
				player1.Deck = append(player1.Deck, card2)
			} else {
				debug(fmt.Sprintf("Player 2 wins round %d of gameDepth %d!", round, gameDepth))
				player2.Deck = append(player2.Deck, card2)
				player2.Deck = append(player2.Deck, card1)
			}
		}
	}
}

func (p Player) Render() string {
	var sb strings.Builder
	delimiter := ", "
	sb.WriteString(fmt.Sprintf("Player %d's deck: ", p.ID))

	for i, card := range p.Deck {
		sb.WriteString(fmt.Sprintf("%d", card))
		if i < len(p.Deck)-1 {
			sb.WriteString(delimiter)
		}
	}

	return sb.String()
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
