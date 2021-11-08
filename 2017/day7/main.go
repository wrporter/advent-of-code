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

	year, day := 2017, 7
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var nodeRegex = regexp.MustCompile(`^([a-z]+)\s\((\d+)\).*$`)

type Node struct {
	Name        string
	Weight      int
	Children    map[string]bool
	TotalWeight int
}

func part1(input []string) interface{} {
	_, root := constructTree(input)
	return root.Name
}

func constructTree(input []string) (tree map[string]Node, root Node) {
	tree = make(map[string]Node)
	for _, line := range input {
		match := nodeRegex.FindStringSubmatch(line)
		name := match[1]
		weight := conversion.StringToInt(match[2])
		node := Node{Name: name, Weight: weight}

		if strings.Contains(line, "->") {
			node.Children = make(map[string]bool)
			childrenStr := line[strings.Index(line, "-> ")+3:]
			children := strings.Split(childrenStr, ", ")
			for _, child := range children {
				node.Children[child] = true
			}
		}
		tree[name] = node
	}

	for _, node := range tree {
		isRoot := true

		for _, node2 := range tree {
			if node2.Children[node.Name] {
				isRoot = false
				break
			}
		}

		if isRoot {
			root = node
		}
	}

	return tree, root
}

func sumWeights(tree map[string]Node, node Node) int {
	sum := node.Weight

	for child := range tree[node.Name].Children {
		sum += sumWeights(tree, tree[child])
	}

	node.TotalWeight = sum
	tree[node.Name] = node

	return sum
}

func part2(input []string) interface{} {
	tree, root := constructTree(input)

	for child := range tree[root.Name].Children {
		sumWeights(tree, tree[child])
	}

	balance := make(map[int][]string)
	for child := range tree[root.Name].Children {
		weight := tree[child].TotalWeight
		balance[weight] = append(balance[weight], child)
	}

	desiredWeight := 0
	badWeight := 0
	for weight, children := range balance {
		if len(children) == 1 {
			badWeight = weight
			fmt.Println(children[0])
		} else {
			desiredWeight = weight
		}
	}

	fmt.Println(desiredWeight, badWeight)

	return (desiredWeight - badWeight) + badWeight
}
