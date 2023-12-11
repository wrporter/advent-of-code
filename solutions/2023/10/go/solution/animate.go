package solution

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/wrporter/advent-of-code/internal/common/v2/animate"
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"image/color"
	"log"
	"math"
	"strings"
)

var (
	gray      = color.RGBA{R: 90, G: 82, B: 85, A: 255}
	green     = color.RGBA{R: 85, G: 158, B: 131, A: 255}
	red       = color.RGBA{R: 174, G: 90, B: 65, A: 255}
	blue      = color.RGBA{R: 27, G: 133, B: 184, A: 255}
	darkBlue  = color.RGBA{R: 12, G: 23, B: 39, A: 255}
	fontColor = color.White
)

type mode int

const (
	modeStart mode = iota
	modePart1Done
	modePart2Done
)

func Animate() {
	solution := New()
	input := solution.ReadInput()
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

type Game struct {
	*animate.AbstractGame

	grid      [][]string
	start     *geometry.Point
	startPipe Pipe
	current   *geometry.Point
	pipe      Pipe
	loop      map[geometry.Point]string

	inside  map[geometry.Point]bool
	outside map[geometry.Point]bool

	mode mode

	part1 int
	part2 int
}

func NewGame(input string) *Game {
	grid, start := parseInput(input)
	startPipe := getStartPipe(grid, start)
	grid[start.Y][start.X] = startPipe.char

	g := &Game{
		grid:      grid,
		start:     start,
		startPipe: startPipe,
	}

	g.AbstractGame = animate.New(g)
	g.AbstractGame.TileSize = 7

	//audio.NewPlayer()
	ebiten.SetWindowSize(g.AbstractGame.TileSize*len(grid[0])+g.BorderHorizontal*2, g.AbstractGame.TileSize*len(grid)+g.BorderVertical*2)

	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.pipe = g.startPipe
	g.current = g.start.Copy()
	g.loop = map[geometry.Point]string{*g.start.Copy(): g.pipe.char}
	g.mode = modeStart
	g.inside = make(map[geometry.Point]bool)
	g.outside = make(map[geometry.Point]bool)

	g.AbstractGame.Restart()
}

func (g *Game) Play() {
	for i := 0; i < g.Speed; i++ {
		g.step()

		if g.mode == modeStart && g.current.Equals(g.start) {
			g.part1 = len(g.loop) / 2
			g.mode = modePart1Done
		}
	}
}

func (g *Game) step() {
	if g.mode == modeStart {
		g.current.Move(g.pipe.next)
		char := g.grid[g.current.Y][g.current.X]
		g.pipe = Pipes[IntoPipe{char, g.pipe.next}]
		g.loop[*g.current.Copy()] = g.pipe.char
	} else if g.mode == modePart1Done {
		for _, ray := range getTopAndLeftEdges(g.grid) {
			intersections := 0

			for ray.Y < len(g.grid) && ray.X < len(g.grid[ray.Y]) {
				tile := g.grid[ray.Y][ray.X]

				if _, isOnLoop := g.loop[*ray]; isOnLoop && strings.Contains("|-JF", tile) {
					intersections += 1
				} else if !isOnLoop {
					if intersections%2 == 1 {
						g.inside[*ray.Copy()] = true
					} else {
						g.outside[*ray.Copy()] = true
					}
				}

				ray.Move(geometry.DownRight)
			}
		}

		g.Mode = animate.ModeDone
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(darkBlue)
	//rectangle := geometry.MapToGrid(g.loop)
	cellBorder := divCeil(g.TileSize, 4)

	if g.Mode == animate.ModeTitle {
		animate.DrawText(screen, "Press [Enter] to Start! (Pipe Maze)", 8, 16, fontColor)
	} else if g.Mode == animate.ModePlay || g.Mode == animate.ModePause {
		animate.DrawText(screen, fmt.Sprintf("Pipes: %d", len(g.loop)), 8, 16, fontColor)
	} else if g.Mode == animate.ModeDone {
		animate.DrawText(screen, fmt.Sprintf("Part 1: %d, Part 2: %d", g.part1, g.part2), 8, 16, fontColor)
	}
	_, windowHeight := ebiten.WindowSize()
	animate.DrawText(screen, fmt.Sprintf("[Speed: %d]", g.Speed), 8, windowHeight-8, fontColor)

	for y := range g.grid {
		for x := range g.grid[y] {
			p := *geometry.NewPoint(x, y)
			if _, ok := g.loop[p]; !ok {
				clr := gray
				if g.inside[p] {
					clr = green
				} else if g.outside[p] {
					clr = red
				}
				g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, 0, 0, clr)
			} else {
				switch g.loop[p] {
				case "-":
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize+cellBorder, 2*cellBorder, 0, blue)
				case "|":
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize, 0, 2*cellBorder, blue)
				case "F":
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, cellBorder, 0, blue) // middle right
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, 0, cellBorder, blue) // bottom
				case "7":
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize+cellBorder, cellBorder, 0, blue)            // middle left
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, 0, cellBorder, blue) // bottom
				case "L":
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize+cellBorder, cellBorder, 0, blue) // middle right
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize, 0, cellBorder, blue)            // top
				case "J":
					g.drawFilledRect(screen, x*g.TileSize, y*g.TileSize+cellBorder, cellBorder, 0, blue) // middle left
					g.drawFilledRect(screen, x*g.TileSize+cellBorder, y*g.TileSize, 0, cellBorder, blue) // top
				}
			}
		}
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
