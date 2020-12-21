// Context-Free Grammars (CGFs) describe the structure and design of a
// language. CFG rules can be expressed in the following form.
//
//     S: a S
//     S: !
//
// The language described by the grammar above is that of all strings of any
// number of `a`s. The second rule is what stops the grammar from producing
// `a`s, as it allows the terminal `S` to produce an empty string, denoted by
// `!`.
//
// A way to simulate whether or not certain strings follow the rules of a
// language is by constructing a Push-Down Automaton (PDA). A PDA is a machine
// that can process input and match it to the given grammar rules.
//
// This package implements an LL parser.
package pda

import (
	"strings"
)

const (
	endSymbol   = "$"
	emptySymbol = "!"
)

type (
	PDA struct {
		startSymbol string
		rules       map[string][]rule
	}

	rule struct {
		symbol  string
		product []string
	}
)

// Constructs a PDA instance.
func NewPDA(startSymbol string) *PDA {
	return &PDA{
		startSymbol: startSymbol,
		rules:       make(map[string][]rule),
	}
}

// Adds Backus-Naur Form rules of the form `S: a | b`, allowing for shorthand
// to CFG.
func (p *PDA) AddBNFRules(rules []string) {
	for _, r := range rules {
		parts := strings.Split(r, ": ")
		symbol := parts[0]

		for _, sequence := range strings.Split(parts[1], " | ") {
			product := strings.Split(sequence, " ")
			p.AddRule(symbol, product)
		}
	}
}

// Adds a CFG rule of the form `S: a S` where `S` is the symbol and
// []string{"a", "S"} is the product.
func (p *PDA) AddRule(symbol string, product []string) {
	p.rules[symbol] = append(p.rules[symbol], rule{
		symbol:  symbol,
		product: product,
	})
}

// Match returns whether the grammar defined in the PDA matches the given
// input.
func (p *PDA) Match(input string) bool {
	stack := []string{p.startSymbol, endSymbol}
	return p.match(strings.Split(input, ""), stack)
}

func (p *PDA) match(input []string, stack []string) bool {
	// The input has matched the grammar rules.
	if stack[0] == endSymbol && len(input) == 0 {
		return true
	}

	// A terminal empty symbol has been reached, remove it.
	if stack[0] == emptySymbol {
		return p.match(input, stack[1:])
	}

	// The terminal symbol on the stack matches the next terminal symbol in the
	// input, remove them and continue.
	if len(input) != 0 && input[0] == stack[0] {
		return p.match(input[1:], stack[1:])
	}

	// Transition on the next non-terminal symbol.
	for _, trans := range p.rules[stack[0]] {
		expand := append(trans.product, stack[1:]...)
		if p.match(input, expand) {
			return true
		}
	}

	// The input does not match the grammar.
	return false
}
