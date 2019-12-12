package solarsystem

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/math"
	"regexp"
)

type Coordinate struct {
	X int
	Y int
	Z int
}

type Particle struct {
	Position Coordinate
	Velocity Coordinate
}

type SolarSystem struct {
	Particles []Particle
}

func New(particlePositions []string) *SolarSystem {
	var particles []Particle
	regex := regexp.MustCompile(`^<x=(-?\d+),\s*y=(-?\d+),\s*z=(-?\d+)>$`)

	for _, positionString := range particlePositions {
		matches := regex.FindStringSubmatch(positionString)
		particles = append(particles, Particle{
			Position: Coordinate{
				conversion.StringToInt(matches[1]),
				conversion.StringToInt(matches[2]),
				conversion.StringToInt(matches[3]),
			},
			Velocity: Coordinate{},
		})
	}

	return &SolarSystem{particles}
}

func (s *SolarSystem) StepForward(numSteps int) []Particle {
	for i := 0; i < numSteps; i++ {
		particles := make([]Particle, len(s.Particles))
		for pIndex, particle := range s.Particles {
			particles[pIndex] = particle.ApplyGravity(s.Particles)
		}
		s.Particles = particles
	}

	return s.Particles
}

func (s *SolarSystem) TotalEnery() int {
	energy := 0
	for _, particle := range s.Particles {
		energy += particle.TotalEnergy()
	}
	return energy
}

func (s *SolarSystem) RenderParticles() string {
	result := ""

	for _, particle := range s.Particles {
		result += (&particle).Render()
		result += "\n"
	}

	return result
}

func (p Particle) Copy() Particle {
	return Particle{
		Position: Coordinate{
			p.Position.X,
			p.Position.Y,
			p.Position.Z,
		},
		Velocity: Coordinate{
			p.Velocity.X,
			p.Velocity.Y,
			p.Velocity.Z,
		},
	}
}

func (p Particle) ApplyGravity(particles []Particle) Particle {
	next := p.Copy()

	for _, p2 := range particles {
		if p != p2 {
			xModifier := modifier(p.Position.X, p2.Position.X)
			yModifier := modifier(p.Position.Y, p2.Position.Y)
			zModifier := modifier(p.Position.Z, p2.Position.Z)
			next.Velocity.X += xModifier
			next.Velocity.Y += yModifier
			next.Velocity.Z += zModifier
		}
	}
	next.ApplyVelocity()

	return next
}

func (p *Particle) ApplyVelocity() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
	p.Position.Z += p.Velocity.Z
}

func (p Particle) PotentialEnergy() int {
	return math.Abs(p.Position.X) + math.Abs(p.Position.Y) + math.Abs(p.Position.Z)
}

func (p Particle) KineticEnergy() int {
	return math.Abs(p.Velocity.X) + math.Abs(p.Velocity.Y) + math.Abs(p.Velocity.Z)
}

func (p Particle) TotalEnergy() int {
	return p.PotentialEnergy() * p.KineticEnergy()
}

func modifier(value1 int, value2 int) int {
	if value1 < value2 {
		return 1
	} else if value1 > value2 {
		return -1
	}
	return 0
}

func (p *Particle) Render() string {
	return fmt.Sprintf(
		"pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>",
		p.Position.X,
		p.Position.Y,
		p.Position.Z,
		p.Velocity.X,
		p.Velocity.Y,
		p.Velocity.Z,
	)
}
