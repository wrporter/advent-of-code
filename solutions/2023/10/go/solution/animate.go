package solution

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/teacat/noire"
	"github.com/wrporter/advent-of-code/internal/common/v2/animate"
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"image/color"
	"log"
	"math"
	"strings"
)

var (
	gray         = color.RGBA{R: 90, G: 82, B: 85, A: 255}
	green        = color.RGBA{R: 85, G: 158, B: 131, A: 255}
	insideColor  = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	outsideColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	red          = color.RGBA{R: 174, G: 90, B: 65, A: 255}
	yellow       = color.RGBA{R: 195, G: 203, B: 113, A: 255}
	blue         = color.RGBA{R: 27, G: 133, B: 184, A: 255}
	darkBlue     = color.RGBA{R: 36, G: 54, B: 83, A: 255}
	darkestBlue  = color.RGBA{R: 12, G: 23, B: 39, A: 255}
	fontColor    = color.White
	multiColor   = noire.NewHSL(0, 50, 60)

	speedDefault  = 40
	speedRayStart = 2
)

type mode int

const (
	modePart1 mode = iota
	modePart2
)

var totalPipes int

func Animate() {
	solution := New()
	input := solution.ReadInput()
	totalPipes = solution.Part1(input).(int)
	// Add input directly when compiling for WASM
	//	input := `.....
	//.S-7.
	//.|.|.
	//.L-J.
	//.....`

	ebiten.SetWindowTitle(fmt.Sprintf("Advent of Code - %d Day %d", solution.Year, solution.Day))
	if err := ebiten.RunGame(NewGame(input)); err != nil {
		log.Fatal(err)
	}
}

type loopPipe struct {
	char  string
	color color.Color
}

type Game struct {
	*animate.AbstractGame

	grid [][]string

	start     *geometry.Point
	startPipe Pipe

	current *geometry.Point
	pipe    Pipe
	loop    map[geometry.Point]loopPipe

	rayStarts     []*geometry.Point
	rayIndex      int
	rayEnd        *geometry.Point
	intersections int

	inside  map[geometry.Point]bool
	outside map[geometry.Point]bool

	mode mode

	part1 int
	part2 int
}

func NewGame(input string) *Game {
	grid, start := parseInput(input)
	topLeft := getTopAndLeftEdges(grid)
	startPipe := getStartPipe(grid, start)
	grid[start.Y][start.X] = startPipe.char

	g := &Game{
		grid:      grid,
		rayStarts: topLeft,
		start:     start,
		startPipe: startPipe,
	}

	g.AbstractGame = animate.New(g)
	g.AbstractGame.TileSize = 7

	//_, _ = audio.NewPlayer()
	ebiten.SetWindowSize(g.AbstractGame.TileSize*len(grid[0])+g.BorderHorizontal*2, g.AbstractGame.TileSize*len(grid)+g.BorderVertical*2)

	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.pipe = g.startPipe
	g.current = g.start.Copy()
	g.loop = map[geometry.Point]loopPipe{*g.start.Copy(): g.newLoopPipe(g.pipe.char)}
	g.rayIndex = 0
	g.rayEnd = g.rayStarts[g.rayIndex].Copy()
	g.mode = modePart1
	g.inside = make(map[geometry.Point]bool)
	g.outside = make(map[geometry.Point]bool)

	g.AbstractGame.Restart()
	g.Speed = speedDefault
}

func (g *Game) Play() {
	for i := 0; i < g.Speed; i++ {
		g.step()
	}
}

func (g *Game) step() {
	if g.Mode == animate.ModeDone {
		return
	}

	if g.mode == modePart1 {
		g.current.Move(g.pipe.next)
		char := g.grid[g.current.Y][g.current.X]
		g.pipe = Pipes[IntoPipe{char, g.pipe.next}]
		g.loop[*g.current.Copy()] = g.newLoopPipe(g.pipe.char)

		if g.current.Equals(g.start) {
			g.part1 = len(g.loop) / 2
			g.mode = modePart2
			g.Speed = speedRayStart
		}
	}

	if g.mode == modePart2 {
		tile := g.grid[g.rayEnd.Y][g.rayEnd.X]

		if _, isOnLoop := g.loop[*g.rayEnd]; isOnLoop && strings.Contains("|-JF", tile) {
			g.intersections += 1
		} else if !isOnLoop {
			if g.intersections%2 == 1 {
				g.inside[*g.rayEnd.Copy()] = true
			} else {
				g.outside[*g.rayEnd.Copy()] = true
			}
		}

		g.rayEnd.Move(geometry.DownRight)

		if g.rayEnd.Y >= len(g.grid) || g.rayEnd.X >= len(g.grid[g.rayEnd.Y]) {
			g.rayIndex++
			if g.rayIndex < len(g.rayStarts) {
				g.rayEnd = g.rayStarts[g.rayIndex].Copy()
			} else {
				g.part2 = len(g.inside)
				g.Mode = animate.ModeDone
			}

			if g.rayIndex > 5 && g.Speed < speedDefault {
				g.Speed += 1
			}
		}
	}
}

