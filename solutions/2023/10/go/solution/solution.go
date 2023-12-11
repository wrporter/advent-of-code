package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	grid, start := parseInput(input)

	pipe := getStartPipe(grid, start)
	current := start.Copy()
	steps := 0

	for ok := true; ok; ok = !current.Equals(start) {
		current.Move(pipe.next)
		char := grid[current.Y][current.X]
		pipe = Pipes[IntoPipe{char, pipe.next}]
		steps++
	}

	return steps / 2
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	grid, start := parseInput(input)

	pipe := getStartPipe(grid, start)
	grid[start.Y][start.X] = pipe.char
	current := start.Copy()

	loop := map[geometry.Point]bool{*current.Copy(): true}

	for ok := true; ok; ok = !current.Equals(start) {
		current.Move(pipe.next)
		char := grid[current.Y][current.X]
		pipe = Pipes[IntoPipe{char, pipe.next}]
		loop[*current.Copy()] = true
	}

	// Use the ray-casting algorithm
	inside := 0
	for _, ray := range getTopAndLeftEdges(grid) {
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

func getTopAndLeftEdges(grid [][]string) []*geometry.Point {
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
	char string
	next geometry.Direction
	prev geometry.Direction
}

type IntoPipe struct {
	char string
	prev geometry.Direction
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

func parseInput(input string) ([][]string, *geometry.Point) {
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

func getStartPipe(grid [][]string, start *geometry.Point) Pipe {
	var startPipe Pipe

	for _, pipe := range Pipes {
		prev := start.Copy().Move(pipe.prev)
		next := start.Copy().Move(pipe.next)

		if prev.Y < 0 || prev.Y >= len(grid) ||
			prev.X < 0 || prev.X >= len(grid[prev.Y]) ||
			next.Y < 0 || next.Y >= len(grid) ||
			next.X < 0 || next.X >= len(grid[next.Y]) {
			continue
		}

		prevPipe, okPrev := Pipes[IntoPipe{grid[prev.Y][prev.X], pipe.prev}]
		nextPipe, okNext := Pipes[IntoPipe{grid[next.Y][next.X], pipe.next}]

		if okPrev && okNext &&
			prev.Copy().Move(prevPipe.prev).Equals(start) &&
			next.Copy().Move(nextPipe.prev).Equals(start) {
			startPipe = pipe
		}
	}

	return startPipe
}
