package solution

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"image/color"
	"log"
	"os"
)

var (
	backgroundColor = color.RGBA{R: 70, G: 70, B: 70, A: 50}
	elfColor        = color.RGBA{R: 0, G: 150, B: 50, A: 150}
	elfSize         = 4
	border          = 20
	FPSSteps        = []int{1, 2, 3, 4, 5, 6, 10, 20, 30, 60}
)

func Visual() {
	data, _ := os.ReadFile("solutions/2022/23/input.txt")
	input := string(data)

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

type GameMode uint8

const (
	ModeTitle GameMode = iota
	ModePlay
	ModePause
	ModeDone
)

type Game struct {
	originalGrove map[geometry.Point]bool

	grove          map[geometry.Point]bool
	round          int
	firstDirection int

	emptyTilesAfterRound10 int

	mode    GameMode
	tick    int
	fpsStep int
}

func NewGame(grove string) *Game {
	originalGrove := parseInput(grove)
	g := &Game{originalGrove: originalGrove}
	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.grove = g.originalGrove
	g.round = 1
	g.firstDirection = 0

	g.mode = ModeTitle
	g.tick = 0
	g.fpsStep = 0
}

func (g *Game) GetFPS() int {
	return 60 / FPSSteps[g.fpsStep]
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.mode = ModePlay
		}
	case ModePlay:
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.fpsStep < (len(FPSSteps)-1) {
			g.fpsStep += 1
		} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.fpsStep > 0 {
			g.fpsStep -= 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = ModePause
		} else if (g.tick % FPSSteps[g.fpsStep]) == 0 {
			firstHalf := make(map[geometry.Point][]Move)
			anyElfHasMoved := step(g.grove, g.firstDirection, firstHalf)
			g.grove = move(firstHalf)

			if g.round == 10 {
				g.emptyTilesAfterRound10 = sumEmptyTiles(geometry.MapToGrid(g.grove))
			}

			if !anyElfHasMoved {
				g.mode = ModeDone
			} else {
				g.firstDirection = (g.firstDirection + 1) % len(directionGroups)
				g.round += 1
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

	g.tick += 1

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	rectangle := geometry.MapToGrid(g.grove)

	if g.mode == ModeTitle {
		ebitenutil.DebugPrint(screen, "Advent of Code Day 23 - Elves disbursing to plant star fruit trees. Press Enter to Start!")
	} else if g.mode == ModePlay || g.mode == ModePause {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Round: %d [FPS: %d]", g.round, g.GetFPS()))
	} else if g.mode == ModeDone {
		emptyTiles := sumEmptyTiles(rectangle)
		ebitenutil.DebugPrint(screen, fmt.Sprintf(
			"Done! Rounds: %d - Empty Tiles: %d - Empty Tiles After Round 10: %d",
			g.round, emptyTiles, g.emptyTilesAfterRound10))
	}

	for y := 0; y < len(rectangle); y++ {
		for x := 0; x < len(rectangle[y]); x++ {
			if rectangle[y][x] == '.' {
				continue
			}
			ebitenutil.DrawRect(
				screen,
				float64(x*elfSize+border), float64(y*elfSize+border),
				float64(elfSize), float64(elfSize),
				elfColor,
			)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
