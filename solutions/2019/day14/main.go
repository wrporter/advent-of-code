package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	nanofactory2 "github.com/wrporter/advent-of-code/solutions/2019/day14/internal/nanofactory"
)

func main() {
	lines, _ := file.ReadFile("./2019/day14/input.txt")
	factory := nanofactory2.New()
	fmt.Println(factory.GetRequiredOre(lines, 1))
	fmt.Println(factory.OreToFuel(lines, 1000000000000))
}
