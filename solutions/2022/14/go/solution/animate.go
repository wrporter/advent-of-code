package solution

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/teacat/noire"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/v2/animate"
	"github.com/wrporter/advent-of-code/internal/common/v2/contain"
	"image/color"
	"log"
)

var (
	backgroundColor = color.RGBA{R: 1, G: 31, B: 75, A: 255}
	fontColor       = color.White
	rockColor       = color.RGBA{R: 77, G: 77, B: 77, A: 255}

	// sandColor pulled from https://icolorpalette.com/color/wet-sand
	sandColor   = noire.NewHex("ab8a5a")
	trailColor  = animate.ToColor(sandColor.Darken(0.20))
	sourceColor = animate.ToColor(sandColor.Lighten(0.40))
)

func Animate() {
	solution := New()
	input := solution.ReadInput()

	//input := `498,4 -> 498,6 -> 496,6
	//503,4 -> 502,4 -> 502,9 -> 494,9`
	game := NewGame(input)

	width := 2*game.floor*game.TileSize + 2*game.BorderHorizontal
	height := game.floor*game.TileSize + 2*game.BorderVertical
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

	unit          geometry.Point
	hasComeToRest bool
	bottom        int
	floor         int

	sand          int
	sandUntilVoid int

	rockColor     uint8
	rockColorDiff int8

	sandColorPulse *animate.ColorPulse
}

func NewGame(input string) *Game {
	source, scan, bottom := parseRockScan(input)
	floor := bottom + 2
	for x := source.X - floor; x <= source.X+floor; x++ {
		scan[geometry.NewPoint(x, floor)] = "#"
	}

	g := &Game{
		originalScan:   scan,
		source:         source,
		bottom:         bottom,
		floor:          floor,
		sandColorPulse: animate.NewColorPulse(sandColor),
	}
	g.AbstractGame = animate.New(g)

	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.scan = contain.CopyMap(g.originalScan)
	g.unit = g.source
	g.hasComeToRest = false
	g.trail = make(map[geometry.Point]string)

	g.sand = 0
	g.sandUntilVoid = 0

	g.rockColor = 50
	g.rockColorDiff = 1

	g.sandColorPulse.Reset()

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

		if g.unit.Y > g.bottom && g.sandUntilVoid == 0 {
			g.sandUntilVoid = g.sand
		}

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
		animate.DrawText(screen, "Press [Enter] to Start! (Regolith Reservoir)", 8, 16, fontColor)
	} else if g.Mode == animate.ModePlay || g.Mode == animate.ModePause {
		animate.DrawText(screen, fmt.Sprintf("Sand: %d", g.sand), 8, 16, fontColor)
	} else if g.Mode == animate.ModeDone {
		animate.DrawText(screen, fmt.Sprintf("Part 1: %d, Part 2: %d", g.sandUntilVoid, g.sand), 8, 16, fontColor)

		sc = g.sandColorPulse.Update()
	}
	_, height := ebiten.WindowSize()
	animate.DrawText(screen, fmt.Sprintf("[TPS: %d, Actual: %d]", ebiten.TPS(), int(ebiten.ActualTPS())), 8, height-8, fontColor)

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

			vector.DrawFilledRect(
				screen,
				float32(x*g.TileSize+g.BorderHorizontal), float32(y*g.TileSize+g.BorderVertical),
				float32(g.TileSize), float32(g.TileSize),
				c,
			)
		}
	}
}
