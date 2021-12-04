package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
)

var andOrRegex = regexp.MustCompile(`^([a-z]+) (AND|OR) ([a-z]+) -> ([a-z]+)$`)
var valueAndRegex = regexp.MustCompile(`^(\d+) (AND|OR) ([a-z]+) -> ([a-z]+)$`)
var shiftRegex = regexp.MustCompile(`^([a-z]+) (LSHIFT|RSHIFT) (\d+) -> ([a-z]+)$`)
var notRegex = regexp.MustCompile(`^(NOT) ([a-z]+) -> ([a-z]+)$`)
var valueRegex = regexp.MustCompile(`^(\d+) -> ([a-z]+)$`)
var directRegex = regexp.MustCompile(`^([a-z]+) -> ([a-z]+)$`)

type operator string

const (
	and      operator = "AND"
	valueAnd operator = "VALUEAND"
	or       operator = "OR"
	not      operator = "NOT"
	lshift   operator = "LSHIFT"
	rshift   operator = "RSHIFT"
	value    operator = "VALUE"
	direct   operator = "DIRECT"
)

type instruction struct {
	left        string
	operator    operator
	right       string
	shiftAmount uint16
	value       uint16
	result      string
	expression  string
}

func assemble(wireInstructions []string, wire string) uint16 {
	tree := make(map[string]instruction)
	for _, ins := range wireInstructions {
		if andOrRegex.MatchString(ins) {
			match := andOrRegex.FindStringSubmatch(ins)
			tree[match[4]] = instruction{
				expression: ins,
				left:       match[1],
				operator:   operator(match[2]),
				right:      match[3],
				result:     match[4],
			}
		} else if valueAndRegex.MatchString(ins) {
			match := valueAndRegex.FindStringSubmatch(ins)
			tree[match[4]] = instruction{
				expression: ins,
				value:      uint16(convert.StringToInt(match[1])),
				operator:   valueAnd,
				right:      match[3],
				result:     match[4],
			}
		} else if shiftRegex.MatchString(ins) {
			match := shiftRegex.FindStringSubmatch(ins)
			tree[match[4]] = instruction{
				expression:  ins,
				left:        match[1],
				operator:    operator(match[2]),
				shiftAmount: uint16(convert.StringToInt(match[3])),
				result:      match[4],
			}
		} else if notRegex.MatchString(ins) {
			match := notRegex.FindStringSubmatch(ins)
			tree[match[3]] = instruction{
				expression: ins,
				operator:   operator(match[1]),
				right:      match[2],
				result:     match[3],
			}
		} else if valueRegex.MatchString(ins) {
			match := valueRegex.FindStringSubmatch(ins)
			tree[match[2]] = instruction{
				expression: ins,
				operator:   value,
				value:      uint16(convert.StringToInt(match[1])),
				result:     match[2],
			}
		} else if directRegex.MatchString(ins) {
			match := directRegex.FindStringSubmatch(ins)
			tree[match[2]] = instruction{
				expression: ins,
				operator:   direct,
				left:       match[1],
				result:     match[2],
			}
		} else {
			fmt.Printf("Unknown expression: %s\n", ins)
		}
	}

	signals := make(map[string]uint16)
	return getSignal(tree, wire, signals)
}

func getSignal(tree map[string]instruction, wire string, signals map[string]uint16) uint16 {
	ins := tree[wire]
	var signal uint16

	//fmt.Printf("Get %s for expression [%s]\n", wire, ins.expression)
	if _, ok := signals[wire]; ok {
		return signals[wire]
	}

	if ins.operator == and {
		signal = getSignal(tree, ins.left, signals) & getSignal(tree, ins.right, signals)
	} else if ins.operator == valueAnd {
		signal = ins.value & getSignal(tree, ins.right, signals)
	} else if ins.operator == or {
		signal = getSignal(tree, ins.left, signals) | getSignal(tree, ins.right, signals)
	} else if ins.operator == not {
		signal = ^getSignal(tree, ins.right, signals)
	} else if ins.operator == lshift {
		signal = getSignal(tree, ins.left, signals) << ins.shiftAmount
	} else if ins.operator == rshift {
		signal = getSignal(tree, ins.left, signals) >> ins.shiftAmount
	} else if ins.operator == direct {
		signal = getSignal(tree, ins.left, signals)
	} else {
		signal = ins.value
	}

	signals[wire] = signal
	return signal
}

func main() {
	lines, _ := file.ReadFile("./2015/day7/input.txt")
	fmt.Println(assemble(lines, "a"))
	linesPart2, _ := file.ReadFile("./2015/day7/input-part2.txt")
	fmt.Println(assemble(linesPart2, "a"))
}
