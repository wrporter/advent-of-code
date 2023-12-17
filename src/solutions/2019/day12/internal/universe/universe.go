package universe

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/ints"
	"fmt"
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

type Universe struct {
	Particles []Particle
}

func New(particlePositions []string) *Universe {
	var particles []Particle
	regex := regexp.MustCompile(`^<x=(-?\d+),\s*y=(-?\d+),\s*z=(-?\d+)>$`)

	for _, positionString := range particlePositions {
		matches := regex.FindStringSubmatch(positionString)
		particles = append(particles, Particle{
			Position: Coordinate{
				convert.StringToInt(matches[1]),
				convert.StringToInt(matches[2]),
				convert.StringToInt(matches[3]),
			},
			Velocity: Coordinate{},
		})
	}

	return &Universe{particles}
}

func (u *Universe) copy() *Universe {
	particles := make([]Particle, len(u.Particles))
	for i, particle := range u.Particles {
		particles[i] = particle
	}
	return &Universe{particles}
}

func (u *Universe) Simulate(numSteps int) []Particle {
	for i := 0; i < numSteps; i++ {
		u.step()
	}
	return u.Particles
}

func (u *Universe) step() {
	for p1 := range u.Particles {
		for p2 := range u.Particles {
			if p1 != p2 {
				u.Particles[p1].Velocity.X += modifier(u.Particles[p1].Position.X, u.Particles[p2].Position.X)
				u.Particles[p1].Velocity.Y += modifier(u.Particles[p1].Position.Y, u.Particles[p2].Position.Y)
				u.Particles[p1].Velocity.Z += modifier(u.Particles[p1].Position.Z, u.Particles[p2].Position.Z)
			}
		}
	}
	for i := range u.Particles {
		u.Particles[i].ApplyVelocity()
	}
}

func (u *Universe) StepsForFullCycle() int {
	steps := make([]int, 3)
	done := make([]bool, 3)
	initial := u.copy()

	for !all(done...) {
		u.Simulate(1)

		for axis := range steps {
			if !done[axis] {
				steps[axis]++
				if allAtInitialAxis(u, initial, axis) {
					done[axis] = true
				}
			}
		}
	}

	return ints.LCM(steps[0], steps[1], steps[2])
}

func allAtInitialAxis(u *Universe, initial *Universe, axis int) bool {
	allAtInitialAxis := true
	for p := range u.Particles {
		if !atInitialAxis(u.Particles[p], initial.Particles[p], axis) {
			allAtInitialAxis = false
		}
	}
	return allAtInitialAxis
}

func atInitialAxis(current Particle, initial Particle, axis int) bool {
	if axis == 0 {
		return current.Position.X == initial.Position.X && current.Velocity.X == initial.Velocity.X
	} else if axis == 1 {
		return current.Position.Y == initial.Position.Y && current.Velocity.Y == initial.Velocity.Y
	}
	return current.Position.Z == initial.Position.Z && current.Velocity.Z == initial.Velocity.Z
}

func (u *Universe) TotalEnergy() int {
	energy := 0
	for _, particle := range u.Particles {
		energy += particle.TotalEnergy()
	}
	return energy
}

func (u *Universe) RenderParticles() string {
	result := ""

	for _, particle := range u.Particles {
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

func (p *Particle) ApplyVelocity() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
	p.Position.Z += p.Velocity.Z
}

func (p Particle) PotentialEnergy() int {
	return ints.Abs(p.Position.X) + ints.Abs(p.Position.Y) + ints.Abs(p.Position.Z)
}

func (p Particle) KineticEnergy() int {
	return ints.Abs(p.Velocity.X) + ints.Abs(p.Velocity.Y) + ints.Abs(p.Velocity.Z)
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

func all(values ...bool) bool {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}
