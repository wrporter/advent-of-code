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
		input: `1000`,
		args1: []interface{}{},
		args2: []interface{}{},
		want1: "TBD",
		want2: "TBD",
	},
	{
		input: input,
		args1: []interface{}{},
		args2: []interface{}{},
		want1: "TBD",
		want2: "TBD",
	},
}

func TestPart1(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part1(tt.input, tt.args1...); !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Part2() = %v, want %v", got, tt.want1)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := s.Part1(tt.input, tt.args2...); !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("Part2() = %v, want %v", got, tt.want2)
			}
		})
	}
}
