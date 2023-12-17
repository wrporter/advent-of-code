package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 7
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var nodeRegex = regexp.MustCompile(`^([a-z]+)\s\((\d+)\).*$`)

type Node struct {
	ID          string
	Parent      *Node
	Children    []*Node
	Weight      int
	TotalWeight int
}

type NodeDef struct {
	ID       string
	Weight   int
	Children []string
}

func part1(input []string) interface{} {
	root := constructTree(input)
	return root.ID
}

func constructTree(input []string) *Node {
	// Parse input into node definitions
	nodeDefs := make(map[string]*NodeDef)
	for _, line := range input {
		match := nodeRegex.FindStringSubmatch(line)
		id := match[1]
		weight := convert.StringToInt(match[2])
		nodeDef := &NodeDef{ID: id, Weight: weight}

		if strings.Contains(line, "->") {
			childrenStr := line[strings.Index(line, "-> ")+3:]
			children := strings.Split(childrenStr, ", ")
			for _, child := range children {
				nodeDef.Children = append(nodeDef.Children, child)
			}
		}

		nodeDefs[id] = nodeDef
	}

	// Populate nodes
	nodes := make(map[string]*Node)
	for _, nodeDef := range nodeDefs {
		nodes[nodeDef.ID] = &Node{
			ID:     nodeDef.ID,
			Weight: nodeDef.Weight,
		}
	}

	// Set parent and child relationships
	for _, node := range nodes {
		for _, childId := range nodeDefs[node.ID].Children {
			node.Children = append(node.Children, nodes[childId])
			nodes[childId].Parent = node
		}
	}

	// Find the root node, the only node without a parent
	var root *Node
	for _, node := range nodes {
		if node.Parent == nil {
			root = node
		}
	}

	return root
}

func setTotalWeights(node *Node) int {
	totalWeight := node.Weight

	for _, child := range node.Children {
		totalWeight += setTotalWeights(child)
	}

	node.TotalWeight = totalWeight

	return totalWeight
}

func findUnbalancedNode(node *Node) *Node {
	weights := make(map[int][]*Node)
	for _, child := range node.Children {
		weights[child.TotalWeight] = append(weights[child.TotalWeight], child)
	}

	for _, nodes := range weights {
		if len(nodes) == 1 {
			return findUnbalancedNode(nodes[0])
		}
	}

	return node
}

func calculateBalancedWeight(unbalancedNode *Node) int {
	weights := make(map[int][]*Node)
	for _, sibling := range unbalancedNode.Parent.Children {
		weights[sibling.TotalWeight] = append(weights[sibling.TotalWeight], sibling)
	}

	var desiredTotalWeight, badWeight, programWeight int
	for weight, nodes := range weights {
		if len(nodes) == 1 {
			badWeight = weight
			programWeight = nodes[0].Weight
		} else {
			desiredTotalWeight = weight
		}
	}
	desiredProgramWeight := (desiredTotalWeight - badWeight) + programWeight

	return desiredProgramWeight
}

func part2(input []string) interface{} {
	root := constructTree(input)
	setTotalWeights(root)
	unbalancedNode := findUnbalancedNode(root)
	balancedWeight := calculateBalancedWeight(unbalancedNode)

	return balancedWeight
}
