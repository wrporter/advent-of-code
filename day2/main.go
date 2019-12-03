package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/file"
)

func main() {
	moduleMassStrings, _ := file.ReadFile("./day1/input.txt")
	moduleMasses, _ := conversion.ToInts(moduleMassStrings)
	fmt.Println(calculateRequiredFuel(moduleMasses))
}

func calculateRequiredFuel(moduleMasses []int) int {
	sum := 0
	for _, mass := range moduleMasses {
		sum += calculateModuleFuel(mass)
	}
	return sum
}

func calculateFuel(mass int) int {
	return mass / 3 - 2
}

func calculateModuleFuel(mass int) int {
	nextMass := calculateFuel(mass)
	totalFuel := 0
	for nextMass > 0 {
		totalFuel += nextMass
		nextMass = calculateFuel(nextMass)
	}
	return totalFuel
}
