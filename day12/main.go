package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day12/internal/solarsystem"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day12/input.txt")
	system := solarsystem.New(lines)
	//system := solarsystem.New([]string{
	//	"<x=-1, y=0, z=2>",
	//	"<x=2, y=-10, z=-7>",
	//	"<x=4, y=-8, z=8>",
	//	"<x=3, y=5, z=-1>",
	//})
	//system := solarsystem.New([]string{
	//	"<x=-8, y=-10, z=0>",
	//	"<x=5, y=5, z=10>",
	//	"<x=2, y=-7, z=3>",
	//	"<x=9, y=-8, z=-3>",
	//})
	system.StepForward(1000)
	fmt.Println(system.RenderParticles())
	fmt.Println(system.TotalEnery())
}
