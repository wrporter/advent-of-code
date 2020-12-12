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
	//defer timeit.Track(time.Now(), "part 1")
	player := Character{
		HitPoints: 50,
		Mana:      500,
	}
	boss := Character{
		HitPoints: 71,
		Damage:    10,
	}
	//player := Character{
	//	HitPoints: 10,
	//	Mana:      250,
	//}
	//boss := Character{
	//	HitPoints: 13,
	//	Damage:    8,
	//}
	//player := Character{
	//	HitPoints: 10,
	//	Mana:      250,
	//}
	//boss := Character{
	//	HitPoints: 14,
	//	Damage:    8,
	//}

	//minMana := ints.MaxInt
	//var minState *State
	minMana, _ := winMinMana(player, boss, "easy")
	//if minState != nil {
	//	fmt.Println(minState.Path)
	//}
	return minMana
}

func part2(input []string) interface{} {
	player := Character{
		HitPoints: 50,
		Mana:      500,
	}
	boss := Character{
		HitPoints: 71,
		Damage:    10,
	}

	minMana, _ := winMinMana(player, boss, "hard")
	return minMana
}

//boss = { hp: 71, damage: 10 }
//player = { hp: 50, mana: 500 }
//
//hard: 1937
//* Shield -> Recharge -> Poison -> Shield -> Recharge -> Poison -> Shield -> Recharge -> Poison -> Shield -> Magic Missile -> Poison -> Magic Missile
//
//easy: 1824
//* Poison -> Recharge -> Shield -> Poison -> Recharge -> Shield -> Poison -> Recharge -> Shield -> Magic Missile -> Poison -> Magic Missile
//* Poison -> Recharge -> Shield -> Poison -> Recharge -> Shield -> Poison -> Recharge -> Shield -> Poison -> Magic Missile -> Magic Missile
//* Recharge -> Poison -> Shield -> Recharge -> Poison -> Shield -> Recharge -> Poison -> Shield -> Magic Missile -> Poison -> Magic Missile

func winMinMana(startPlayer, startBoss Character, difficulty string) (int, *State) {
	var queue []*State
	var state *State
	minMana := ints.MaxInt
	var minState *State

	for _, spell := range spells {
		queue = append(queue, &State{Player: startPlayer, Boss: startBoss, Spell: spell})
	}

	for len(queue) > 0 {
		state, queue = queue[len(queue)-1], queue[:len(queue)-1]

		// Prune paths
		if state.Player.ManaSpent >= minMana {
			continue
		}

		//state.Path += "-- Player turn--\n"
		//state.Path += renderState(state)
		if difficulty == "hard" {
			state.Player.HitPoints -= 1
		}
		if state.Player.HitPoints <= 0 {
			continue
		}
		applyEffects(state)
		//state.Path += renderEffects(state.Effects)
		if state.Boss.HitPoints <= 0 {
			//state.Path += "The boss is dead, the player wins!\n"
			if state.Player.ManaSpent < minMana {
				minMana = state.Player.ManaSpent
				minState = state
			}
			continue
		}
		if state.Player.Mana < state.Spell.ManaCost {
			continue
		}
		//state.Path += renderSpell(state.Spell)

		state.Player.HitPoints += state.Spell.Heal
		state.Player.Mana -= state.Spell.ManaCost
		state.Player.ManaSpent += state.Spell.ManaCost
		state.Boss.HitPoints -= state.Spell.Damage
		if state.Spell.Effect.NumTurns > 0 {
			state.Effects[state.Spell.Name] = state.Spell.Effect
		}
		if state.Boss.HitPoints <= 0 {
			//state.Path += "The boss is dead, the player wins!\n"
			if state.Player.ManaSpent < minMana {
				minMana = state.Player.ManaSpent
				minState = state
			}
			continue
		}
		//state.Path += "\n"

		// Boss Turn
		//state.Path += "-- Boss turn--\n"
		//state.Path += renderState(state)
		applyEffects(state)
		if state.Boss.HitPoints <= 0 {
			//state.Path += "The boss is dead, the player wins!\n"
			if state.Player.ManaSpent < minMana {
				minMana = state.Player.ManaSpent
				minState = state
			}
			continue
		}
		state.Player.HitPoints -= state.Boss.Damage - state.Player.Armor
		//state.Path += fmt.Sprintf("Boss attacks for %d - %d = %d damage!", state.Boss.Damage, state.Player.Armor, state.Boss.Damage-state.Player.Armor)
		if state.Player.HitPoints <= 0 {
			continue
		}
		//state.Path += "\n\n"

		for _, spell := range spells {
			if !spellActive(state, spell) {
				nextState := copyState(state)
				nextState.Spell = spell
				queue = append(queue, nextState)
			}
		}
	}

	return minMana, minState
}

