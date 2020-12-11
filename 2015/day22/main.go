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
	player := Character{
		HitPoints: 10,
		Mana:      250,
	}
	boss := Character{
		HitPoints: 13,
		Damage:    9,
	}

	min := ints.MaxInt
	for _, spell := range spells {
		won, mana := fight(copyCharacter(player), copyCharacter(boss), &spell)
		if won {
			min = ints.Min(min, mana)
		}
	}

	return min
}

func part2(input []string) interface{} {
	return 0
}

func fight(player, boss *Character, spell *Spell) (bool, int) {
	minMana := player.Mana
	win := false
	if player.HitPoints <= 0 || boss.HitPoints <= 0 /* OR cannot cast any spell */ {
		return player.HitPoints > 0 /* AND can cast any spell */, player.Mana
	}

	// Player Turn
	applyEffects(player)
	applyEffects(boss)
	player.HitPoints += spell.Heal
	player.Mana -= spell.ManaCost
	boss.HitPoints -= spell.Damage
	if spell.Effect != nil {
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
		return true, player.Mana
	}
	player.HitPoints -= boss.Damage

	for _, spell := range spells {
		if spell.ManaCost <= player.Mana {
			playerCopy := copyCharacter(player)
			bossCopy := copyCharacter(boss)

			won, mana := fight(playerCopy, bossCopy, &spell)
			minMana = ints.Min(minMana, mana)
			if won {
				return true, minMana
			}
			win = won
		}
	}

	return win, minMana
}

func applyEffects(character *Character) {
	var remainingEffects []*Effect
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

func copyCharacter(char *Character) *Character {
	return &Character{
		HitPoints: char.HitPoints,
		Damage:    char.Damage,
		Mana:      char.Mana,
		Armor:     char.Armor,
		Effects:   char.Effects,
	}
}

type (
	Character struct {
		HitPoints int
		Damage    int
		Mana      int
		Armor     int
		Effects   []*Effect
	}

	Spell struct {
		Name     string
		ManaCost int
		Damage   int
		Heal     int
		Effect   *Effect
	}

	Effect struct {
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
		Effect: &Effect{
			Self:          true,
			NumTurns:      6,
			ArmorIncrease: 7,
		},
	},
	{
		Name:     "Poison",
		ManaCost: 173,
		Effect: &Effect{
			Self:     false,
			NumTurns: 6,
			Damage:   3,
		},
	},
	{
		Name:     "Recharge",
		ManaCost: 229,
		Effect: &Effect{
			Self:          true,
			NumTurns:      6,
			ArmorIncrease: 7,
		},
	},
}
