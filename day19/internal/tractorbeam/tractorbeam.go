package tractorbeam

import (
	"github.com/wrporter/advent-of-code-2019/internal/common/intcode"
)

type Drone struct {
	code []int
}

func NewDrone(code []int) *Drone {
	return &Drone{code}
}

func (d *Drone) Scan(area int) int {
	affectedArea := 0

	for y := 0; y < area; y++ {
		for x := 0; x < area; x++ {
			if d.isInBeam(x, y) {
				affectedArea++
			}
		}
	}

	return affectedArea
}

func (d *Drone) Fit(ship int) int {
	size := ship - 1
	x, y := 0, ship

	for {
		for !d.isInBeam(x, y) {
			x++
		}

		if d.isInBeam(x, y) && // bottom left
			d.isInBeam(x+size, y) && // bottom right
			d.isInBeam(x, y-size) && // top left
			d.isInBeam(x+size, y-size) { // top right
			return x*10000 + (y - size)
		}
		y++
	}
}

func (d *Drone) isInBeam(x int, y int) bool {
	in, out := d.deploy()
	in <- x
	in <- y
	return <-out == 1
}

func (d *Drone) deploy() (chan int, chan int) {
	cpu := intcode.New()
	program := intcode.NewProgram(d.code)
	in := program.Input
	out := program.Output
	cpu.Run(program)
	return in, out
}
