package solution

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid, start := ParseInput(input)

	pipe := GetStartPipe(grid, start)
	current := start.Copy()
	steps := 0

	for ok := true; ok; ok = !current.Equals(start) {
		current.Move(pipe.Next)
		char := grid[current.Y][current.X]
		pipe = Pipes[IntoPipe{char, pipe.Next}]
		steps++
	}

	return steps / 2
}

func part2(input string, _ ...interface{}) interface{} {
	grid, start := ParseInput(input)

	pipe := GetStartPipe(grid, start)
	grid[start.Y][start.X] = pipe.Char
	current := start.Copy()

	loop := map[geometry.Point]bool{*current.Copy(): true}

	for ok := true; ok; ok = !current.Equals(start) {
		current.Move(pipe.Next)
		char := grid[current.Y][current.X]
		pipe = Pipes[IntoPipe{char, pipe.Next}]
		loop[*current.Copy()] = true
	}

	// Use the ray-casting algorithm
	inside := 0
	for _, ray := range GetTopAndLeftEdges(grid) {
		intersections := 0

		for ray.Y < len(grid) && ray.X < len(grid[ray.Y]) {
			tile := grid[ray.Y][ray.X]

			// Exclude L and 7 because we are scanning from top-left to bottom-right and we pass
			// right through the inside of the loop to the outside.
			if loop[*ray] && strings.Contains("|-JF", tile) {
				intersections += 1
			} else if !loop[*ray] && intersections%2 == 1 {
				inside += 1
			}

			ray.Move(geometry.DownRight)
		}
	}
	return inside
}

func GetTopAndLeftEdges(grid [][]string) []*geometry.Point {
	topLeft := make([]*geometry.Point, len(grid)+len(grid[0])-1)
	i := 0
	for x := 0; x < len(grid[0]); x, i = x+1, i+1 {
		topLeft[i] = geometry.NewPoint(x, 0)
	}
	for y := 1; y < len(grid); y, i = y+1, i+1 {
		topLeft[i] = geometry.NewPoint(0, y)
	}
	return topLeft
}

type Pipe struct {
	Char string
	Next geometry.Direction
	Prev geometry.Direction
}

type IntoPipe struct {
	Char string
	Prev geometry.Direction
}

var Pipes = map[IntoPipe]Pipe{
	{"|", geometry.Down}:  {"|", geometry.Down, geometry.Up},
	{"|", geometry.Up}:    {"|", geometry.Up, geometry.Down},
	{"-", geometry.Right}: {"-", geometry.Right, geometry.Left},
	{"-", geometry.Left}:  {"-", geometry.Left, geometry.Right},
	{"L", geometry.Down}:  {"L", geometry.Right, geometry.Up},
	{"L", geometry.Left}:  {"L", geometry.Up, geometry.Right},
	{"J", geometry.Down}:  {"J", geometry.Left, geometry.Up},
	{"J", geometry.Right}: {"J", geometry.Up, geometry.Left},
	{"7", geometry.Right}: {"7", geometry.Down, geometry.Left},
	{"7", geometry.Up}:    {"7", geometry.Left, geometry.Down},
	{"F", geometry.Left}:  {"F", geometry.Down, geometry.Right},
	{"F", geometry.Up}:    {"F", geometry.Right, geometry.Down},
}

func ParseInput(input string) ([][]string, *geometry.Point) {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	var start *geometry.Point

	for y, row := range lines {
		grid[y] = strings.Split(row, "")
		for x, value := range grid[y] {
			if value == "S" {
				start = geometry.NewPoint(x, y)
			}
		}
	}

	return grid, start
}

func GetStartPipe(grid [][]string, start *geometry.Point) Pipe {
	var startPipe Pipe

	for _, pipe := range Pipes {
		prev := start.Copy().Move(pipe.Prev)
		next := start.Copy().Move(pipe.Next)

		if prev.Y < 0 || prev.Y >= len(grid) ||
			prev.X < 0 || prev.X >= len(grid[prev.Y]) ||
			next.Y < 0 || next.Y >= len(grid) ||
			next.X < 0 || next.X >= len(grid[next.Y]) {
			continue
		}

		prevPipe, okPrev := Pipes[IntoPipe{grid[prev.Y][prev.X], pipe.Prev}]
		nextPipe, okNext := Pipes[IntoPipe{grid[next.Y][next.X], pipe.Next}]

		if okPrev && okNext &&
			prev.Copy().Move(prevPipe.Prev).Equals(start) &&
			next.Copy().Move(nextPipe.Prev).Equals(start) {
			startPipe = pipe
		}
	}

	return startPipe
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 10, Part1: part1, Part2: part2}
}
