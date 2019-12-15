package nanofactory

import (
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	mymath "github.com/wrporter/advent-of-code-2019/internal/common/math"
	"math"
	"regexp"
	"sort"
	"strings"
)

type NanoFactory struct{}

type Reaction struct {
	Input  []Chemical
	Output Chemical
}

type Chemical struct {
	Name   string
	Amount int
}

const (
	Ore  = "ORE"
	Fuel = "FUEL"
)

var chemicalRegex = regexp.MustCompile(`^(\d+) ([A-Z]+)$`)

func New() *NanoFactory {
	return &NanoFactory{}
}

func (n *NanoFactory) GetRequiredOre(reactionStrings []string, fuel int) int {
	reactions := parseReactions(reactionStrings)
	tree := map[string]*Chemical{Fuel: {Fuel, fuel}}
	toProcess := []Chemical{{Fuel, 1}}
	var next Chemical

	for len(toProcess) > 0 {
		next, toProcess = poll(toProcess)
		if _, processable := tree[next.Name]; !processable {
			continue
		}

		need := leastAmount(tree, next, reactions)
		delete(tree, next.Name)

		for _, input := range reactions[next.Name].Input {
			var chemical Chemical
			if _, ok := tree[input.Name]; ok {
				tree[input.Name].Amount += input.Amount * need
				chemical = Chemical{input.Name, input.Amount}
			} else {
				chemical = Chemical{input.Name, input.Amount * need}
				tree[input.Name] = &chemical
			}
			if chemical.Name != Ore {
				toProcess = append(toProcess, chemical)
			}
		}

		sort.Slice(toProcess, func(i, j int) bool {
			return distanceToOre(reactions, toProcess[i].Name) > distanceToOre(reactions, toProcess[j].Name)
		})
	}

	return tree[Ore].Amount
}

func distanceToOre(tree map[string]Reaction, chemical string) int {
	return distanceToOreRec(tree, chemical, 1)
}

func distanceToOreRec(tree map[string]Reaction, chemical string, level int) int {
	distance := level
	for _, input := range tree[chemical].Input {
		if input.Name == Ore {
			return level
		} else {
			nextDist := distanceToOreRec(tree, input.Name, level+1)
			distance = mymath.Max(distance, nextDist)
		}
	}
	return distance
}

func leastAmount(tree map[string]*Chemical, next Chemical, reactions map[string]Reaction) int {
	return int(math.Ceil(
		float64(tree[next.Name].Amount) /
			float64(reactions[next.Name].Output.Amount),
	))
}

func onlyInputIsOre(reactions map[string]Reaction, input Chemical) bool {
	return len(reactions[input.Name].Input) == 1 && reactions[input.Name].Input[0].Name == Ore
}

func parseReactions(reactionStrings []string) map[string]Reaction {
	reactions := make(map[string]Reaction)

	for _, reactionString := range reactionStrings {
		reaction := parseReaction(reactionString)
		reactions[reaction.Output.Name] = reaction
	}
	return reactions
}

func parseReaction(reactionString string) Reaction {
	equation := strings.Split(reactionString, " => ")
	input := parseInput(equation)
	output := parseChemical(equation[1])
	reaction := Reaction{input, output}
	return reaction
}

func parseInput(equation []string) []Chemical {
	var input []Chemical
	inputStrings := strings.Split(equation[0], ", ")
	for _, str := range inputStrings {
		input = append(input, parseChemical(str))
	}
	return input
}

func parseChemical(chemical string) Chemical {
	match := chemicalRegex.FindStringSubmatch(chemical)
	return Chemical{
		Name:   match[2],
		Amount: conversion.StringToInt(match[1]),
	}
}

func reverse(reactions []Reaction) {
	for i, j := 0, len(reactions)-1; i < j; i, j = i+1, j-1 {
		reactions[i], reactions[j] = reactions[j], reactions[i]
	}
}

func poll(array []Chemical) (Chemical, []Chemical) {
	return array[0], array[1:]
}
