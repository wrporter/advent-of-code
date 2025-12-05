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
	//	{
	//		input: `3-5
	//10-14
	//16-20
	//12-18
	//
	//1
	//5
	//8
	//11
	//17
	//32`,
	//		want1: 3,
	//		want2: 14,
	//	},
	{
		input: `16-20
17-19

1`,
		want1: 0,
		want2: 5,
	},
	{
		input: input,
		want1: 607,
		want2: 342433357244012,
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

var rangeTests = []struct {
	input string
	args  []interface{}
	want  interface{}
}{
	{
		// single
		input: `1-3`,
		want:  3,
	},
	{
		// before
		input: `1-3
4-6`,
		want: 6,
	},
	{
		// end meets next
		input: `1-4
4-6`,
		want: 6,
	},
	{
		// overlaps into next
		input: `1-5
4-6`,
		want: 6,
	},
	{
		// next is completely inside
		input: `1-6
4-6`,
		want: 6,
	},
	{
		// next is completely inside
		input: `1-6
4-6`,
		want: 6,
	},
	{
		// first is completely inside next
		input: `1-3
1-6`,
		want: 6,
	},
}

func Test_Part2_Ranges(t *testing.T) {
	for i, tt := range rangeTests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part2(tt.input, tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
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
