package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day13/public/computer"
	"github.com/wrporter/advent-of-code-2019/day19/internal/tractorbeam"
)

func main() {
	code := computer.ReadCode("./day19/input.txt")
	drone := tractorbeam.NewDrone(code)
	affectedArea := drone.Scan(50)
	fmt.Printf("Affected area: %d\n", affectedArea)
}
