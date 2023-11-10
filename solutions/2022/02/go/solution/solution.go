package solution

import (
	mymath "github.com/wrporter/advent-of-code/internal/common/ints"
	"strings"
)

type Shape = int

const (
	//_ Shape = iota
	Rock Shape = iota
	Paper
	Scissors
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	shapes := map[byte]Shape{
		'A': Rock,
		'B': Paper,
		'C': Scissors,

		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}

	score := 0
	for _, round := range strings.Split(input, "\n") {
		opponent := shapes[round[0]]
		self := shapes[round[2]]

		if (self+1)%3 == opponent {
			score += 0 // Lose
		} else if self == opponent {
			score += 3 // Draw
		} else {
			score += 6 // Win
		}

		score += self + 1
	}

	return score
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	score := 0
	for _, round := range strings.Split(input, "\n") {
		opponent := Shape(round[0] - 'A')
		outcome := round[2]
		var self Shape
		const (
			lose = 'X'
			draw = 'Y'
			win  = 'Z'
		)

		if outcome == lose {
			self = mymath.WrapMod(opponent-1, 3)
			score += 0
		} else if outcome == draw {
			self = opponent
			score += 3
		} else if outcome == win {
			self = (opponent + 1) % 3
			score += 6
		}

		score += self + 1
	}

	return score
}
