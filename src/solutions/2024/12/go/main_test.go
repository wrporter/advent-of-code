package main

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
		input: `AAAA
BBCD
BBCC
EEEC`,
		want1: 140,
		want2: 80,
	},
	{
		input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
		want1: 772,
		want2: 436,
	},
	{
		input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
		want1: 692,
		want2: 236,
	},
	{
		input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
		want1: 1184,
		want2: 368,
	},
	{
		input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
		want1: 1930,
		want2: 1206,
	},
	{
		input: input,
		want1: 1461806,
		want2: 887932,
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

func BenchmarkSolution_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Part1(input)
	}
}

func BenchmarkSolution_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Part2(input)
	}
}
