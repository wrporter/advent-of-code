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
		input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
		want1: 2,
		want2: 2,
	},
	{
		input: `LLR

		AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ)`,
		want1: 6,
		want2: 6,
	},
	{
		input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
		want1: -1,
		want2: 6,
	},
	{
		input: input,
		want1: 18827,
		want2: 20220305520997,
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
