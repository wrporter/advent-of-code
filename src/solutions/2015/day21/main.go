package main

import (
	"aoc/src/lib/go/ints"
	"fmt"
)

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type Shop struct {
	weapons []Item
	armor   []Item
	rings   []Item
}

var itemShop = Shop{
	weapons: []Item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	},
	armor: []Item{
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	},
	rings: []Item{
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	},
}

type EquipAmount struct {
	min int
	max int
}

type EquipTypes struct {
	weapons EquipAmount
	armor   EquipAmount
	rings   EquipAmount
}

type Character struct {
	HitPoints int
	Damage    int
	Armor     int
}

func (c *Character) Equip(item Item) {
	c.Damage += item.damage
	c.Armor += item.armor
}

var equipment = EquipTypes{
	weapons: EquipAmount{1, 1},
	armor:   EquipAmount{0, 1},
	rings:   EquipAmount{0, 2},
}

func PermuteSize(values []Item, startSize int, endSize int, emit func([]Item)) {
	var permuteSize func([]Item, int, int)

	permuteSize = func(current []Item, index int, size int) {
		if len(current) == size {
			emit(current)
			return
		}

		for i := index; i < len(values); i++ {
			current = append(current, values[i])
			permuteSize(current, i+1, size)
			current = current[:len(current)-1]
		}
	}

	for size := startSize; size <= endSize; size++ {
		permuteSize(nil, 0, size)
	}
}

func playerWins(player Character, boss Character) bool {
	for {
		boss.HitPoints -= attackDamage(player, boss)
		if boss.HitPoints <= 0 {
			return true
		}
		player.HitPoints -= attackDamage(boss, player)
		if player.HitPoints <= 0 {
			return false
		}
	}
}

func attackDamage(attacker Character, defender Character) int {
	if attacker.Damage <= defender.Armor {
		return 1
	}
	return attacker.Damage - defender.Armor
}

type Purchase struct {
	gold  int
	items []Item
}

func minGoldToDefeatBoss() (Purchase, Purchase) {
	minWinningPurchase := Purchase{
		gold:  ints.MaxInt,
		items: nil,
	}
	maxLosingPurchase := Purchase{
		gold:  0,
		items: nil,
	}
	boss := Character{
		HitPoints: 100,
		Damage:    8,
		Armor:     2,
	}

	PermuteSize(itemShop.weapons, equipment.weapons.min, equipment.weapons.max, func(weapons []Item) {
		PermuteSize(itemShop.armor, equipment.armor.min, equipment.armor.max, func(armor []Item) {
			PermuteSize(itemShop.rings, equipment.rings.min, equipment.rings.max, func(rings []Item) {
				cost := 0
				player := &Character{
					HitPoints: 100,
					Damage:    0,
					Armor:     0,
				}
				for _, item := range weapons {
					cost += item.cost
					player.Equip(item)
				}
				for _, item := range armor {
					cost += item.cost
					player.Equip(item)
				}
				for _, item := range rings {
					cost += item.cost
					player.Equip(item)
				}

				if playerWins(*player, boss) && cost < minWinningPurchase.gold {
					minWinningPurchase.gold = cost
					minWinningPurchase.items = make([]Item, 0)
					minWinningPurchase.items = append(minWinningPurchase.items, weapons...)
					minWinningPurchase.items = append(minWinningPurchase.items, armor...)
					minWinningPurchase.items = append(minWinningPurchase.items, rings...)
				}

				if !playerWins(*player, boss) && cost > maxLosingPurchase.gold {
					maxLosingPurchase.gold = cost
					maxLosingPurchase.items = make([]Item, 0)
					maxLosingPurchase.items = append(maxLosingPurchase.items, weapons...)
					maxLosingPurchase.items = append(maxLosingPurchase.items, armor...)
					maxLosingPurchase.items = append(maxLosingPurchase.items, rings...)
				}
			})
		})
	})

	return minWinningPurchase, maxLosingPurchase
}

func main() {
	fmt.Println(minGoldToDefeatBoss())
}
