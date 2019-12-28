package main

import (
	"fmt"
	nanofactory2 "github.com/wrporter/advent-of-code-2019/2019/day14/internal/nanofactory"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day14/input.txt")
	factory := nanofactory2.New()
	fmt.Println(factory.GetRequiredOre(lines, 1))
	fmt.Println(factory.OreToFuel(lines, 1000000000000))
}
