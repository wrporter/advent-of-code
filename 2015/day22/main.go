package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
)

func main() {
	year, day := 2015, 22
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	//player := Character{
	//	HitPoints: 50,
	//	Mana:      500,
	//}
	//boss := Character{
	//	HitPoints: 71,
	//	Damage:    10,
	//}
	//player := Character{
	//	HitPoints: 10,
	//	Mana:      250,
	//}
	//boss := Character{
	//	HitPoints: 13,
	//	Damage:    9,
	//}
	player := Character{
		HitPoints: 10,
		Mana:      250,
	}
	boss := Character{
		HitPoints: 14,
		Damage:    9,
	}

	return winMinMana(player, boss)
}

func part2(input []string) interface{} {
	return 0
}

func winMinMana(startPlayer, startBoss Character) int {
	queue := []State{{startPlayer, startBoss}}
	var state State
	minMana := ints.MaxInt

	for len(queue) > 0 {
		state, queue = queue[len(queue)-1], queue[:len(queue)-1]

		if state.Player.HitPoints <= 0 {
			continue
		}

		for _, spell := range spells {
			if state.Player.Mana < spell.ManaCost || spellActive(state.Player, spell) || spellActive(state.Boss, spell) {
				continue
			}
			nextState := copyState(state)
			player, boss := &nextState.Player, &nextState.Boss

			// Player Turn
			applyEffects(player)
			applyEffects(boss)
			if !canCastSpell(*player) {
				break
			}
			if boss.HitPoints <= 0 {
				minMana = ints.Min(minMana, player.ManaSpent)
				continue
			}

			player.HitPoints += spell.Heal
			player.Mana -= spell.ManaCost
			player.ManaSpent += spell.ManaCost
			boss.HitPoints -= spell.Damage
			if spell.Effect.NumTurns > 0 {
				if spell.Effect.Self {
					player.Effects = append(player.Effects, spell.Effect)
				} else {
					boss.Effects = append(boss.Effects, spell.Effect)
				}
			}

			// Boss Turn
			applyEffects(player)
			applyEffects(boss)
			if boss.HitPoints <= 0 {
				minMana = ints.Min(minMana, player.ManaSpent)
				continue
			}
			player.HitPoints -= boss.Damage - player.Armor

			queue = append(queue, nextState)
		}
	}

	return minMana
}

func spellActive(char Character, spell Spell) bool {
	for _, effect := range char.Effects {
		if spell.Effect.Name == effect.Name && effect.NumTurns > 1 {
			return true
		}
	}
	return false
}

func canCastSpell(player Character) bool {
	for _, spell := range spells {
		if player.Mana >= spell.ManaCost {
			return true
		}
	}
	return false
}

func applyEffects(character *Character) {
	var remainingEffects []Effect
	for _, effect := range character.Effects {
		character.Armor = effect.ArmorIncrease
		character.Mana += effect.AddMana
		character.HitPoints -= effect.Damage

		effect.NumTurns--
		if effect.NumTurns != 0 {
			remainingEffects = append(remainingEffects, effect)
		}
	}
	character.Effects = remainingEffects
}

func copyCharacter(char Character) Character {
	newChar := char
	newChar.Effects = make([]Effect, len(char.Effects))
	copy(newChar.Effects, char.Effects)
	return newChar
}

func copyState(state State) State {
	return State{
		Player: copyCharacter(state.Player),
		Boss:   copyCharacter(state.Boss),
	}
}

type (
	State struct {
		Player Character
		Boss   Character
	}

	Character struct {
		HitPoints int
		Damage    int
		Mana      int
		Armor     int
		Effects   []Effect
		ManaSpent int
	}

	Spell struct {
		Name     string
		ManaCost int
		Damage   int
		Heal     int
		Effect   Effect
	}

	Effect struct {
		Name          string
		Self          bool
		NumTurns      int
		ArmorIncrease int
		Damage        int
		AddMana       int
	}
)

var spells = []Spell{
	{
		Name:     "Magic Missile",
		ManaCost: 53,
		Damage:   4,
	},
	{
		Name:     "Drain",
		ManaCost: 73,
		Damage:   2,
		Heal:     2,
	},
	{
		Name:     "Shield",
		ManaCost: 113,
		Effect: Effect{
			Name:          "Shield",
			Self:          true,
			NumTurns:      6,
			ArmorIncrease: 7,
		},
	},
	{
		Name:     "Poison",
		ManaCost: 173,
		Effect: Effect{
			Name:     "Poison",
			Self:     false,
			NumTurns: 6,
			Damage:   3,
		},
	},
	{
		Name:     "Recharge",
		ManaCost: 229,
		Effect: Effect{
			Name:     "Recharge",
			Self:     true,
			NumTurns: 5,
			AddMana:  101,
		},
	},
}
