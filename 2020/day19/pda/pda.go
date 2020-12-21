package pda

import (
	"fmt"
	"strings"
)

const (
	endSymbol   = "$"
	emptySymbol = "!"
)

type (
	PDA struct {
		startSymbol string
		transitions map[string][]transition
	}

	transition struct {
		symbol     string
		derivation []string
	}
)

func NewPDA(startSymbol string) *PDA {
	return &PDA{
		startSymbol: startSymbol,
		transitions: make(map[string][]transition),
	}
}

func (p *PDA) AddRules(rules []string) {
	for _, rule := range rules {
		parts := strings.Split(rule, ": ")
		symbol := parts[0]

		for _, sequence := range strings.Split(parts[1], " | ") {
			derivation := strings.Split(sequence, " ")
			p.AddRule(symbol, derivation)
		}
	}
}

func (p *PDA) AddRule(symbol string, derivation []string) {
	p.transitions[symbol] = append(p.transitions[symbol], transition{
		symbol:     symbol,
		derivation: derivation,
	})
}

func (p *PDA) Match(input string) bool {
	stack := []string{p.startSymbol, endSymbol}
	return p.match(strings.Split(input, ""), stack)
}

func (p *PDA) match(input []string, stack []string) bool {
	if stack[0] == endSymbol && len(input) == 0 {
		return true
	}

	if stack[0] == emptySymbol {
		return p.match(input, stack[1:])
	}

	if len(input) != 0 && input[0] == stack[0] {
		return p.match(input[1:], stack[1:])
	}

	for _, trans := range p.transitions[stack[0]] {
		expand := append(trans.derivation, stack[1:]...)
		if p.match(input, expand) {
			return true
		}
	}

	return false
}

func (p *PDA) PrintTransitions() {
	for _, symbolTransitions := range p.transitions {
		for _, trans := range symbolTransitions {
			fmt.Println(trans.String())
		}
	}
}

func (t *transition) String() string {
	return fmt.Sprintf("%s -> %v", t.symbol, t.derivation)
}
