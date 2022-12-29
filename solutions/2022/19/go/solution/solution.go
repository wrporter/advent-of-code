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
	top := 3
	if len(blueprints) < top {
		top = len(blueprints)
	}
	blueprints = blueprints[:top]
	product := 1

	for _, blueprint := range blueprints {
		maxGeodes := findMaxGeodes(blueprint, 32)
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

		next := Node{
			Resources: [4]int{
				current.Resources[0] + current.Robots[0],
				current.Resources[1] + current.Robots[1],
				current.Resources[2] + current.Robots[2],
				current.Resources[3] + current.Robots[3],
			},
			Robots: [4]int{
				current.Robots[0],
				current.Robots[1],
				current.Robots[2],
				current.Robots[3],
			},
			TimeLeft: current.TimeLeft,
		}

		states = append(states, next)

		if current.Resources[0] >= cost[3][0] && current.Resources[2] >= cost[3][2] {
			states = append(states, Node{
				Resources: [4]int{
					next.Resources[0] - cost[3][0],
					next.Resources[1],
					next.Resources[2] - cost[3][2],
					next.Resources[3],
				},
				Robots: [4]int{
					next.Robots[0],
					next.Robots[1],
					next.Robots[2],
					next.Robots[3] + 1,
				},
				TimeLeft: next.TimeLeft,
			})
		} else if current.Resources[0] >= cost[2][0] && current.Resources[1] >= cost[2][1] {
			states = append(states, Node{
				Resources: [4]int{
					next.Resources[0] - cost[2][0],
					next.Resources[1] - cost[2][1],
					next.Resources[2],
					next.Resources[3],
				},
				Robots: [4]int{
					next.Robots[0],
					next.Robots[1],
					next.Robots[2] + 1,
					next.Robots[3],
				},
				TimeLeft: next.TimeLeft,
			})
		} else {
			if current.Resources[0] >= cost[1][0] {
				states = append(states, Node{
					Resources: [4]int{
						next.Resources[0] - cost[1][0],
						next.Resources[1],
						next.Resources[2],
						next.Resources[3],
					},
					Robots: [4]int{
						next.Robots[0],
						next.Robots[1] + 1,
						next.Robots[2],
						next.Robots[3],
					},
					TimeLeft: next.TimeLeft,
				})
			}
			if current.Resources[0] >= cost[0][0] {
				states = append(states, Node{
					Resources: [4]int{
						next.Resources[0] - cost[0][0],
						next.Resources[1],
						next.Resources[2],
						next.Resources[3],
					},
					Robots: [4]int{
						next.Robots[0] + 1,
						next.Robots[1],
						next.Robots[2],
						next.Robots[3],
					},
					TimeLeft: next.TimeLeft,
				})
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
	maxRobots := [4]int{}

	for robot := range maxRobots {
		for mineral := range maxRobots {
			maxRobots[robot] = ints.Max(maxRobots[robot], blueprint.Cost[mineral][robot])
		}
	}

	//reflectCost := reflect.ValueOf(&blueprint.Cost).Elem()
	//reflectMaxRobots := reflect.ValueOf(&maxRobots).Elem()
	//for i := 0; i < reflectMaxRobots.NumField(); i++ {
	//	for j := 0; j < reflectMaxRobots.NumField(); j++ {
	//		cost := reflectCost.Field(i)
	//		max := reflectMaxRobots.Field(i)
	//		max.Set(reflect.ValueOf(ints.Max(int(reflect.ValueOf(max).Int()), int(reflect.ValueOf(cost).Int()))))
	//	}
	//}

	return maxRobots
}

func keyOf(node Node) string {
	return `${resources.ore},${resources.clay},${resources.obsidian},${resources.geode}-${robots.ore},${robots.clay},${robots.obsidian},${robots.geode}-${timeLeft}`
}

//func clone(node Node, resourceDiff map[string]int, robotDiff map[string]int) Node {
//	copy := Node{
//		Resources: {...resources},
//		Robots:    {...robots},
//		TimeLeft:  node.TimeLeft,
//	}
//
//	for _, mineral := range MINERALS {
//		copy.Resources[mineral] += resourceDiff[mineral]
//		copy.Robots[mineral] += robotDiff[mineral]
//	}
//
//	return copy
//}

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

//type Minerals struct {
//	Ore      int
//	Clay     int
//	Obsidian int
//	Geode    int
//}

type Node struct {
	Resources [4]int
	Robots    [4]int
	TimeLeft  int
}

//type Cost struct {
//	Ore      Minerals
//	Clay     Minerals
//	Obsidian Minerals
//	Geode    Minerals
//}

//var MINERALS = []string{"ore", "clay", "obsidian", "geode"}

type Blueprint struct {
	ID   int
	Cost [4][4]int
}
