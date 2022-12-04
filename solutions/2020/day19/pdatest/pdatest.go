package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/2020/day19/pda"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

func main() {
	input, _ := file.ReadFile("./2020/day19/pdatest/test1.txt")
	p := pda.NewPDA("0")

	messages, rules := parse(input)

	rules["8"] = Rule{
		Symbol: "8",
		Sequences: [][]string{
			{"42"},
			{"42", "8"},
			{"!"},
		},
	}
	rules["11"] = Rule{
		Symbol: "11",
		Sequences: [][]string{
			{"42", "31"},
			{"42", "11", "31"},
			{"!"},
		},
	}

	for _, rule := range rules {
		for _, sequence := range rule.Sequences {
			p.AddRule(rule.Symbol, sequence)
		}
	}

	count := 0
	for _, message := range messages {
		match := p.Match(message)
		if match {
			count++
		}
		fmt.Printf("%v\t%v\n", match, message)
	}
	fmt.Println(count)
}

type Rule struct {
	Symbol    string
	Sequences [][]string
}

func parse(input []string) ([]string, map[string]Rule) {
	var messages []string
	rules := make(map[string]Rule)

	section := 0
	for _, line := range input {
		if line == "" {
			section++
			continue
		}

		if section == 0 {
			splitRule := strings.Split(line, ": ")
			derivation := strings.ReplaceAll(splitRule[1], "\"", "")

			rule := Rule{
				Symbol: splitRule[0],
			}

			for _, sequence := range strings.Split(derivation, " | ") {
				rule.Sequences = append(rule.Sequences, strings.Split(sequence, " "))
			}

			rules[rule.Symbol] = rule
		} else if section == 1 {
			messages = append(messages, line)
		}
	}

	return messages, rules
}
