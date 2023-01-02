package solution

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/teacat/noire"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/v2/animate"
	"image/color"
	"log"
)

var (
	backgroundColor = color.RGBA{R: 77, G: 77, B: 77, A: 255}
	elfColor        = color.RGBA{R: 23, G: 139, B: 3, A: 255}
	fillColor       = noire.NewHexA("333337", 1)
	fontColor       = color.White
)

func Animate() {
	solution := New()
	input := solution.ReadInput()

	//	input := `....#..
	//..###.#
	//#...#.#
	//.#...##
	//#.###..
	//##.#.##
	//.#..#..`

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Advent of Code - 2022 Day 23")
	if err := ebiten.RunGame(NewGame(input)); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	*animate.AbstractGame

	originalGrove map[geometry.Point]bool

	grove          map[geometry.Point]bool
	round          int
	firstDirection int

	emptyTilesAfterRound10 int

	fillColorPulse *animate.ColorPulse
}

func NewGame(grove string) *Game {
	originalGrove := parseInput(grove)

	g := &Game{
		originalGrove:  originalGrove,
		fillColorPulse: animate.NewColorPulse(fillColor).SetRange(-0.05, 0.01).SetStep(0.001),
	}
	g.AbstractGame = animate.New(g)

	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.grove = g.originalGrove
	g.round = 1
	g.firstDirection = 0

	g.AbstractGame.Restart()
	g.fillColorPulse.Reset()
}

func (g *Game) Play() {
	firstHalf := make(map[geometry.Point][]Move)
	anyElfHasMoved := step(g.grove, g.firstDirection, firstHalf)
	g.grove = move(firstHalf)

	if g.round == 10 {
		g.emptyTilesAfterRound10 = sumEmptyTiles(geometry.MapToGrid(g.grove))
	}

	if !anyElfHasMoved {
		g.Mode = animate.ModeDone
	} else {
		g.firstDirection = (g.firstDirection + 1) % len(directionGroups)
		g.round += 1
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	rectangle := geometry.MapToGrid(g.grove)

	if g.Mode == animate.ModeTitle {
		animate.DrawText(screen, "Press [Enter] to Start! (Unstable Diffusion)", 8, 16, fontColor)
	} else if g.Mode == animate.ModePlay || g.Mode == animate.ModePause {
		animate.DrawText(screen, fmt.Sprintf("Rounds: %d", g.round), 8, 16, fontColor)
	} else if g.Mode == animate.ModeDone {
		c := g.fillColorPulse.Update()
		fc := animate.ToColor(c)

		vector.DrawFilledRect(screen,
			float32(g.BorderHorizontal), float32(g.BorderVertical),
			float32(g.TileSize*len(rectangle[0])), float32(g.TileSize*len(rectangle)),
			fc)

		animate.DrawText(screen, fmt.Sprintf("Part 1: %d, Part 2: %d", g.emptyTilesAfterRound10, g.round), 8, 16, fontColor)
	}
	_, height := ebiten.WindowSize()
	animate.DrawText(screen, fmt.Sprintf("[TPS: %d, Actual: %d]", ebiten.TPS(), int(ebiten.ActualTPS())), 8, height-8, fontColor)

	for y := 0; y < len(rectangle); y++ {
		for x := 0; x < len(rectangle[y]); x++ {
			if rectangle[y][x] == '.' {
				continue
			}
			vector.DrawFilledRect(
				screen,
				float32(x*g.TileSize+g.BorderHorizontal), float32(y*g.TileSize+g.BorderVertical),
				float32(g.TileSize), float32(g.TileSize),
				elfColor,
			)
		}
	}
}
