package animate

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	DefaultSpeed            = 1
	DefaultTileSize         = 4
	DefaultBorderVertical   = 32
	DefaultBorderHorizontal = 20
)

type Game interface {
	Restart()
	Play()
	Draw(screen *ebiten.Image)
}

type GameMode uint8

const (
	ModeTitle GameMode = iota
	ModePlay
	ModePause
	ModeDone
)

type AbstractGame struct {
	Game

	Mode GameMode
	// Speed represents how many operations should be performed every
	// update. This creates an artificial speed for when the game is drawn.
	Speed            int
	TileSize         int
	BorderVertical   int
	BorderHorizontal int

	Height int
	Width  int
}

func New(game Game) *AbstractGame {
	ebiten.SetRunnableOnUnfocused(false)
	return &AbstractGame{
		Game:             game,
		Mode:             ModeTitle,
		Speed:            DefaultSpeed,
		TileSize:         DefaultTileSize,
		BorderVertical:   DefaultBorderVertical,
		BorderHorizontal: DefaultBorderHorizontal,
	}
}

func (g *AbstractGame) Restart() {
	g.Mode = ModeTitle
	g.Speed = DefaultSpeed
}

func (g *AbstractGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		g.Game.Restart()
		g.Mode = ModeTitle
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return errors.New("quit")
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.Speed < 100 {
			g.Speed += 1
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.Speed > 1 {
			g.Speed -= 1
		}
	}

	switch g.Mode {
	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Mode = ModePlay
		}
	case ModePlay:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Mode = ModePause
		}
		g.Play()
	case ModePause:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Mode = ModePlay
		}
	case ModeDone:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Restart()
		}
	}

	return nil
}

func (g *AbstractGame) Draw(screen *ebiten.Image) {
	g.Game.Draw(screen)
}

func (g *AbstractGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
