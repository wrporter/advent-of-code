package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 20
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>$`)

func part1(input []string) interface{} {
	particles := parseInput(input)
	var closestParticle *Particle
	for i := 0; i < 500; i++ {
		minDistance := ints.MaxInt
		for _, particle := range particles {
			particle.Step()
			if particle.ManhattanDistance() < minDistance {
				minDistance = particle.ManhattanDistance()
				closestParticle = particle
			}
		}
	}
	return closestParticle.ID

	//sort.SliceStable(particles, func(i, j int) bool {
	//	a, b := particles[i], particles[j]
	//
	//	if a.ManhattanAcceleration() == b.ManhattanAcceleration() {
	//		if a.ManhattanVelocity() == b.ManhattanVelocity() {
	//			return a.ManhattanDistance() < b.ManhattanDistance()
	//		}
	//		return a.ManhattanVelocity() < b.ManhattanVelocity()
	//	}
	//	return a.ManhattanAcceleration() < b.ManhattanAcceleration()
	//})
	//
	//return particles[0].ID
}

func part2(input []string) interface{} {
	particles := parseInput(input)

	for i := 0; i < 500; i++ {
		positions := make(map[Point3D][]int)
		for _, particle := range particles {
			positions[particle.Position] = append(positions[particle.Position], particle.ID)
		}

		for _, ids := range positions {
			if len(ids) > 1 {
				for _, id := range ids {
					delete(particles, id)
				}
			}
		}

		for _, particle := range particles {
			particle.Step()
		}
	}

	return len(particles)
}

func parseInput(input []string) map[int]*Particle {
	particles := make(map[int]*Particle)
	for id, line := range input {
		match := regex.FindStringSubmatch(line)
		particle := &Particle{
			ID: id,
			Position: Point3D{
				X: convert.StringToInt(match[1]),
				Y: convert.StringToInt(match[2]),
				Z: convert.StringToInt(match[3]),
			},
			Velocity: Point3D{
				X: convert.StringToInt(match[4]),
				Y: convert.StringToInt(match[5]),
				Z: convert.StringToInt(match[6]),
			},
			Acceleration: Point3D{
				X: convert.StringToInt(match[7]),
				Y: convert.StringToInt(match[8]),
				Z: convert.StringToInt(match[9]),
			},
		}
		particles[id] = particle
	}
	return particles
}

type (
	Particle struct {
		ID           int
		Position     Point3D
		Velocity     Point3D
		Acceleration Point3D
	}

	Point3D struct {
		X int
		Y int
		Z int
	}
)

func (p *Particle) Collides(p2 *Particle) bool {
	return p.Position.X == p2.Position.X &&
		p.Position.Y == p2.Position.Y &&
		p.Position.Z == p2.Position.Z
}

func (p *Particle) Step() {
	p.Velocity.X += p.Acceleration.X
	p.Velocity.Y += p.Acceleration.Y
	p.Velocity.Z += p.Acceleration.Z

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
	p.Position.Z += p.Velocity.Z
}

func (p *Particle) ManhattanDistance() int {
	return ints.Abs(p.Position.X) + ints.Abs(p.Position.Y) + ints.Abs(p.Position.Z)
}

func (p *Particle) ManhattanVelocity() int {
	return ints.Abs(p.Velocity.X) + ints.Abs(p.Velocity.Y) + ints.Abs(p.Velocity.Z)
}

func (p *Particle) ManhattanAcceleration() int {
	return ints.Abs(p.Acceleration.X) + ints.Abs(p.Acceleration.Y) + ints.Abs(p.Acceleration.Z)
}
