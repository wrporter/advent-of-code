package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 12
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^([A-Z])(\d+)$`)

func part1(input []string) interface{} {
	direction := geometry.Right
	position := geometry.NewPoint(0, 0)

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		action := match[1]
		value := convert.StringToInt(match[2])

		switch action {
		case "N":
			position.Y += value
		case "S":
			position.Y -= value
		case "E":
			position.X += value
		case "W":
			position.X -= value
		case "L":
			direction = direction.Rotate(-value)
		case "R":
			direction = direction.Rotate(value)
		case "F":
			position = position.AddAmount(direction, value)
		}
	}

	return position.GetManhattanDistance()
}

func part2(input []string) interface{} {
	position := geometry.NewPoint(0, 0)
	waypoint := geometry.NewPoint(10, 1)

	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		action := match[1]
		value := convert.StringToInt(match[2])

		switch action {
		case "N":
			waypoint.Y += value
		case "S":
			waypoint.Y -= value
		case "E":
			waypoint.X += value
		case "W":
			waypoint.X -= value
		case "L":
			waypoint = waypoint.Rotate(-value)
		case "R":
			waypoint = waypoint.Rotate(value)
		case "F":
			position.X += waypoint.X * value
			position.Y += waypoint.Y * value
		}
	}

	return position.GetManhattanDistance()
}
