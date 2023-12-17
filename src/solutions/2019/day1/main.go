package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"fmt"
)

func main() {
	moduleMassStrings, _ := file.ReadFile("./2019/day1/input.txt")
	moduleMasses, _ := convert.ToInts(moduleMassStrings)
	fmt.Println(calculateRequiredFuel(moduleMasses))
}

func calculateRequiredFuel(moduleMasses []int) int {
	sum := 0
	for _, mass := range moduleMasses {
		sum += calculateModuleFuel(mass)
	}
	return sum
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

func calculateFuel(mass int) int {
	return mass/3 - 2
}
