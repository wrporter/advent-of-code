package pda

import (
	"fmt"
	"strings"
)

const endSymbol = "$"

type (
	PDA struct {
		startSymbol string
		transitions []transition
	}

	transition struct {
		symbol     string
		derivation []string
	}
)

func NewPDA(startSymbol string) *PDA {
	return &PDA{
		startSymbol: startSymbol,
	}
}

func (p *PDA) AddRules(rules []string) {
	for _, rule := range rules {
		parts := strings.Split(rule, ": ")
		symbol := parts[0]

		for _, sequence := range strings.Split(parts[1], " | ") {
			p.AddRule(symbol, strings.Split(sequence, " "))
		}
	}
}

func (p *PDA) AddRule(symbol string, derivation []string) {
	p.transitions = append(p.transitions, transition{
		symbol:     symbol,
		derivation: derivation,
	})
}

func (p *PDA) Match(input string) bool {
	stack := []string{p.startSymbol, endSymbol}
	return p.match(strings.Split(input, ""), stack)

}

func (p *PDA) match(input []string, stack []string) bool {
	if (stack[0] == endSymbol && len(input) == 0) ||
		(stack[0] == endSymbol && input[0] == "!") {
		return true
	}

	for _, trans := range p.transitions {
		if stack[0] == "!" {
			return p.match(input, stack[1:])
		}

		if len(input) != 0 && input[0] == stack[0] {
			return p.match(input[1:], stack[1:])
		}

		// transition on a variable
		if stack[0] == trans.symbol {
			// pop symbol from stack and push
			nextStack := append(trans.derivation, stack[1:]...)
			if p.match(input, nextStack) {
				return true
			}
		}
	}

	// no transition could be made, and the input didn't match the grammar
	return false
}

func (p *PDA) PrintTransitions() {
	for _, trans := range p.transitions {
		fmt.Println(trans.String())
	}
}

func (t *transition) String() string {
	return fmt.Sprintf("%s -> %v", t.symbol, t.derivation)
}
