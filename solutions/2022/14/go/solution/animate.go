package solution

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/teacat/noire"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/v2/animate"
	"github.com/wrporter/advent-of-code/internal/common/v2/contain"
	"image/color"
	"log"
)

var (
	backgroundColor = color.RGBA{R: 1, G: 31, B: 75, A: 255}
	rockColor       = color.RGBA{R: 77, G: 77, B: 77, A: 255}
	sourceColor     = color.RGBA{R: 234, G: 226, B: 214, A: 255}

	// sandColor pulled from https://icolorpalette.com/color/wet-sand
	sandColor  = noire.NewHex("ab8a5a")
	trailColor = animate.ToColor(sandColor.Darken(0.20))

	pulseStep = 0.002
)

func Animate() {
	solution := New()
	input := solution.ReadInput()

	//input := `498,4 -> 498,6 -> 496,6
	//503,4 -> 502,4 -> 502,9 -> 494,9`
	game := NewGame(input)

	width := 2*game.floor*game.TileSize + 2*game.BorderSize
	height := game.floor*game.TileSize + 2*game.BorderSize
	if width < 400 {
		width = 400
	}
	if height < 400 {
		height = 400
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Advent of Code - 2022 Day 14")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	*animate.AbstractGame

	originalScan map[geometry.Point]string

	scan   map[geometry.Point]string
	trail  map[geometry.Point]string
	source geometry.Point

	sand          int
	unit          geometry.Point
	hasComeToRest bool
	floor         int

	rockColor     uint8
	rockColorDiff int8

	sandColorPercent float64
	sandColorStep    float64
}

func NewGame(input string) *Game {
	source, scan, bottom := parseRockScan(input)
	floor := bottom + 2
	for x := source.X - floor; x <= source.X+floor; x++ {
		scan[geometry.NewPoint(x, floor)] = "#"
	}

	g := &Game{
		originalScan: scan,
		source:       source,
		floor:        floor,
	}
	g.AbstractGame = animate.New(g)

	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.scan = contain.CopyMap(g.originalScan)
	g.unit = g.source
	g.hasComeToRest = false
	g.sand = 0
	g.trail = make(map[geometry.Point]string)

	g.rockColor = 50
	g.rockColorDiff = 1

	g.sandColorPercent = 0
	g.sandColorStep = pulseStep

	g.AbstractGame.Restart()
	g.TPS = 2000
	ebiten.SetTPS(g.TPS)
}

func (g *Game) Play() {
	var next geometry.Point

	g.unit = g.source
	g.trail = make(map[geometry.Point]string)
	hasComeToRest := false

	for !hasComeToRest {
		if next = g.unit.Down(); g.shouldFallTo(next) {
			g.unit = next
		} else if next = g.unit.DownLeft(); g.shouldFallTo(next) {
			g.unit = next
		} else if next = g.unit.DownRight(); g.shouldFallTo(next) {
			g.unit = next
		} else {
			g.scan[g.unit] = "o"
			g.sand++
			hasComeToRest = true
		}

		g.trail[g.unit] = "-"

		if g.shouldExit(g.unit) {
			g.Mode = animate.ModeDone
		}
	}
}

func (g *Game) shouldExit(unit geometry.Point) bool {
	return unit == g.source
}

func (g *Game) shouldFallTo(unit geometry.Point) bool {
	_, exists := g.scan[unit]
	return !exists && unit.Y != g.floor
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	g.scan[g.source] = "S"
	gridMap := geometry.MapToGridV2(g.scan)
	geometry.Imprint(gridMap, g.trail)
	delete(g.scan, g.source)

	sc := sandColor

	if g.Mode == animate.ModeTitle {
		ebitenutil.DebugPrint(screen, "Falling Sand -- Press [Enter] to Start!")
	} else if g.Mode == animate.ModePlay || g.Mode == animate.ModePause {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Sand: %d", g.sand))
	} else if g.Mode == animate.ModeDone {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Sand: %d - Done!", g.sand))

		g.sandColorPercent += g.sandColorStep
		if g.sandColorPercent >= 0.10 || g.sandColorPercent <= -0.20 {
			g.sandColorStep = -g.sandColorStep
		}
		sc = sc.Lighten(g.sandColorPercent)
	}
	_, height := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[TPS: %d, Actual: %d]", ebiten.TPS(), int(ebiten.ActualTPS())), 0, height-16)

	var c color.Color
	for y := 0; y < len(gridMap.Grid); y++ {
		for x := 0; x < len(gridMap.Grid[y]); x++ {
			if gridMap.Grid[y][x] == '.' {
				continue
			}

			c = rockColor
			if gridMap.Grid[y][x] == 'S' {
				c = sourceColor
			} else if gridMap.Grid[y][x] == '-' {
				c = trailColor
			} else if gridMap.Grid[y][x] == 'o' {
				c = animate.ToColor(sc)
			}

			ebitenutil.DrawRect(
				screen,
				float64(x*g.TileSize+g.BorderSize), float64(y*g.TileSize+g.BorderSize),
				float64(g.TileSize), float64(g.TileSize),
				c,
			)
		}
	}
}

func (g *Game) getRockColor() color.RGBA {
	g.rockColor = uint8(int8(g.rockColor) + g.rockColorDiff)
	if g.rockColor == 160 {
		g.rockColorDiff = -1
	}
	if g.rockColor == 60 {
		g.rockColorDiff = 1
	}
	return color.RGBA{R: g.rockColor, G: g.rockColor, B: g.rockColor, A: 255}
}
