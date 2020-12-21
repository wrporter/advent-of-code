package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 16
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

type Range struct {
	Low  int
	High int
}

type Rule struct {
	Range1 Range
	Range2 Range
}

var (
	ruleRegex = regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
)

func part1(input []string) interface{} {
	rules, _, nearbyTickets := parse(input)

	ticketScanningErrorRate := 0
	for _, ticket := range nearbyTickets {
		for _, value := range ticket {
			if !isValid(rules, value) {
				ticketScanningErrorRate += value
			}
		}
	}

	return ticketScanningErrorRate
}

func part2(input []string) interface{} {
	rules, yourTicket, nearbyTickets := parse(input)

	var validTickets [][]int
	for _, ticket := range nearbyTickets {
		if isValidTicket(ticket, rules) {
			validTickets = append(validTickets, ticket)
		}
	}

	validTickets = append(validTickets, yourTicket)

	numColumns := len(yourTicket)
	columns := make([]string, numColumns)
	taken := make(map[string]bool)
	potentialColumns := make([]map[string]bool, numColumns)
	for col := 0; col < len(yourTicket); col++ {
		buckets := make(map[string]int) // rule name to number of tickets

		for _, ticket := range validTickets {
			for ruleName, rule := range rules {
				if meetsRule(ticket[col], rule) {
					buckets[ruleName]++
				}
			}
		}

		columnRules := make(map[string]bool)
		for ruleName, count := range buckets {
			if count == len(validTickets) {
				if !taken[ruleName] {
					columns[col] = ruleName
					taken[ruleName] = true
				}
				columnRules[ruleName] = true
				potentialColumns[col] = columnRules
			}
		}
	}

	unassigned := make(map[string]bool)
	for ruleName := range rules {
		unassigned[ruleName] = true
	}
	assigned := make(map[string]int)

	for len(unassigned) > 0 {
		for rule := range unassigned {
			numCandidates := 0
			column := -1
			for col, colRules := range potentialColumns {
				if colRules[rule] {
					numCandidates++
					column = col
				}
			}
			if _, ok := assigned[rule]; !ok && numCandidates == 1 {
				assigned[rule] = column
				delete(unassigned, rule)
				potentialColumns[column] = map[string]bool{}
				break
			}
		}
	}

	result := 1
	for ruleName, col := range assigned {
		if strings.HasPrefix(ruleName, "departure") {
			result *= yourTicket[col]
		}
	}

	return result
}

func isValidTicket(ticket []int, rules map[string]Rule) bool {
	for _, value := range ticket {
		if !isValid(rules, value) {
			return false
		}
	}
	return true
}

func parse(input []string) (map[string]Rule, []int, [][]int) {
	rules := make(map[string]Rule)
	var yourTicket []int
	var nearbyTickets [][]int

	section := 0
	for _, line := range input {
		if line == "" {
			section++
			continue
		}

		if section == 0 {
			match := ruleRegex.FindStringSubmatch(line)
			rule := Rule{
				Range1: Range{
					Low:  conversion.StringToInt(match[2]),
					High: conversion.StringToInt(match[3]),
				},
				Range2: Range{
					Low:  conversion.StringToInt(match[4]),
					High: conversion.StringToInt(match[5]),
				},
			}
			rules[match[1]] = rule
		} else if section == 1 && !strings.HasPrefix(line, "your ticket") {
			yourTicket, _ = conversion.ToInts(strings.Split(line, ","))
		} else if section == 2 && !strings.HasPrefix(line, "nearby tickets") {
			tickets, _ := conversion.ToInts(strings.Split(line, ","))
			nearbyTickets = append(nearbyTickets, tickets)
		}
	}
	return rules, yourTicket, nearbyTickets
}

func isValid(rules map[string]Rule, ticket int) bool {
	for _, rule := range rules {
		if meetsRule(ticket, rule) {
			return true
		}
	}
	return false
}

func meetsRule(value int, rule Rule) bool {
	return (value >= rule.Range1.Low && value <= rule.Range1.High) ||
		(value >= rule.Range2.Low && value <= rule.Range2.High)
}
