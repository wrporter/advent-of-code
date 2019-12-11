package paint

import (
	"github.com/wrporter/advent-of-code-2019/day5/public/computer"
	"sync"
)

const NumDirections = 4

const (
	Up    = 0
	Left  = 1
	Down  = 2
	Right = 3
)

const (
	Black = 0
	White = 1
)

const (
	TurnRight = -1
	TurnLeft  = 1
)

type Position struct {
	Row int
	Col int
}

type Robot struct {
	program   *computer.Program
	direction int
	position  Position
}

func NewRobot(code []int) *Robot {
	return &Robot{
		computer.NewProgram(code),
		Up,
		Position{0, 0},
	}
}

func (r *Robot) Paint() int {
	var wg sync.WaitGroup
	panels := make(map[Position]int)
	cpu := computer.New()

	wg.Add(1)
	go cpu.ThreadProgram(&wg, r.program)

	wg.Add(1)
	go func() {
		for {
			r.program.Input <- r.camera(panels)

			if color, running := <-r.program.Output; running {
				panels[r.position] = color
			} else {
				break
			}

			if turn, running := <-r.program.Output; running {
				r.turn(turn)
			} else {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()

	return len(panels)
}

func (r *Robot) camera(panels map[Position]int) int {
	if color, ok := panels[r.position]; ok {
		return color
	}
	return Black
}

func (r *Robot) turn(turn int) {
	r.direction = wrap(r.direction+toTurnDirection(turn), NumDirections)
	switch r.direction {
	case Up:
		r.position.Row -= 1
	case Left:
		r.position.Col -= 1
	case Down:
		r.position.Row += 1
	case Right:
		r.position.Col += 1
	}
}

func toTurnDirection(turn int) int {
	if turn == 0 {
		return TurnLeft
	} else if turn == 1 {
		return TurnRight
	}
	return -999
}

func wrap(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
