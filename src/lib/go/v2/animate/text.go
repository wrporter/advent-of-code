package animate

import (
	"aoc/src/lib/go/v2/animate/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	nasalizationFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.Nasalization_otf)
	if err != nil {
		log.Fatal(err)
	}

	nasalizationFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	text.FaceWithLineHeight(nasalizationFont, 16)
}

func DrawText(screen *ebiten.Image, txt string, x, y int, c color.Color) {
	text.Draw(screen, txt, nasalizationFont, x, y, c)
}
