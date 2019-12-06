package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day6/internal/orbit"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	orbits, _ := file.ReadFile("./day6/input.txt")
	//orbits := []string{"B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "COM)B", "K)L"}
	orbitMap := orbit.New()
	fmt.Println(orbitMap.Count(orbits))
}
