package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
)

var keypad1 = [][]rune{
	{'1', '2', '3'},
	{'4', '5', '6'},
	{'7', '8', '9'},
}

var keypad2 = [][]rune{
	{'0', '0', '1', '0', '0'},
	{'0', '2', '3', '4', '0'},
	{'5', '6', '7', '8', '9'},
	{'0', 'A', 'B', 'C', '0'},
	{'0', '0', 'D', '0', '0'},
}

func solveKeypadCombo(input []string, keypad [][]rune) string {
	combo := make([]rune, len(input))

	for index, instructionSet := range input {
		position := geometry.NewPoint(1, 1)

		for _, instruction := range instructionSet {
			var direction geometry.Direction
			switch instruction {
			case 'U':
				direction = geometry.Up
			case 'R':
				direction = geometry.Right
			case 'D':
				direction = geometry.Down
			case 'L':
				direction = geometry.Left
			default:
				direction = geometry.Up
			}

			nextPosition := position.Move(direction)
			if nextPosition.X >= 0 && nextPosition.X < len(keypad[0]) &&
				nextPosition.Y >= 0 && nextPosition.Y < len(keypad) &&
				keypad[nextPosition.Y][nextPosition.X] != '0' {
				position = nextPosition
			}
		}

		combo[index] = keypad[position.Y][position.X]
	}

	return string(combo)
}

func main() {
	input, _ := file.ReadFile("./2016/day2/input.txt")
	answer1 := solveKeypadCombo(input, keypad1)
	answer2 := solveKeypadCombo(input, keypad2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
