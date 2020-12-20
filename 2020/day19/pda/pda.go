package pda

import (
	"fmt"
	"strings"
)

type (
	PDA struct {
		startSymbol string
		startRule   string
		transitions []transition
	}

	transition struct {
		symbol     string
		derivation []string
	}
)

func NewPDA(startRule string) *PDA {
	return &PDA{
		startSymbol: "$",
		startRule:   startRule,
	}
}

func (p *PDA) AddRule(symbol string, derivation []string) {
	p.transitions = append(p.transitions, transition{
		symbol:     symbol,
		derivation: derivation,
	})
}

func (p *PDA) Match(input string) bool {
	stack := []string{p.startRule, p.startSymbol}
	return p.match(strings.Split(input, ""), stack, 0, nil, 0)

}

func (p *PDA) match(input []string, stack []string, depth int, rec [][]string, count int) bool {
	if depth > 200 {
		return false
	}

	if (stack[0] == p.startSymbol && len(input) == 0) ||
		(stack[0] == p.startSymbol && input[0] == "!") {
		return true
	}

	for _, trans := range p.transitions {
		if len(input) != 0 && input[0] == stack[0] {
			return p.match(input[1:], stack[1:], depth+1, rec, count)
		}

		// if the recursive-checking-list is not empty, let's analyze it to see if we have a recursion issue
		if len(rec) != 0 {
			var tmp []string
			for i := 1; i < len(rec); i++ {
				tmp = rec[i-1]
				if strings.Join(tmp, "") == strings.Join(rec[i], "") {
					count++
				}
			}
		}

		// if a rule recurses 50 or more times, let's skip it
		if count >= 50 {
			rec = nil
			count = 0
			continue
		}

		// transition on a variable
		if stack[0] == trans.symbol {
			rec = append(rec, trans.derivation)

			// pop symbol from stack and push
			if p.match(input, append(trans.derivation, stack[1:]...), depth+1, rec, count) {
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
