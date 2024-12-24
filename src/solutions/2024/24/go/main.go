package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"fmt"
	"github.com/samber/lo"
	"regexp"
	"sort"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	values, gates := parse(input)
	return calculateZValue(gates, values)
}

func part2(input string, _ ...interface{}) interface{} {
	values, gates := parse(input)
	//xyWires := lo.Keys(values)
	//sort.Slice(xyWires, func(i, j int) bool {
	//	return xyWires[i] > xyWires[j]
	//})
	//x := lo.Reduce(xyWires, func(agg int, wire string, index int) int {
	//	if wire[0] == 'x' {
	//		return (agg << 1) | values[wire]
	//	}
	//	return agg
	//}, 0)
	//y := lo.Reduce(xyWires, func(agg int, wire string, index int) int {
	//	if wire[0] == 'y' {
	//		return (agg << 1) | values[wire]
	//	}
	//	return agg
	//}, 0)
	//
	//z := calculateZValue(gates, values)

	inputBitCount := len(values) / 2

	flags := make(map[string]bool)

	var FAGate0s []Gate
	for _, g := range gates {
		if isDirect(g) && isGate("XOR")(g) {
			FAGate0s = append(FAGate0s, g)
		}
	}

	for _, g := range FAGate0s {
		isFirst := g.left == "x00" || g.right == "x00"
		if isFirst {
			if g.output != "z00" {
				flags[g.output] = true
			}
			continue
		} else if g.output == "z00" {
			flags[g.output] = true
		}

		if isOutput(g) {
			flags[g.output] = true
		}
	}

	var FAGate3s []Gate
	for _, g := range gates {
		if isGate("XOR")(g) && !isDirect(g) {
			FAGate3s = append(FAGate3s, g)
		}
	}

	for _, g := range FAGate3s {
		if !isOutput(g) {
			flags[g.output] = true
		}
	}

	var outputGates []Gate
	for _, g := range gates {
		if isOutput(g) {
			outputGates = append(outputGates, g)
		}
	}

	for _, g := range outputGates {
		isLast := g.output == fmt.Sprintf("z%02d", inputBitCount)
		if isLast {
			if g.operator != "OR" {
				flags[g.output] = true
			}
			continue
		} else if g.operator != "XOR" {
			flags[g.output] = true
		}
	}

	var checkNext []Gate
	for _, g := range FAGate0s {
		if _, flagged := flags[g.output]; flagged || g.output == "z00" {
			continue
		}

		matches := filterGates(FAGate3s, hasInput(g.output))
		if len(matches) == 0 {
			checkNext = append(checkNext, g)
			flags[g.output] = true
		}
	}

	for _, g := range checkNext {
		intendedResult := "z" + g.left[1:]
		matches := filterGates(FAGate3s, hasOutput(intendedResult))
		if len(matches) != 1 {
			return "TBD"
		}

		match := matches[0]
		toCheck := []string{match.left, match.right}
		orMatches := filterGates(lo.Values(gates), func(g Gate) bool {
			return g.operator == "OR" && (g.output == toCheck[0] || g.output == toCheck[1])
		})
		if len(orMatches) != 1 {
			return "TBD"
		}

		orMatchOutput := orMatches[0].output
		for _, output := range toCheck {
			if output != orMatchOutput {
				flags[output] = true
				break
			}
		}
	}

	if len(flags) != 8 {
		return "TBD"
	}

	flagsArr := lo.Keys(flags)
	sort.Strings(flagsArr)
	return strings.Join(flagsArr, ",")
}

func calculateZValue(gates map[string]Gate, values map[string]int) int {
	trees := lo.FilterMap(lo.Keys(gates), func(gate string, index int) (*Node, bool) {
		if gate[0] == 'z' {
			return expand(gates, values, gate), true
		}
		return nil, false
	})
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].output > trees[j].output
	})

	for _, root := range trees {
		root.value = postOrderTraversal(root)
	}

	bits := lo.Map(trees, func(node *Node, index int) int {
		return node.value
	})

	return calculateDecimal(bits)
}

func calculateDecimal(bits []int) int {
	value := 0
	for _, bit := range bits {
		value = (value << 1) | bit
	}
	return value
}

func parse(input string) (map[string]int, map[string]Gate) {
	chunks := strings.Split(input, "\n\n")
	values := make(map[string]int)
	for _, line := range strings.Split(chunks[0], "\n") {
		parts := strings.Split(line, ": ")
		values[parts[0]] = convert.StringToInt(parts[1])
	}
	gates := make(map[string]Gate)
	for _, line := range strings.Split(chunks[1], "\n") {
		match := regex.FindStringSubmatch(line)
		left, operator, right, gate := match[1], match[2], match[3], match[4]
		gates[gate] = Gate{left, operator, right, gate}
	}
	return values, gates
}

func postOrderTraversal(node *Node) int {
	if node == nil {
		return -1
	}

	if node.left == nil && node.right == nil {
		return node.value
	}

	left := postOrderTraversal(node.left)
	right := postOrderTraversal(node.right)

	switch node.operator {
	case "OR":
		return left | right
	case "AND":
		return left & right
	case "XOR":
		return left ^ right
	}

	return -1
}

func expand(gates map[string]Gate, values map[string]int, gate string) *Node {
	if gate == "" {
		return nil
	}

	g := gates[gate]

	node := &Node{
		value: -1,

		left:     expand(gates, values, g.left),
		operator: g.operator,
		right:    expand(gates, values, g.right),
		output:   gate,
	}

	if value, ok := values[gate]; ok {
		node.value = value
	}

	return node
}

var regex = regexp.MustCompile(`^(\w{3}) (OR|AND|XOR) (\w{3}) -> (\w{3})$`)

type Node struct {
	value int

	left     *Node
	operator string
	right    *Node
	output   string
}
type Gate struct {
	left     string
	operator string
	right    string
	output   string
}

func isDirect(g Gate) bool {
	return strings.HasPrefix(g.left, "x") || strings.HasPrefix(g.right, "x")
}

func isOutput(g Gate) bool {
	return strings.HasPrefix(g.output, "z")
}

func isGate(operator string) func(Gate) bool {
	return func(g Gate) bool {
		return g.operator == operator
	}
}

func hasOutput(output string) func(Gate) bool {
	return func(g Gate) bool {
		return g.output == output
	}
}

func hasInput(input string) func(Gate) bool {
	return func(g Gate) bool {
		return g.left == input || g.right == input
	}
}

func filterGates(gates []Gate, predicate func(Gate) bool) []Gate {
	var result []Gate
	for _, g := range gates {
		if predicate(g) {
			result = append(result, g)
		}
	}
	return result
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 24, Part1: part1, Part2: part2}
}
