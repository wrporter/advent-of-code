package animate

import (
	"github.com/teacat/noire"
	"image/color"
)

const (
	DefaultPulsePercent = 0
	DefaultPulseStep    = 0.002
)

type ColorPulse struct {
	Percent float64
	Step    float64
	Color   noire.Color
}

func NewColorPulse(c noire.Color) *ColorPulse {
	return &ColorPulse{
		Percent: DefaultPulsePercent,
		Step:    DefaultPulseStep,
		Color:   c,
	}
}

func (p *ColorPulse) Reset() {
	p.Percent = DefaultPulsePercent
	p.Step = DefaultPulseStep
}

func (p *ColorPulse) Update() noire.Color {
	p.Percent += p.Step
	if p.Percent >= 0.05 || p.Percent <= -0.15 {
		p.Step = -p.Step
	}
	return p.Color.Lighten(p.Percent)
}

func ToColor(c noire.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a * 255),
	}
}
