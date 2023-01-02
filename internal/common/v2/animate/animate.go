package animate

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	DefaultTPS              = 60
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

	Mode             GameMode
	TPS              int
	TileSize         int
	BorderVertical   int
	BorderHorizontal int
}

func New(game Game) *AbstractGame {
	return &AbstractGame{
		Game:             game,
		Mode:             ModeTitle,
		TPS:              DefaultTPS,
		TileSize:         DefaultTileSize,
		BorderVertical:   DefaultBorderVertical,
		BorderHorizontal: DefaultBorderHorizontal,
	}
}

func (g *AbstractGame) Restart() {
	g.Mode = ModeTitle
	g.TPS = DefaultTPS
}

func (g *AbstractGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.TPS > 0 {
		g.TPS = g.TPS - 100
		if g.TPS <= 0 {
			g.TPS = 60
		}
		ebiten.SetTPS(g.TPS)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.TPS < 2000 {
		g.TPS = g.TPS + 100
		if g.TPS == 160 {
			g.TPS = 100
		}
		ebiten.SetTPS(g.TPS)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		g.Game.Restart()
		g.Mode = ModeTitle
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return errors.New("quit")
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
