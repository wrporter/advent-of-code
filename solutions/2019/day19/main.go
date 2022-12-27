package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/intcode"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	tractorbeam2 "github.com/wrporter/advent-of-code/solutions/2019/day19/internal/tractorbeam"
	"time"
)

func main() {
	code := intcode.ReadCode("./day19/input.txt")
	part1(code)
	part2(code)
}

func part1(code []int) {
	defer timeit.Track(time.Now(), "part 1")
	drone := tractorbeam2.NewDrone(code)
	affectedArea := drone.Scan(50)
	fmt.Printf("Affected area: %d\n", affectedArea)
}

func part2(code []int) {
	defer timeit.Track(time.Now(), "part 2")
	drone := tractorbeam2.NewDrone(code)
	affectedArea := drone.Fit(100)
	fmt.Printf("Affected area: %d\n", affectedArea)
}
