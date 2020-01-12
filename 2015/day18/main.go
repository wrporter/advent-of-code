package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"strings"
	"time"
)

const (
	Alive = '#'
	Dead  = '.'
)

type GameOfLife struct {
	grid     [][]byte
	numAlive int
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

func New(lines []string) *GameOfLife {
	grid := make([][]byte, len(lines))
	numAlive := 0

	for y, line := range lines {
		row := make([]byte, len(line))
		for x, char := range line {
			row[x] = byte(char)
			if row[x] == Alive {
				numAlive++
			}
		}
		grid[y] = row
	}

	return &GameOfLife{grid, numAlive}
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

			if (cell == Alive && (numAliveNeighbors == 2 || numAliveNeighbors == 3)) ||
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
	game := New(lines)
	game.Step(100)
	game.Display()
	fmt.Println(game.CountAlive())
}
