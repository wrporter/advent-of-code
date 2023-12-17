package solution

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/ints"
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
	var maxGeodesPossible []int
	for t := 0; t <= maxTime; t++ {
		maxGeodesPossible = append(maxGeodesPossible, (t-1)*t/2)
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
			resources[3]+(robots[3]*time)+maxGeodesPossible[time] <= maxGeodes {
			return
		}

		for time > 0 {
			if resources[0] >= cost[robot][0] && resources[1] >= cost[robot][1] &&
				resources[2] >= cost[robot][2] && resources[3] >= cost[robot][3] {
				for r := 0; r < 4; r++ {
					dfs(r, time-1, buildRobot(robots, robot), spendResources(resources, robots, cost, robot))
				}
				return
			}

			time -= 1
			resources = harvestResources(resources, robots)
		}

		maxGeodes = ints.Max(maxGeodes, resources[3])
	}

	for robot := 0; robot < 4; robot++ {
		dfs(robot, maxTime, [4]int{1, 0, 0, 0}, [4]int{0, 0, 0, 0})
	}

	return maxGeodes
}

func harvestResources(resources, robots [4]int) [4]int {
	return [4]int{
		resources[0] + robots[0],
		resources[1] + robots[1],
		resources[2] + robots[2],
		resources[3] + robots[3],
	}
}

func spendResources(resources, robots [4]int, cost [4][4]int, robot int) [4]int {
	return [4]int{
		resources[0] + robots[0] - cost[robot][0],
		resources[1] + robots[1] - cost[robot][1],
		resources[2] + robots[2] - cost[robot][2],
		resources[3] + robots[3] - cost[robot][3],
	}
}

func buildRobot(robots [4]int, robot int) [4]int {
	next := [4]int{robots[0], robots[1], robots[2], robots[3]}
	next[robot] += 1
	return next
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

type Blueprint struct {
	ID   int
	Cost [4][4]int
}
