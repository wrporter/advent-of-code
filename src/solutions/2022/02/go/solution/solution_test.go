package solution

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	s     = New()
	input = s.ReadInputFromTests()
)

var tests = []struct {
	input string
	args1 []interface{}
	args2 []interface{}
	want1 interface{}
	want2 interface{}
}{
	{input: `A Z`, want1: 3, want2: 8}, // Lose, Win
	{input: `C Y`, want1: 2, want2: 6}, // Lose, Draw
	{input: `B X`, want1: 1, want2: 1}, // Lose, Lose
	{input: `A X`, want1: 4, want2: 3}, // Draw, Lose
	{input: `B Y`, want1: 5, want2: 5}, // Draw, Draw
	{input: `C Z`, want1: 6, want2: 7}, // Draw, Win
	{input: `C X`, want1: 7, want2: 2}, // Win, Lose
	{input: `A Y`, want1: 8, want2: 4}, // Win, Draw
	{input: `B Z`, want1: 9, want2: 9}, // Win, Win
	{
		input: `A Y
B X
C Z`,
		want1: 15,
		want2: 12,
	},
	{
		input: input,
		want1: 12740,
		want2: 11980,
	},
}

func Test_Part1(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part1(tt.input, tt.args1...); !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Part1() = %v, want %v", got, tt.want1)
			}
		})
	}
}

func Test_Part2(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part2(tt.input, tt.args2...); !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("Part2() = %v, want %v", got, tt.want2)
			}
		})
	}
}