func (g *Game) newLoopPipe(char string) loopPipe {
	return loopPipe{char: char, color: animate.ToColor(multiColor.AdjustHue(float64(len(g.loop)) / float64(totalPipes) * 360))}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(darkBlue)
	cellBorder := divCeil(g.TileSize, 4)

	if g.Mode == animate.ModeTitle {
		animate.DrawText(screen, "Press [Enter] to Start! (Pipe Maze)", 8, 16, fontColor)
	} else if g.Mode == animate.ModePlay || g.Mode == animate.ModePause {
		animate.DrawText(screen, fmt.Sprintf("Pipes: %d", len(g.loop)), 8, 16, fontColor)
	} else if g.Mode == animate.ModeDone {
		animate.DrawText(screen, fmt.Sprintf("Part 1: %d, Part 2: %d", g.part1, g.part2), 8, 16, fontColor)
	}
	animate.DrawText(screen, fmt.Sprintf("[Speed: %d]", g.Speed), 8, 2*g.BorderVertical+len(g.grid)*g.TileSize-8, fontColor)

	for y := range g.grid {
		for x := range g.grid[y] {
			p := *geometry.NewPoint(x, y)
			if _, ok := g.loop[p]; !ok {
				clr := color.RGBA{R: 255, G: 255, B: 255, A: 255}
				if g.inside[p] {
					clr = insideColor
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize, 2*cellBorder, 2*cellBorder, darkestBlue)
				} else if g.outside[p] {
					clr = outsideColor
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize, 2*cellBorder, 2*cellBorder, darkestBlue)
				}

				g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, 0, 0, clr)
			} else {
				//pipeColor := animate.ToColor(multiColor.AdjustHue(float64(len(g.loop)) / float64(totalPipes) * 360))
				g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize, 2*cellBorder, 2*cellBorder, color.Black)

				switch pipe := g.loop[p]; pipe.char {
				case "-":
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize+cellBorder, 2*cellBorder, 0, pipe.color)
				case "|":
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize, 0, 2*cellBorder, pipe.color)
				case "F":
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, cellBorder, 0, pipe.color) // middle right
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, 0, cellBorder, pipe.color) // bottom
				case "7":
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize+cellBorder, cellBorder, 0, pipe.color)            // middle left
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, 0, cellBorder, pipe.color) // bottom
				case "L":
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, cellBorder, 0, pipe.color) // middle right
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize, 0, cellBorder, pipe.color)            // top
				case "J":
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize+cellBorder, cellBorder, 0, pipe.color) // middle left
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize, 0, cellBorder, pipe.color) // top
				}
			}
		}
	}

	if g.mode == modePart2 && g.Mode != animate.ModeDone {
		vector.StrokeLine(screen,
			float32(g.BorderHorizontal+g.rayStarts[g.rayIndex].X*g.TileSize+2*cellBorder),
			float32(g.BorderVertical+g.rayStarts[g.rayIndex].Y*g.TileSize+2*cellBorder),
			float32(g.BorderHorizontal+g.rayEnd.X*g.TileSize+2*cellBorder),
			float32(g.BorderVertical+g.rayEnd.Y*g.TileSize+2*cellBorder),
			float32(g.TileSize/2), color.RGBA{R: 255, G: 150, B: 40, A: 1}, false)
	}
}

func (g *Game) drawFilledRect(screen *ebiten.Image, x, y, width, height int, clr color.Color) {
	cellSize := g.TileSize / 2
	vector.DrawFilledRect(
		screen,
		float32(g.BorderHorizontal+x),
		float32(g.BorderVertical+y),
		float32(cellSize+width),
		float32(cellSize+height),
		clr,
		false,
	)
}

func divCeil(x, div int) int {
	return int(math.Ceil(float64(x) / float64(div)))
}
