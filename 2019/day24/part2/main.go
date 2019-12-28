package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"github.com/wrporter/advent-of-code-2019/internal/common/geometry"
	"github.com/wrporter/advent-of-code-2019/internal/common/timeit"
	"strings"
	"time"
)

const (
	Bug        = '#'
	EmptySpace = '.'
)

var Middle = geometry.NewPoint(2, 2)

const (
	Height = 5
	Width  = 5
)

type BugSensor struct {
	seen  map[string]int
	state map[int][][]byte
	low   int
	high  int
}

func New(lines []string) *BugSensor {
	layout := make([][]byte, len(lines))
	for y, line := range lines {
		row := make([]byte, len(line))
		for x, char := range line {
			row[x] = byte(char)
		}
		layout[y] = row
	}

	state := map[int][][]byte{
		-1: newGrid(Height, Width),
		0:  layout,
		1:  newGrid(Height, Width),
	}

	return &BugSensor{
		state: state,
		low:   -1,
		high:  1,
	}
}

func (b *BugSensor) CountBugs(steps int) int {
	for i := 0; i < steps; i++ {
		b.Step()
	}

	numBugs := 0

	for _, layout := range b.state {
		for _, row := range layout {
			for _, tile := range row {
				if tile == Bug {
					numBugs++
				}
			}
		}
	}

	return numBugs
}

func (b *BugSensor) Step() {
	b.state[b.low-1] = newGrid(Height, Width)
	b.state[b.high+1] = newGrid(Height, Width)
	nextState := make(map[int][][]byte)
	nextState[b.low-1] = b.state[b.low-1]
	nextState[b.high+1] = b.state[b.high+1]

	for depth := b.low; depth <= b.high; depth++ {
		nextLayout := stepLayout(b.state, depth)
		nextState[depth] = nextLayout
	}

	b.low--
	b.high++
	b.state = nextState
}

func stepLayout(state map[int][][]byte, depth int) [][]byte {
	layout := state[depth]
	nextLayout := make([][]byte, len(layout))
	for y, row := range layout {
		nextRow := make([]byte, len(row))

		for x := range row {
			nextRow[x] = EmptySpace
			point := geometry.NewPoint(x, y)
			if point == Middle {
				continue
			}
			if next(state, depth, point) {
				nextRow[x] = Bug
			}
		}

		nextLayout[y] = nextRow
	}
	return nextLayout
}

func next(state map[int][][]byte, depth int, point geometry.Point) bool {
	current := state[depth]
	below := state[depth+1]
	above := state[depth-1]
	numAdjacentBugs := 0

	for _, direction := range geometry.Directions {
		p := point.Add(direction)
		if p == Middle {
			if direction == geometry.Left {
				numAdjacentBugs += bugsInCol(below, Width-1)
			} else if direction == geometry.Right {
				numAdjacentBugs += bugsInCol(below, 0)
			} else if direction == geometry.Up {
				numAdjacentBugs += bugsInRow(below, Height-1)
			} else if direction == geometry.Down {
				numAdjacentBugs += bugsInRow(below, 0)
			}
		} else if (p.Y < 0 && above[Middle.Y-1][Middle.X] == Bug) ||
			(p.Y >= Height && above[Middle.Y+1][Middle.X] == Bug) ||
			(p.X < 0 && above[Middle.Y][Middle.X-1] == Bug) ||
			(p.X >= Width && above[Middle.Y][Middle.X+1] == Bug) ||
			(inBounds(p) && current[p.Y][p.X] == Bug) {
			numAdjacentBugs++
		}
	}

	isBug := current[point.Y][point.X] == Bug
	return (isBug && numAdjacentBugs == 1) ||
		(!isBug && (numAdjacentBugs == 1 || numAdjacentBugs == 2))
}

func inBounds(p geometry.Point) bool {
	return p.Y >= 0 && p.Y < Height && p.X >= 0 && p.X < Width
}

func bugsInCol(layout [][]byte, col int) int {
	numBugs := 0
	for _, row := range layout {
		if row[col] == Bug {
			numBugs++
		}
	}
	return numBugs
}

func bugsInRow(layout [][]byte, row int) int {
	numBugs := 0
	for _, tile := range layout[row] {
		if tile == Bug {
			numBugs++
		}
	}
	return numBugs
}

func newGrid(height, width int) [][]byte {
	grid := make([][]byte, height)
	for y := 0; y < height; y++ {
		row := make([]byte, width)
		for x := 0; x < width; x++ {
			row[x] = EmptySpace
		}
		grid[y] = row
	}
	return grid
}

func (b *BugSensor) Display() {
	out := &strings.Builder{}
	out.WriteString("\033[2J\033[H")
	for depth := b.low; depth <= b.high; depth++ {
		out.WriteString(fmt.Sprintf("Depth %d:\n", depth))
		for _, row := range b.state[depth] {
			out.Write(row)
			out.WriteByte('\n')
		}
	}
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func main() {
	lines, _ := file.ReadFile("./day24/input.txt")
	defer timeit.Track(time.Now(), "Time")
	sensor := New(lines)
	fmt.Printf("Bugs: %d\n", sensor.CountBugs(200))
	//sensor.Display()
}
