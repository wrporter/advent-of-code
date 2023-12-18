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
	{
		input: `.....
..##.
..#..
.....
..##.
.....`,
		args1: []interface{}{},
		args2: []interface{}{},
		want1: 25,
		want2: 4,
	},
	{
		input: `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`,
		args1: []interface{}{},
		args2: []interface{}{},
		want1: 110,
		want2: 20,
	},
	{
		input: input,
		args1: []interface{}{},
		args2: []interface{}{},
		want1: 3906,
		want2: 895,
	},
}

func Test_Part1(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part1(tt.input, tt.args1...); !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Part2() = %v, want %v", got, tt.want1)
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
