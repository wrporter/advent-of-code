package solution

import (
	"aoc/src/lib/go/convert"
	"regexp"
	"sort"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	monkeys := parseInput(input)
	return calculateMonkeyBusiness(monkeys, 20, func(worry int) int {
		return worry / 3
	})
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	monkeys := parseInput(input)

	// Chinese Remainder Theorem: Get product of all prime moduli
	modulus := 1
	for _, m := range monkeys {
		modulus *= m.divisibleBy
	}

	return calculateMonkeyBusiness(monkeys, 10_000, func(worry int) int {
		return worry % modulus
	})
}

func calculateMonkeyBusiness(monkeys []*monkey, rounds int, manageWorry func(worry int) int) interface{} {
	for round := 0; round < rounds; round += 1 {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.numItemsInspected += 1
				item = m.operation(item)
				item = manageWorry(item)

				if item%m.divisibleBy == 0 {
					monkeys[m.trueMonkey].items = append(monkeys[m.trueMonkey].items, item)
				} else {
					monkeys[m.falseMonkey].items = append(monkeys[m.falseMonkey].items, item)
				}
			}
			m.items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].numItemsInspected > monkeys[j].numItemsInspected
	})

	return monkeys[0].numItemsInspected * monkeys[1].numItemsInspected
}

var regex = regexp.MustCompile(`^Monkey (\d+):
 {2}Starting items: (\d+(?:, \d+)*)
 {2}Operation: new = old ([+*]) (old|\d+)
 {2}Test: divisible by (\d+)
 {4}If true: throw to monkey (\d+)
 {4}If false: throw to monkey (\d+)$`)

func parseInput(input string) []*monkey {
	var monkeys []*monkey
	for _, group := range strings.Split(input, "\n\n") {
		match := regex.FindStringSubmatch(group)

		operator := match[3]
		modifier := match[4]
		modifierValue := convert.StringToInt(match[4])

		m := &monkey{
			id:    convert.StringToInt(match[1]),
			items: convert.ToIntsV2(strings.Split(match[2], ", ")),
			operation: func(level int) int {
				amount := modifierValue
				if modifier == "old" {
					amount = level
				}

				if operator == "+" {
					return level + amount
				} else if operator == "*" {
					return level * amount
				}
				return level
			},
			divisibleBy:       convert.StringToInt(match[5]),
			trueMonkey:        convert.StringToInt(match[6]),
			falseMonkey:       convert.StringToInt(match[7]),
			numItemsInspected: 0,
		}

		monkeys = append(monkeys, m)
	}

	return monkeys
}

type monkey struct {
	id                int
	items             []int
	operation         func(level int) int
	divisibleBy       int
	trueMonkey        int
	falseMonkey       int
	numItemsInspected int
}
