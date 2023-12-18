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
		input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`,
		want1: 21,
		want2: 525152,
	},
	{
		input: input,
		want1: 7204,
		want2: 1672318386674,
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

func Test_countArrangements(t *testing.T) {
	tests := []struct {
		record Record
		want   int
	}{
		{Record{"???", []int{1}}, 3},
		{Record{"??..??...?##.", []int{1, 1, 3}}, 4},
		{Record{"?###????????", []int{3, 2, 1}}, 10},
		{Record{"?#?????", []int{3, 1}}, 5},
		{Record{"???.###", []int{1, 1, 3}}, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := countArrangements(tt.record); got != tt.want {
				t.Errorf("countArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
