package main

import (
	"aoc/src/lib/go/file"
	nanofactory2 "aoc/src/solutions/2019/day14/lib/nanofactory"
	"fmt"
)

func main() {
	lines, _ := file.ReadFile("./2019/day14/input.txt")
	factory := nanofactory2.New()
	fmt.Println(factory.GetRequiredOre(lines, 1))
	fmt.Println(factory.OreToFuel(lines, 1000000000000))
}
