package solution

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	s     = New()
	input = s.ReadInputPrefix("../../../../../")
)

var tests = []struct {
	input string
	args1 []interface{}
	args2 []interface{}
	want1 interface{}
	want2 interface{}
}{
	{
		input: `+1`,
		want1: 1,
		want2: 0,
	},
	{
		input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
		want1: 4361,
		want2: 467835,
	},
	{
		input: input,
		want1: 550064,
		want2: 85010461,
	},
}

func TestPart1(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part1(tt.input, tt.args1...); !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Part1() = %v, want %v", got, tt.want1)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part2(tt.input, tt.args2...); !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("Part2() = %v, want %v", got, tt.want2)
			}
		})
	}
}

func BenchmarkSolution_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Part1(tests[0].input)
	}
}

func BenchmarkSolution_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Part2(tests[0].input)
	}
}
