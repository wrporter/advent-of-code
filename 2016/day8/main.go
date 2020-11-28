package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
)

type Command string

const (
	Rect         Command = "rect"
	RotateRow    Command = "rotate row"
	RotateColumn Command = "rotate column"
)

type Operation struct {
	Command Command
	A       int
	B       int
}

var regex = regexp.MustCompile(`^(rect|rotate row|rotate column) (|y=|x=)(\d+)(x| by |)(\d+)$`)

func part1(input []string) int {
	operations := parseOperations(input)
	screen := makeGrid(6, 50)
	performOperations(operations, screen)
	return countOnPixels(screen)
}

func part2(input []string) string {
	operations := parseOperations(input)
	screen := makeGrid(6, 50)
	performOperations(operations, screen)
	return renderGrid(screen)
}

func performOperations(operations []Operation, screen [][]bool) {
	for _, operation := range operations {
		switch operation.Command {
		case Rect:
			fillRect(screen, operation.A, operation.B)
		case RotateRow:
			screen[operation.A] = rotate(screen[operation.A], operation.B)
		case RotateColumn:
			column := make([]bool, len(screen))
			for row := range screen {
				column[row] = screen[row][operation.A]
			}

			column = rotate(column, operation.B)
			for row := range screen {
				screen[row][operation.A] = column[row]
			}
		}
	}
}

func countOnPixels(screen [][]bool) int {
	count := 0
	for _, row := range screen {
		for _, pixel := range row {
			if pixel {
				count++
			}
		}
	}
	return count
}

func renderGrid(grid [][]bool) string {
	result := ""
	for _, row := range grid {
		for _, position := range row {
			if position {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func rotate(values []bool, amount int) []bool {
	if amount < 0 || len(values) == 0 {
		return values
	}

	rotation := len(values) - (amount % len(values))
	values = append(values[rotation:], values[:rotation]...)

	return values
}

func fillRect(screen [][]bool, width int, height int) {
	for row := 0; row < height && row < len(screen); row++ {
		for col := 0; col < width && col < len(screen[row]); col++ {
			screen[row][col] = true
		}
	}
}

func makeGrid(height, width int) [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

func parseOperations(input []string) []Operation {
	operations := make([]Operation, len(input))
	for i, line := range input {
		operations[i] = toOperation(line)
	}
	return operations
}

func toOperation(operationString string) Operation {
	match := regex.FindStringSubmatch(operationString)
	operation := Operation{
		Command: Command(match[1]),
		A:       conversion.StringToInt(match[3]),
		B:       conversion.StringToInt(match[5]),
	}
	return operation
}

func main() {
	input, _ := file.ReadFile("./2016/day8/input.txt")
	answer1 := part1(input)
	answer2 := part2(input)
	fmt.Println(answer1)
	fmt.Println(answer2)
	fmt.Println("ZJHRKCPLYJ")
}
