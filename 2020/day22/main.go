package main

import (
	"container/list"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"strings"
)

var debugFlag = true

func main() {
	year, day := 2020, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/sample-input.txt", year, day))

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
	Deck *list.List
}

func NewPlayer(id int) Player {
	return Player{
		ID:   id,
		Deck: list.New(),
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
func playRecursive(player1 Player, player2 Player, game int) Player {
	debug(fmt.Sprintf("=== Game %d ===", game))
	seen := make(map[string]bool)

	for round := 1; ; round++ {
		debug(fmt.Sprintf("\n-- Round %d (Game %d) --", round, game))
		debug(player1.Render())
		debug(player2.Render())

		if seen[player1.String()] || seen[player2.String()] {
			player1.Deck.PushBack(player1.Deck.Remove(player1.Deck.Front()))
			player1.Deck.PushBack(player2.Deck.Remove(player2.Deck.Front()))
			seen[player1.String()] = true
			seen[player2.String()] = true
			continue
		}

		seen[player1.String()] = true
		seen[player2.String()] = true

		player1Card := player1.Deck.Remove(player1.Deck.Front())
		player2Card := player2.Deck.Remove(player2.Deck.Front())

		debug(fmt.Sprintf("Player 1 plays: %d", player1Card))
		debug(fmt.Sprintf("Player 2 plays: %d", player2Card))

		if player1Card.(int) <= player1.Deck.Len() && player2Card.(int) <= player2.Deck.Len() {
			debug("Playing a sub-game to determine the winner...\n")
			player1Copy := player1.Copy()
			player2Copy := player2.Copy()
			winner := playRecursive(player1Copy, player2Copy, game+1)
			debug(fmt.Sprintf("\n...anyway, back to game %d.", game))
			if winner.ID == player1.ID {
				debug(fmt.Sprintf("Player 1 wins round %d of game %d!", round, game))
				player1.Deck.PushBack(player1Card)
				player1.Deck.PushBack(player2Card)
			} else {
				debug(fmt.Sprintf("Player 2 wins round %d of game %d!", round, game))
				player2.Deck.PushBack(player2Card)
				player2.Deck.PushBack(player1Card)
			}
		} else {
			if player1Card.(int) > player2Card.(int) {
				debug(fmt.Sprintf("Player 1 wins round %d of game %d!", round, game))
				player1.Deck.PushBack(player1Card)
				player1.Deck.PushBack(player2Card)
			} else {
				debug(fmt.Sprintf("Player 2 wins round %d of game %d!", round, game))
				player2.Deck.PushBack(player2Card)
				player2.Deck.PushBack(player1Card)
			}
		}

		if player1.Deck.Len() == 0 {
			debug(fmt.Sprintf("The winner of game %d is player 2!", game))
			return player2
		}
		if player2.Deck.Len() == 0 {
			debug(fmt.Sprintf("The winner of game %d is player 1!", game))
			return player1
		}
	}
}

func copyDeck(deck *list.List) *list.List {
	result := list.New()
	for node := deck.Front(); node != nil; node = node.Next() {
		result.PushBack(node.Value)
	}
	return result
}

func (p Player) Render() string {
	var sb strings.Builder
	delimiter := ", "
	sb.WriteString(fmt.Sprintf("Player %d's deck: ", p.ID))

	for card := p.Deck.Front(); card != nil; card = card.Next() {
		sb.WriteString(fmt.Sprintf("%d", card.Value))
		if card.Next() != nil {
			sb.WriteString(delimiter)
		}
	}

	return sb.String()
}

func (p Player) String() string {
	var sb strings.Builder
	delimiter := ','
	sb.WriteString(fmt.Sprintf("%d: ", p.ID))

	for card := p.Deck.Front(); card != nil; card = card.Next() {
		sb.WriteString(fmt.Sprintf("%d", card.Value))
		if card.Next() != nil {
			sb.WriteRune(delimiter)
		}
	}

	return sb.String()
}

func (p Player) Copy() Player {
	result := NewPlayer(p.ID)
	result.Deck = copyDeck(p.Deck)
	return result
}

func score(winner Player) int {
	result := 0

	multiplier := 1
	for node := winner.Deck.Back(); node != nil; node = node.Prev() {
		result += node.Value.(int) * multiplier
		multiplier++
	}

	return result
}

func play(player1 Player, player2 Player) Player {
	for {
		if player1.Deck.Len() == 0 {
			return player2
		} else if player2.Deck.Len() == 0 {
			return player1
		}

		player1Card := player1.Deck.Remove(player1.Deck.Front())
		player2Card := player2.Deck.Remove(player2.Deck.Front())

		if player1Card.(int) > player2Card.(int) {
			player1.Deck.PushBack(player1Card)
			player1.Deck.PushBack(player2Card)
		} else {
			player2.Deck.PushBack(player2Card)
			player2.Deck.PushBack(player1Card)
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
			player1.Deck.PushBack(conversion.StringToInt(line))
		} else {
			player2.Deck.PushBack(conversion.StringToInt(line))
		}
	}
	return player1, player2
}
