package main

import (
	"fmt"
	universe2 "github.com/wrporter/advent-of-code/solutions/2019/day12/internal/universe"
)

func main() {
	//lines, _ := file.ReadFile("./2019/day12/input.txt")
	//u := universe.New(lines)
	u := universe2.New([]string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>",
	})
	//u := universe.New([]string{
	//	"<x=-8, y=-10, z=0>",
	//	"<x=5, y=5, z=10>",
	//	"<x=2, y=-7, z=3>",
	//	"<x=9, y=-8, z=-3>",
	//})
	//u.Simulate(10)
	//fmt.Println(u.RenderParticles())
	//fmt.Println(u.TotalEnergy())
	fmt.Println(u.StepsForFullCycle())
}
