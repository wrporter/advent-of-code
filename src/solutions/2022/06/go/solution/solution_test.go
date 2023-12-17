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
		input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
		want1: 7,
		want2: 19,
	},
	{
		input: `bvwbjplbgvbhsrlpgdmjqwftvncz`,
		want1: 5,
		want2: 23,
	},
	{
		input: `nppdvjthqldpwncqszvftbrmjlhg`,
		want1: 6,
		want2: 23,
	},
	{
		input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
		want1: 10,
		want2: 29,
	},
	{
		input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
		want1: 11,
		want2: 26,
	},
	{
		input: input,
		want1: 1042,
		want2: 2980,
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
