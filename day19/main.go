package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day19/internal/tractorbeam"
	"github.com/wrporter/advent-of-code-2019/internal/common/intcode"
	"github.com/wrporter/advent-of-code-2019/internal/common/timeit"
	"time"
)

func main() {
	code := intcode.ReadCode("./day19/input.txt")
	part1(code)
	part2(code)
}

func part1(code []int) {
	defer timeit.Track(time.Now(), "part 1")
	drone := tractorbeam.NewDrone(code)
	affectedArea := drone.Scan(50)
	fmt.Printf("Affected area: %d\n", affectedArea)
}

func part2(code []int) {
	defer timeit.Track(time.Now(), "part 2")
	drone := tractorbeam.NewDrone(code)
	affectedArea := drone.Fit(100)
	fmt.Printf("Affected area: %d\n", affectedArea)
}
