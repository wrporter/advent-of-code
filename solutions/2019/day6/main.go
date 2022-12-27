package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	orbit2 "github.com/wrporter/advent-of-code/solutions/2019/day6/internal/orbit"
)

func main() {
	orbits, _ := file.ReadFile("./2019/day6/input.txt")
	//orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	orbitFactory := orbit2.New()
	orbitMap := orbitFactory.Generate(orbits)
	fmt.Println(orbitFactory.Distance(orbitMap.Tree, orbit2.You, orbit2.Santa))
}
