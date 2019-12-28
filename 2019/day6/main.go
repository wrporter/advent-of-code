package main

import (
	"fmt"
	orbit2 "github.com/wrporter/advent-of-code-2019/2019/day6/internal/orbit"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	orbits, _ := file.ReadFile("./day6/input.txt")
	//orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	orbitFactory := orbit2.New()
	orbitMap := orbitFactory.Generate(orbits)
	fmt.Println(orbitFactory.Distance(orbitMap.Tree, orbit2.You, orbit2.Santa))
}
