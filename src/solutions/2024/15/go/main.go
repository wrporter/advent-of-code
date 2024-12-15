package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/geometry"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	grid, directions := parse(input)

	var robot *geometry.Point
	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				robot = geometry.NewPoint(x, y)
				break
			}
		}
		if robot != nil {
			break
		}
	}

	//fmt.Printf("\u001B[H\u001B[2JInitial state:\n%s\n", runegrid.String(grid))
	//time.Sleep(20 * time.Millisecond)

	for _, d := range directions {
		direction := geometry.NewDirection(d)
		next := robot.Copy()
		var items []*geometry.Point

		for grid[next.Y][next.X] != '#' && grid[next.Y][next.X] != '.' {
			items = append(items, next.Copy())
			next.Move(direction)
		}

		if grid[next.Y][next.X] == '.' {
			for i := len(items) - 1; i >= 0; i-- {
				item := items[i]
				next = item.Copy().Move(direction)
				grid[item.Y][item.X], grid[next.Y][next.X] = grid[next.Y][next.X], grid[item.Y][item.X]
			}
			robot.Move(direction)
		}

		//fmt.Printf("\033[H\033[2JMove %s:\n%s\n", string(d), runegrid.String(grid))
		//time.Sleep(20 * time.Millisecond)
	}

	return sumGPS(grid)
}

func part2(input string, _ ...interface{}) interface{} {
	initialGrid, directions := parse(input)

	var robot *geometry.Point
	grid := make([][]rune, len(initialGrid))
	for y, row := range initialGrid {
		grid[y] = make([]rune, 2*len(initialGrid[y]))

		for x, cell := range row {
			switch cell {
			case '#', '.':
				grid[y][2*x] = cell
				grid[y][2*x+1] = cell
			case '@':
				grid[y][2*x] = '@'
				grid[y][2*x+1] = '.'
				robot = geometry.NewPoint(2*x, y)
			case 'O':
				grid[y][2*x] = '['
				grid[y][2*x+1] = ']'
			}
		}
	}

	//fmt.Printf("\u001B[H\u001B[2JInitial state:\n%s\n", runegrid.String(grid))
	//time.Sleep(20 * time.Millisecond)

	for _, d := range directions {
		direction := geometry.NewDirection(d)

		if canPush(grid, robot, direction) {
			push(grid, robot, direction)
			robot.Move(direction)
		}

		//fmt.Printf("\033[H\033[2JMove %s:\n%s\n", string(d), runegrid.String(grid))
		//time.Sleep(20 * time.Millisecond)
	}

	return sumGPS(grid)
}

func parse(input string) ([][]rune, string) {
	parts := strings.Split(input, "\n\n")
	grid := convert.ToRuneGrid(strings.Split(parts[0], "\n"))
	directions := strings.ReplaceAll(parts[1], "\n", "")
	return grid, directions
}

func sumGPS(grid [][]rune) interface{} {
	sum := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' || cell == '[' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func canPush(grid [][]rune, p *geometry.Point, direction geometry.Direction) bool {
	next := p.Copy().Move(direction)
	cell := grid[next.Y][next.X]

	if cell == '#' {
		return false
	} else if cell == '.' {
		return true
	} else if direction == geometry.Left || direction == geometry.Right {
		return canPush(grid, next, direction)
	} else if direction == geometry.Up || direction == geometry.Down {
		if cell == '[' {
			return canPush(grid, next, direction) &&
				canPush(grid, geometry.NewPoint(next.X+1, next.Y), direction)
		} else if cell == ']' {
			return canPush(grid, next, direction) &&
				canPush(grid, geometry.NewPoint(next.X-1, next.Y), direction)
		}
	}

	return false
}

func push(grid [][]rune, p *geometry.Point, direction geometry.Direction) {
	next := p.Copy().Move(direction)
	cell := grid[next.Y][next.X]

	if cell == '#' {
		return
	} else if cell == '.' {
		grid[p.Y][p.X], grid[next.Y][next.X] = grid[next.Y][next.X], grid[p.Y][p.X]
	} else if direction == geometry.Left || direction == geometry.Right {
		next2 := next.Copy().Move(direction)
		push(grid, next2, direction)
		grid[next2.Y][next2.X], grid[next.Y][next.X], grid[p.Y][p.X] = grid[next.Y][next.X], grid[p.Y][p.X], grid[next2.Y][next2.X]
	} else if direction == geometry.Up || direction == geometry.Down {
		push(grid, next, direction)
		if cell == '[' {
			push(grid, geometry.NewPoint(next.X+1, next.Y), direction)
		} else {
			push(grid, geometry.NewPoint(next.X-1, next.Y), direction)
		}
		grid[p.Y][p.X], grid[next.Y][next.X] = grid[next.Y][next.X], grid[p.Y][p.X]
	}
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 15, Part1: part1, Part2: part2}
}
