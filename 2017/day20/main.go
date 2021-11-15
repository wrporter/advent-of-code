package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"sort"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 20
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>$`)

func part1(input []string) interface{} {
	particles := parseInput(input)

	sort.SliceStable(particles, func(i, j int) bool {
		a, b := particles[i], particles[j]

		if a.ManhattanAcceleration() == b.ManhattanAcceleration() {
			if a.ManhattanVelocity() == b.ManhattanVelocity() {
				return a.ManhattanDistance() < b.ManhattanDistance()
			}
			return a.ManhattanVelocity() < b.ManhattanVelocity()
		}
		return a.ManhattanAcceleration() < b.ManhattanAcceleration()
	})

	return particles[0].ID
}

func part2(input []string) interface{} {
	return 0
}

func parseInput(input []string) []*Particle {
	var particles []*Particle
	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		particle := &Particle{
			ID: i,
			Position: Point3D{
				X: conversion.StringToInt(match[1]),
				Y: conversion.StringToInt(match[2]),
				Z: conversion.StringToInt(match[3]),
			},
			Velocity: Point3D{
				X: conversion.StringToInt(match[4]),
				Y: conversion.StringToInt(match[5]),
				Z: conversion.StringToInt(match[6]),
			},
			Acceleration: Point3D{
				X: conversion.StringToInt(match[7]),
				Y: conversion.StringToInt(match[8]),
				Z: conversion.StringToInt(match[9]),
			},
		}
		particles = append(particles, particle)
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
