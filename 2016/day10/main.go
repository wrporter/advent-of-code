package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
	"sort"
)

var botRegex = regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)
var valueRegex = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)

type Bot struct {
	ID    int
	Chips []int
}

type GiveInstruction struct {
	GiverID int
	LowTo   string
	LowID   int
	HighTo  string
	HighID  int
}

func part1(instructions []string) (int, int) {
	bots := make(map[int]*Bot)
	outputs := make(map[int][]int)
	var giveInstructions []GiveInstruction
	botID := -1

	for _, instruction := range instructions {
		if valueRegex.MatchString(instruction) {
			match := valueRegex.FindStringSubmatch(instruction)
			chip := conversion.StringToInt(match[1])
			botID := conversion.StringToInt(match[2])

			if bot, ok := bots[botID]; ok {
				bot.Chips = append(bot.Chips, chip)
				sort.Ints(bot.Chips)
			} else {
				bots[botID] = &Bot{ID: botID, Chips: []int{chip}}
			}
		} else if botRegex.MatchString(instruction) {
			match := botRegex.FindStringSubmatch(instruction)
			giveInstructions = append(giveInstructions, GiveInstruction{
				GiverID: conversion.StringToInt(match[1]),
				LowTo:   match[2],
				LowID:   conversion.StringToInt(match[3]),
				HighTo:  match[4],
				HighID:  conversion.StringToInt(match[5]),
			})
		}
	}

	for len(giveInstructions) > 0 {
		instruction := giveInstructions[0]
		giveInstructions = giveInstructions[1:]

		if giveBot, ok := bots[instruction.GiverID]; ok && len(giveBot.Chips) >= 2 {
			if giveBot.Chips[0] == 17 && giveBot.Chips[len(giveBot.Chips)-1] == 61 {
				botID = giveBot.ID
			}

			if instruction.LowTo == "output" {
				outputs[instruction.LowID] = append(outputs[instruction.LowID], giveBot.Chips[0])
			} else if instruction.LowTo == "bot" {
				if receiveBot, ok := bots[instruction.LowID]; ok {
					receiveBot.Chips = append(receiveBot.Chips, giveBot.Chips[0])
					sort.Ints(receiveBot.Chips)
				} else {
					bots[instruction.LowID] = &Bot{ID: instruction.LowID, Chips: []int{giveBot.Chips[0]}}
				}
			}
			giveBot.Chips = giveBot.Chips[1:]

			if instruction.HighTo == "output" {
				outputs[instruction.HighID] = append(outputs[instruction.HighID], giveBot.Chips[len(giveBot.Chips)-1])
			} else if instruction.HighTo == "bot" {
				if receiveBot, ok := bots[instruction.HighID]; ok {
					receiveBot.Chips = append(receiveBot.Chips, giveBot.Chips[len(giveBot.Chips)-1])
					sort.Ints(receiveBot.Chips)
				} else {
					bots[instruction.HighID] = &Bot{ID: instruction.HighID, Chips: []int{giveBot.Chips[len(giveBot.Chips)-1]}}
				}
			}
			giveBot.Chips = giveBot.Chips[:len(giveBot.Chips)-1]

		} else {
			giveInstructions = append(giveInstructions, instruction)
		}
	}

	return botID, outputs[0][0] * outputs[1][0] * outputs[2][0]
}

func main() {
	input, _ := file.ReadFile("./2016/day10/input.txt")
	answer1, answer2 := part1(input)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
