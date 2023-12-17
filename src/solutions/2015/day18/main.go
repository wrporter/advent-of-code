package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"fmt"
	"strings"
	"time"
)

const (
	Alive = '#'
	Dead  = '.'
)

type GameOfLife struct {
	grid               [][]byte
	numAlive           int
	cornersAlwaysAlive bool
}

var directions = []geometry.Point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func New(lines []string, cornersAlwaysAlive bool) *GameOfLife {
	grid := make([][]byte, len(lines))
	numAlive := 0

	for y, line := range lines {
		row := make([]byte, len(line))
		for x, char := range line {
			if cornersAlwaysAlive && isCorner(x, y, row, grid) {
				row[x] = Alive
			} else {
				row[x] = byte(char)
			}

			if row[x] == Alive {
				numAlive++
			}
		}
		grid[y] = row
	}

	return &GameOfLife{grid, numAlive, cornersAlwaysAlive}
}

func isCorner(x int, y int, row []byte, grid [][]byte) bool {
	return (x == 0 && y == 0) ||
		(x == len(row)-1 && y == 0) ||
		(x == 0 && y == len(grid)-1) ||
		(x == len(row)-1 && y == len(grid)-1)
}

func (g *GameOfLife) Step(steps int) {
	for i := 0; i < steps; i++ {
		g.Next()
	}
}

func (g *GameOfLife) Next() {
	next := make([][]byte, len(g.grid))
	numAlive := 0

	for y, row := range g.grid {
		nextRow := make([]byte, len(row))
		for x, cell := range row {
			numAliveNeighbors := g.getNumAliveNeighbors(y, x)
			nextRow[x] = Dead

			if g.cornersAlwaysAlive && isCorner(x, y, nextRow, next) {
				nextRow[x] = Alive
			} else if (cell == Alive && (numAliveNeighbors == 2 || numAliveNeighbors == 3)) ||
				(cell == Dead && numAliveNeighbors == 3) {
				nextRow[x] = Alive
			}

			if nextRow[x] == Alive {
				numAlive++
			}
		}
		next[y] = nextRow
	}

	g.grid = next
	g.numAlive = numAlive
}

func (g *GameOfLife) CountAlive() int {
	return g.numAlive
}

func (g *GameOfLife) getNumAliveNeighbors(row int, col int) int {
	count := 0
	for _, direction := range directions {
		y, x := row+direction.Y, col+direction.X
		if y >= 0 &&
			y < len(g.grid) &&
			x >= 0 &&
			x < len(g.grid[y]) &&
			g.grid[y][x] == Alive {
			count++
		}
	}
	return count
}

func (g *GameOfLife) Display() {
	out := &strings.Builder{}
	out.WriteString("\033[2J\033[H")
	for _, row := range g.grid {
		out.Write(row)
		out.WriteByte('\n')
	}
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func main() {
	lines, _ := file.ReadFile("./2015/day18/input.txt")
	//lines := []string{
	//	".#.#.#",
	//	"...##.",
	//	"#....#",
	//	"..#...",
	//	"#.#..#",
	//	"####..",
	//}
	game := New(lines, false)
	game.Step(100)
	fmt.Println(game.CountAlive())

	game = New(lines, true)
	game.Step(100)
	fmt.Println(game.CountAlive())
}