//func renderPath(state *State) {
//	for _, s := range state.Path {
//		printState(s)
//	}
//	printState(state)
//}

func renderState(s *State) string {
	result := ""
	result += fmt.Sprintf("- Player has %d hit points, %d armor, %d mana\n", s.Player.HitPoints, s.Player.Armor, s.Player.Mana)
	result += fmt.Sprintf("- Boss has %d hit points\n", s.Boss.HitPoints)
	return result
}

func renderSpell(spell Spell) string {
	result := fmt.Sprintf("Player casts %s", spell.Name)
	switch spell.Name {
	case "Magic Missile":
		result += ", dealing 4 damage.\n"
	case "Drain":
		result += ", dealing 2 damage, and healing 2 hit points.\n"
	case "Recharge":
		result += ".\n"
	case "Shield":
		result += ", increasing armor by 7.\n"
	case "Poison":
		result += ".\n"
	}
	return result
}

func renderEffects(effects map[string]Effect) string {
	result := ""
	for _, effect := range effects {
		result += renderEffect(effect)
	}
	return result
}

func renderEffect(effect Effect) string {
	switch effect.Name {
	case "Recharge":
		return fmt.Sprintf("Recharge provides 101 mana; its timer is now %d.\n", effect.NumTurns)
	case "Shield":
		return fmt.Sprintf("Shield's timer is now %d.\n", effect.NumTurns)
	case "Poison":
		return fmt.Sprintf("Poison deals 3 damage; its timer is now %d.\n", effect.NumTurns)
	}
	return ""
}

func spellActive(state *State, spell Spell) bool {
	effect, ok := state.Effects[spell.Name]
	return ok && effect.NumTurns > 1
}

func canCastSpell(player Character) bool {
	for _, spell := range spells {
		if player.Mana >= spell.ManaCost {
			return true
		}
	}
	return false
}

func applyEffects(state *State) {
	remainingEffects := make(map[string]Effect)

	for _, effect := range state.Effects {
		effect.NumTurns--
		effect.Apply(effect, state)

		if effect.NumTurns != 0 {
			remainingEffects[effect.Name] = effect
		} else {
			effect.Expire(effect, state)
		}
	}

	state.Effects = remainingEffects
}

func copyEffects(effects map[string]Effect) map[string]Effect {
	newEffects := make(map[string]Effect)
	for name, effect := range effects {
		newEffects[name] = effect
	}
	return newEffects
}

func copyState(state *State) *State {
	return &State{
		Player:  state.Player,
		Boss:    state.Boss,
		Effects: copyEffects(state.Effects),
		//Path:    state.Path,
		Spell: state.Spell,
	}
}

type (
	State struct {
		Player  Character
		Boss    Character
		Effects map[string]Effect
		Spell   Spell
		//Path    string
	}

	Character struct {
		HitPoints int
		Damage    int
		Mana      int
		Armor     int
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
		Name     string
		Self     bool
		NumTurns int
		Apply    func(effect Effect, state *State)
		Expire   func(effect Effect, state *State)
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
			Name:     "Shield",
			Self:     true,
			NumTurns: 6,
			Apply: func(effect Effect, state *State) {
				state.Player.Armor = 7
				//state.Path += fmt.Sprintf("Shield's timer is now %d.\n", effect.NumTurns)
			},
			Expire: func(effect Effect, state *State) {
				state.Player.Armor = 0
				//state.Path += "Shield wears off, decreasing armor by 7.\n"
			},
		},
	},
	{
		Name:     "Poison",
		ManaCost: 173,
		Effect: Effect{
			Name:     "Poison",
			Self:     false,
			NumTurns: 6,
			Apply: func(effect Effect, state *State) {
				state.Boss.HitPoints -= 3
				//state.Path += fmt.Sprintf("Poison deals 3 damage; its timer is now %d.\n", effect.NumTurns)
			},
			Expire: func(effect Effect, state *State) {
				//state.Path += "Poison wears off.\n"
			},
		},
	},
	{
		Name:     "Recharge",
		ManaCost: 229,
		Effect: Effect{
			Name:     "Recharge",
			Self:     true,
			NumTurns: 5,
			Apply: func(effect Effect, state *State) {
				state.Player.Mana += 101
				//state.Path += fmt.Sprintf("Recharge provides 101 mana; its timer is now %d.\n", effect.NumTurns)
			},
			Expire: func(effect Effect, state *State) {
				//state.Path += "Recharge wears off.\n"
			},
		},
	},
}
