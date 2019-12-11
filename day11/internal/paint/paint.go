package paint

import (
	"github.com/wrporter/advent-of-code-2019/day5/public/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/math"
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
	BlackChar = "."
	WhiteChar = "#"
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

func (r *Robot) Paint(startColor int) (int, [][]string) {
	var wg sync.WaitGroup
	panels := map[Position]int{Position{0, 0}: startColor}
	cpu := computer.New()
	topLeft := Position{0, 0}
	bottomRight := Position{0, 0}

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

			topLeft.Col = math.Min(topLeft.Col, r.position.Col)
			topLeft.Row = math.Max(topLeft.Row, r.position.Row)
			bottomRight.Col = math.Max(bottomRight.Col, r.position.Col)
			bottomRight.Row = math.Min(bottomRight.Row, r.position.Row)
		}
		wg.Done()
	}()

	wg.Wait()

	numPaintedPanels := len(panels)
	paint := toPaintedRegion(topLeft, bottomRight, panels)

	return numPaintedPanels, paint
}

func toPaintedRegion(topLeft Position, bottomRight Position, panels map[Position]int) [][]string {
	width := math.Abs(topLeft.Col) + math.Abs(bottomRight.Col) + 1
	height := math.Abs(topLeft.Row) + math.Abs(bottomRight.Row) + 1
	region := make([][]string, height)

	for h := 0; h < height; h++ {
		row := make([]string, width)
		region[h] = row

		for w := 0; w < width; w++ {
			color := panels[Position{h, w}]
			if color == White {
				row[w] = WhiteChar
			} else {
				row[w] = BlackChar
			}
		}
	}

	return region
}

func RenderPaintedRegion(region [][]string) string {
	result := ""

	for row := 0; row < len(region); row++ {
		for col := 0; col < len(region[row]); col++ {
			color := region[row][col]
			if color == BlackChar {
				result += " "
			} else {
				result += WhiteChar
			}
		}
		result += "\n"
	}

	return result
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
