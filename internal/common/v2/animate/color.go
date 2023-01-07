package animate

import (
	"github.com/teacat/noire"
	"image/color"
)

const (
	DefaultPulsePercent = 0
	DefaultPulseStep    = 0.004
	DefaultMin          = -0.15
	DefaultMax          = 0.05
)

type ColorPulse struct {
	originalPercent float64
	originalStep    float64

	Percent float64
	Step    float64
	Min     float64
	Max     float64
	Color   noire.Color
}

func NewColorPulse(c noire.Color) *ColorPulse {
	return &ColorPulse{
		originalPercent: DefaultPulsePercent,
		originalStep:    DefaultPulseStep,
		Percent:         DefaultPulsePercent,
		Step:            DefaultPulseStep,
		Min:             DefaultMin,
		Max:             DefaultMax,
		Color:           c,
	}
}

func (p *ColorPulse) Reset() {
	p.Percent = p.originalPercent
	p.Step = p.originalStep
}

func (p *ColorPulse) SetRange(min, max float64) *ColorPulse {
	p.Min = min
	p.Max = max
	return p
}

func (p *ColorPulse) SetStep(step float64) *ColorPulse {
	p.originalStep = step
	p.Step = step
	return p
}

func (p *ColorPulse) Update() noire.Color {
	p.Percent += p.Step
	if p.Percent <= p.Min || p.Percent >= p.Max {
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
