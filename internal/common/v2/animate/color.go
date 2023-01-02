package animate

import (
	"github.com/teacat/noire"
	"image/color"
)

func ToColor(c noire.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a * 255),
	}
}
