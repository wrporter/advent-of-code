package tractorbeam

import (
	"github.com/wrporter/advent-of-code-2019/day13/public/computer"
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
			in, out := d.deploy()
			in <- x
			in <- y
			if <-out == 1 {
				affectedArea++
			}
		}
	}

	return affectedArea
}

func (d *Drone) deploy() (chan int, chan int) {
	cpu := computer.New()
	program := computer.NewProgram(d.code)
	in := program.Input
	out := program.Output
	cpu.Run(program)
	return in, out
}
