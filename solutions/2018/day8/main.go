package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 8
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	data, _ := convert.ToInts(strings.Fields(input[0]))
	root, _ := parseTree(data, 0)
	return sumMetadata(root)
}

func part2(input []string) interface{} {
	data, _ := convert.ToInts(strings.Fields(input[0]))
	root, _ := parseTree(data, 0)
	return getValue(root)
}

type node struct {
	children []*node
	metadata []int
}

func getValue(n *node) int {
	if len(n.children) == 0 {
		return ints.Sum(n.metadata)
	}

	value := 0
	for _, index := range n.metadata {
		i := index - 1
		if i >= 0 && i < len(n.children) {
			value += getValue(n.children[index-1])
		}
	}
	return value
}

func sumMetadata(n *node) int {
	sum := ints.Sum(n.metadata)
	for _, child := range n.children {
		sum += sumMetadata(child)
	}
	return sum
}

func parseTree(data []int, position int) (*node, int) {
	n := &node{}
	numNodes := data[position]
	numMetadataEntries := data[position+1]

	index := position + 2
	for i := 0; i < numNodes; i++ {
		child, nextIndex := parseTree(data, index)
		n.children = append(n.children, child)
		index = nextIndex
	}

	for i := 0; i < numMetadataEntries; i++ {
		n.metadata = append(n.metadata, data[index])
		index++
	}

	return n, index
}
