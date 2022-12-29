package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"regexp"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	blueprints := parseInput(input)
	qualityLevel := 0

	for _, blueprint := range blueprints {
		maxGeodes := findMaxGeodes(blueprint, 24)
		qualityLevel += maxGeodes * blueprint.ID
	}

	return qualityLevel
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	blueprints := parseInput(input)
	product := 1

	for i := 0; i < len(blueprints) && i < 3; i++ {
		maxGeodes := findMaxGeodes(blueprints[i], 32)
		product *= maxGeodes
	}

	return product
}

func findMaxGeodes(blueprint Blueprint, maxTime int) int {
	maxCost := getMaxCost(blueprint)
	cost := blueprint.Cost
	maxGeodes := 0
	seen := make(map[Node]bool)
	states := []Node{{
		Resources: [4]int{0, 0, 0, 0},
		Robots:    [4]int{1, 0, 0, 0},
		TimeLeft:  maxTime,
	}}

	var node Node
	for len(states) > 0 {
		node, states = Pop(states)

		maxGeodes = ints.Max(maxGeodes, node.Resources[3])
		if node.TimeLeft == 0 {
			continue
		}

		current := Node{
			Resources: [4]int{
				getPotentialResource(node.Resources[0], node.Robots[0], maxCost[0], node.TimeLeft),
				getPotentialResource(node.Resources[1], node.Robots[1], cost[2][1], node.TimeLeft),
				getPotentialResource(node.Resources[2], node.Robots[2], cost[3][2], node.TimeLeft),
				node.Resources[3],
			},
			Robots: [4]int{
				ints.Min(node.Robots[0], maxCost[0]),
				ints.Min(node.Robots[1], cost[2][1]),
				node.Robots[2],
				ints.Min(node.Robots[3], cost[3][2]),
			},
			TimeLeft: node.TimeLeft - 1,
		}

		if seen[current] {
			continue
		}
		seen[current] = true

		next := clone(current, current.Robots, [4]int{})
		states = append(states, next)

		if current.Resources[0] >= cost[3][0] && current.Resources[2] >= cost[3][2] {
			states = append(states, clone(
				next,
				[4]int{-cost[3][0], 0, -cost[3][2], 0},
				[4]int{0, 0, 0, 1}),
			)
		} else if current.Resources[0] >= cost[2][0] && current.Resources[1] >= cost[2][1] {
			states = append(states, clone(
				next,
				[4]int{-cost[2][0], -cost[2][1], 0, 0},
				[4]int{0, 0, 1, 0}),
			)
		} else {
			if current.Resources[0] >= cost[1][0] {
				states = append(states, clone(
					next,
					[4]int{-cost[1][0], 0, 0, 0},
					[4]int{0, 1, 0, 0}),
				)
			}
			if current.Resources[0] >= cost[0][0] {
				states = append(states, clone(
					next,
					[4]int{-cost[0][0], 0, 0, 0},
					[4]int{1, 0, 0, 0}),
				)
			}
		}
	}

	return maxGeodes
}

func Pop(states []Node) (Node, []Node) {
	return states[len(states)-1], states[:len(states)-1]
}

func getPotentialResource(resources, robots, cost, timeLeft int) int {
	return ints.Min(resources, timeLeft*cost-robots*(timeLeft-1))
}

func getMaxCost(blueprint Blueprint) [4]int {
	maxCost := [4]int{}
	for robot := range maxCost {
		for mineral := range maxCost {
			maxCost[robot] = ints.Max(maxCost[robot], blueprint.Cost[mineral][robot])
		}
	}
	return maxCost
}

func clone(node Node, resourceDiff [4]int, robotDiff [4]int) Node {
	nodeClone := Node{
		Resources: [4]int{
			node.Resources[0] + resourceDiff[0],
			node.Resources[1] + resourceDiff[1],
			node.Resources[2] + resourceDiff[2],
			node.Resources[3] + resourceDiff[3],
		},
		Robots: [4]int{
			node.Robots[0] + robotDiff[0],
			node.Robots[1] + robotDiff[1],
			node.Robots[2] + robotDiff[2],
			node.Robots[3] + robotDiff[3],
		},
		TimeLeft: node.TimeLeft,
	}
	return nodeClone
}

func parseInput(input string) []Blueprint {
	regex := regexp.MustCompile(`\d+`)
	lines := strings.Split(input, "\n")
	blueprints := make([]Blueprint, len(lines))

	for i, line := range lines {
		nums, _ := convert.ToInts(regex.FindAllString(line, -1))
		blueprints[i] = Blueprint{
			ID: nums[0],
			Cost: [4][4]int{
				{nums[1], 0, 0, 0},
				{nums[2], 0, 0, 0},
				{nums[3], nums[4], 0, 0},
				{nums[5], 0, nums[6], 0},
			},
		}
	}

	return blueprints
}

type Node struct {
	Resources [4]int
	Robots    [4]int
	TimeLeft  int
}

type Blueprint struct {
	ID   int
	Cost [4][4]int
}
