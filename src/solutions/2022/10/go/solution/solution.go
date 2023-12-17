package solution

import (
	"aoc/src/lib/go/convert"
	"math"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	instructions := strings.Split(input, "\n")
	signalStrength := 0

	cpu(instructions, 220, func(cycle int, registerX int) {
		if (cycle-20)%40 == 0 {
			signalStrength += cycle * registerX
		}
	})

	return signalStrength
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	instructions := strings.Split(input, "\n")
	crt := make([][]string, 6)
	for row := range crt {
		crt[row] = make([]string, 40)
		for col := range crt[row] {
			crt[row][col] = "."
		}
	}

	cpu(instructions, 240, func(cycle int, registerX int) {
		row := int(math.Floor(float64((cycle - 1) / 40)))
		col := (cycle - 1) % 40
		isWithinSprite := col >= registerX-1 && col <= registerX+1
		if isWithinSprite {
			crt[row][col] = "#"
		} else {
			crt[row][col] = "."
		}
	})

	var rows []string
	for _, row := range crt {
		rows = append(rows, strings.Join(row, ""))
	}
	render := strings.Join(rows, "\n")

	return render
}

func cpu(instructions []string, maxCycles int, process func(cycles int, registerX int)) {
	address := 0
	registerX := 1
	counter := 1

	for cycle := 1; cycle <= maxCycles && address < len(instructions); cycle += 1 {
		args := strings.Split(instructions[address], " ")
		operation := args[0]

		process(cycle, registerX)

		if operation == "noop" {
			address += 1
			counter = 1
		} else if operation == "addx" && counter == 2 {
			registerX += convert.StringToInt(args[1])
			address += 1
			counter = 1
		} else {
			counter += 1
		}
	}
}
