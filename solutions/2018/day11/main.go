package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/sat"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 11
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	grid := newFuelCellGrid(1723)
	windowSize := 3
	sumTable := sat.NewSummedAreaTable(grid)
	_, topLeftCorner := getTotalPowerCorner(sumTable, windowSize)
	return fmt.Sprintf("%d,%d", topLeftCorner.X, topLeftCorner.Y)
}

func part2(input []string) interface{} {
	grid := newFuelCellGrid(1723)
	sumTable := sat.NewSummedAreaTable(grid)

	maxPower := 0
	var topLeftCorner geometry.Point
	size := 0

	for windowSize := 1; windowSize <= len(grid); windowSize++ {
		power, corner := getTotalPowerCorner(sumTable, windowSize)
		if power > maxPower {
			maxPower = power
			topLeftCorner = corner
			size = windowSize
		}
	}

	return fmt.Sprintf("%d,%d,%d", topLeftCorner.X, topLeftCorner.Y, size)
}

func getTotalPowerCorner(sumTable *sat.SummedAreaTable, windowSize int) (int, geometry.Point) {
	maxPower := 0
	var topLeftCorner geometry.Point

	for y := 0; y <= sumTable.Height-windowSize; y++ {
		for x := 0; x <= sumTable.Width-windowSize; x++ {
			power := sumTable.SumWindow(x, y, x+windowSize-1, y+windowSize-1)
			if power > maxPower {
				maxPower = power
				topLeftCorner = geometry.NewPoint(x+1, y+1)
			}
		}
	}

	return maxPower, topLeftCorner
}

func newFuelCellGrid(serialNumber int) [][]int {
	size := 300
	grid := make([][]int, size)
	for y := 0; y < size; y++ {
		grid[y] = make([]int, size)
		for x := 0; x < size; x++ {
			grid[y][x] = getPowerLevel(x+1, y+1, serialNumber)
		}
	}
	return grid
}

func getPowerLevel(x int, y int, serialNumber int) int {
	rackId := x + 10
	power := rackId * y
	power += serialNumber
	power *= rackId
	power = getDigit(power, 3)
	power -= 5
	return power
}

func getDigit(number, position int) int {
	r := number % int(math.Pow(10, float64(position)))
	return r / int(math.Pow(10, float64(position-1)))
}
