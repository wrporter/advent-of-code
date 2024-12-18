package breakout

import (
	"aoc/src/lib/go/intcode"
	"fmt"
	"strings"
	"time"
)

const (
	GridHeight = 20
	GridWidth  = 44
)

type Tile int

const (
	Empty  Tile = 0
	Wall   Tile = 1
	Block  Tile = 2
	Paddle Tile = 3
	Ball   Tile = 4
)

const (
	JoystickNeutral   = 0
	JoystickTiltLeft  = -1
	JoystickTiltRight = 1
)

type Breakout struct {
	program *intcode.Program
}

type State struct {
	Score     int
	Grid      Grid
	NumBlocks int
}

type Grid [][]Tile

type Point struct {
	X int
	Y int
}

func New(code []int) *Breakout {
	return &Breakout{intcode.NewProgram(code)}
}

func (b *Breakout) InsertQuarters(quarters int) {
	b.program.Memory[0] = quarters
}

func (b *Breakout) Play() State {
	state := State{0, makeGrid(), 0}
	cpu := intcode.New()
	cpu.Run(b.program)
	var ball, paddle Point

	readyForInput := make(chan struct{})
	go func() {
		for range readyForInput {
			b.program.Input <- moveJoystick(ball, paddle)
			state.display()
		}
	}()

	for {
		if x, running := <-b.program.Output; running {
			y := <-b.program.Output

			if x == -1 && y == 0 {
				state.Score = <-b.program.Output
			} else {
				tile := Tile(<-b.program.Output)
				state.Grid[y][x] = tile

				if tile == Ball {
					ball = Point{x, y}
					readyForInput <- struct{}{}
				} else if tile == Paddle {
					paddle = Point{x, y}
				} else if tile == Block {
					state.NumBlocks++
				}
			}
		} else {
			break
		}
	}

	state.display()
	return state
}

func moveJoystick(ball Point, paddle Point) int {
	if paddle.X < ball.X {
		return JoystickTiltRight
	} else if paddle.X > ball.X {
		return JoystickTiltLeft
	}
	return JoystickNeutral
}

func (s *State) display() {
	out := &strings.Builder{}
	out.WriteString("\033[2J")
	out.WriteString("\033[H")
	out.WriteString(s.render())
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func (s *State) render() string {
	return fmt.Sprintf("\033[32mScore: %d\033[0m\n%s", s.Score, s.Grid.render())
}

func (g Grid) render() string {
	result := ""
	for _, row := range g {
		for _, tile := range row {
			result += tile.render()
		}
		result += "\n"
	}
	return result
}

func (t Tile) render() string {
	switch t {
	case Wall:
		return "\u001b[40m \u001b[0m"
	case Block:
		return "\u001b[41m \u001b[0m"
	case Paddle:
		return "‾"
	case Ball:
		return "\033[34mO\033[0m"
	default:
		return " "
	}
}

func makeGrid() Grid {
	grid := make(Grid, GridHeight)
	for y := range grid {
		row := make([]Tile, GridWidth)
		grid[y] = row
	}
	return grid
}
