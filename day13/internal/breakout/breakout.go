package breakout

import (
	"github.com/wrporter/advent-of-code-2019/day13/public/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/math"
)

type TileID int

const (
	Empty  TileID = 0
	Wall   TileID = 1
	Block  TileID = 2
	Paddle TileID = 3
	Ball   TileID = 4
)

type Breakout struct {
	program *computer.Program
}

type Point struct {
	X int
	Y int
}

func New(code []int) *Breakout {
	return &Breakout{computer.NewProgram(code)}
}

func (b *Breakout) Play() int {
	grid := make(map[Point]TileID)
	maxX := 0
	maxY := 0
	x := 0
	y := 0
	tileID := Empty
	numBlocks := 0
	cpu := computer.New()
	cpu.Run(b.program)

	for {
		if nextX, running := <-b.program.Output; running {
			x = nextX
			y = <-b.program.Output
			tileID = TileID(<-b.program.Output)
			if tileID == Block {
				numBlocks++
			}

			grid[Point{x, y}] = tileID
			maxX = math.Max(maxX, x)
			maxY = math.Max(maxY, y)
		} else {
			break
		}
	}

	return numBlocks
}
