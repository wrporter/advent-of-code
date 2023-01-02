package solution

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"image/color"
	"log"
)

var (
	backgroundColor = color.RGBA{R: 1, G: 31, B: 75, A: 255}
	sourceColor     = color.RGBA{R: 234, G: 226, B: 214, A: 255}
	trailColor      = color.RGBA{R: 70, G: 56, B: 35, A: 255}
	sandColor       = color.RGBA{R: 167, G: 134, B: 85, A: 255}
	tileSize        = 4
	border          = 20
)

func Visual() {
	solution := New()
	input := solution.ReadInput()

	//input := `498,4 -> 498,6 -> 496,6
	//503,4 -> 502,4 -> 502,9 -> 494,9`
	game := NewGame(input)

	ebiten.SetTPS(game.tps)

	ebiten.SetWindowSize(2*game.floor*tileSize+2*border, game.floor*tileSize+2*border)
	ebiten.SetWindowTitle("Advent of Code - 2022 Day 14")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type GameMode uint8

const (
	ModeTitle GameMode = iota
	ModePlay
	ModePause
	ModeDone
)

type Game struct {
	originalGrove map[geometry.Point]bool

	scan   map[geometry.Point]string
	trail  map[geometry.Point]string
	source geometry.Point

	sand          int
	unit          geometry.Point
	hasComeToRest bool
	floor         int

	mode         GameMode
	tick         int
	tps          int
	originalScan map[geometry.Point]string
	rgb          uint8
	rgbDiff      int8
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
	g.Restart()
	return g
}

func CopyMap(m map[geometry.Point]string) map[geometry.Point]string {
	c := make(map[geometry.Point]string)
	for k, v := range m {
		c[k] = v
	}
	return c
}

func (g *Game) shouldExit(unit geometry.Point) bool {
	return unit == g.source
}

func (g *Game) shouldFallTo(unit geometry.Point) bool {
	_, exists := g.scan[unit]
	return !exists && unit.Y != g.floor
}

func (g *Game) Restart() {
	g.scan = CopyMap(g.originalScan)
	g.unit = g.source
	g.hasComeToRest = false
	g.sand = 0
	g.trail = make(map[geometry.Point]string)

	g.mode = ModeTitle
	g.tick = 0
	g.tps = 1200
	g.rgb = 50
	g.rgbDiff = 1
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.tps > 0 {
		g.tps = g.tps - 100
		if g.tps <= 0 {
			g.tps = 60
		}
		ebiten.SetTPS(g.tps)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.tps < 2000 {
		g.tps = g.tps + 100
		if g.tps == 160 {
			g.tps = 100
		}
		ebiten.SetTPS(g.tps)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		g.Restart()
		g.mode = ModeTitle
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return errors.New("quit")
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.mode = ModePause
	}

	g.tick += 1

	switch g.mode {
	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.mode = ModePlay
		}
	case ModePlay:
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
				g.mode = ModeDone
			}
		}
	case ModePause:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = ModePlay
		}
	case ModeDone:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Restart()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	g.scan[g.source] = "S"
	gridMap := geometry.MapToGridV2(g.scan)
	geometry.Imprint(gridMap, g.trail)
	delete(g.scan, g.source)

	if g.mode == ModeTitle {
		ebitenutil.DebugPrint(screen, "Advent of Code Day 14 - Sand falling. Press [Enter] to Start!")
	} else if g.mode == ModePlay || g.mode == ModePause {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Sand: %d", g.sand))
	} else if g.mode == ModeDone {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Sand: %d - Done!", g.sand))
	}
	_, height := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[TPS: %d, Actual: %d]", ebiten.TPS(), int(ebiten.ActualTPS())), 0, height-16)

	rockColor := g.getRockColor()

	for y := 0; y < len(gridMap.Grid); y++ {
		for x := 0; x < len(gridMap.Grid[y]); x++ {
			if gridMap.Grid[y][x] == '.' {
				continue
			}

			c := rockColor
			if gridMap.Grid[y][x] == 'S' {
				c = sourceColor
			} else if gridMap.Grid[y][x] == '-' {
				c = trailColor
			} else if gridMap.Grid[y][x] == 'o' {
				c = sandColor
			}

			ebitenutil.DrawRect(
				screen,
				float64(x*tileSize+border), float64(y*tileSize+border),
				float64(tileSize), float64(tileSize),
				c,
			)
		}
	}
}

func (g *Game) getRockColor() color.RGBA {
	g.rgb = uint8(int8(g.rgb) + g.rgbDiff)
	if g.rgb == 160 {
		g.rgbDiff = -1
	}
	if g.rgb == 60 {
		g.rgbDiff = 1
	}
	return color.RGBA{R: g.rgb, G: g.rgb, B: g.rgb, A: 255}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
