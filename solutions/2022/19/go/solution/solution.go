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
	var sequence []int
	for t := 0; t <= maxTime; t++ {
		sequence = append(sequence, (t-1)*t/2)
	}
	maxCost := getMaxCost(blueprint)
	cost := blueprint.Cost
	maxGeodes := 0

	var dfs func(robot, time int, robots, resources [4]int)
	dfs = func(robot, time int, robots, resources [4]int) {
		if (robot == 0 && robots[0] >= maxCost[0]) ||
			(robot == 1 && robots[1] >= maxCost[1]) ||
			(robot == 2 && (robots[2] >= maxCost[2] || robots[1] == 0)) ||
			(robot == 3 && robots[2] == 0) ||
			resources[3]+(robots[3]*time)+sequence[time] <= maxGeodes {
			return
		}

		for time > 0 {
			if robot == 0 && resources[0] >= cost[0][0] {
				for r := 0; r < 4; r++ {
					dfs(r, time-1,
						[4]int{robots[0] + 1, robots[1], robots[2], robots[3]},
						[4]int{resources[0] + robots[0] - cost[0][0], resources[1] + robots[1], resources[2] + robots[2], resources[3] + robots[3]},
					)
				}
				return
			} else if robot == 1 && resources[0] >= cost[1][0] {
				for r := 0; r < 4; r++ {
					dfs(r, time-1,
						[4]int{robots[0], robots[1] + 1, robots[2], robots[3]},
						[4]int{resources[0] + robots[0] - cost[1][0], resources[1] + robots[1], resources[2] + robots[2], resources[3] + robots[3]},
					)
				}
				return
			} else if robot == 2 && resources[0] >= cost[2][0] && resources[1] >= cost[2][1] {
				for r := 0; r < 4; r++ {
					dfs(r, time-1,
						[4]int{robots[0], robots[1], robots[2] + 1, robots[3]},
						[4]int{resources[0] + robots[0] - cost[2][0], resources[1] + robots[1] - cost[2][1], resources[2] + robots[2], resources[3] + robots[3]},
					)
				}
				return
			} else if robot == 3 && resources[0] >= cost[3][0] && resources[2] >= cost[3][2] {
				for r := 0; r < 4; r++ {
					dfs(r, time-1,
						[4]int{robots[0], robots[1], robots[2], robots[3] + 1},
						[4]int{resources[0] + robots[0] - cost[3][0], resources[1] + robots[1], resources[2] + robots[2] - cost[3][2], resources[3] + robots[3]},
					)
				}
				return
			}

			time -= 1
			resources = gather(resources, robots)
		}

		maxGeodes = ints.Max(maxGeodes, resources[3])
	}

	for r := 0; r < 4; r++ {
		dfs(r, maxTime, [4]int{1, 0, 0, 0}, [4]int{0, 0, 0, 0})
	}

	return maxGeodes
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

func gather(resources, robots [4]int) [4]int {
	return [4]int{
		resources[0] + robots[0],
		resources[1] + robots[1],
		resources[2] + robots[2],
		resources[3] + robots[3],
	}
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

const (
	ORE      = 0
	CLAY     = 1
	OBSIDIAN = 2
	GEODE    = 3
)

type Blueprint struct {
	ID   int
	Cost [4][4]int
}
