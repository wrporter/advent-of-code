package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		input []string
		want  interface{}
	}{
		{
			[]string{"389125467"},
			"67384529",
		},
		{
			[]string{"916438275"},
			"39564287",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		input []string
		want  interface{}
	}{
		{
			[]string{"389125467"},
			149245887792,
		},
		{
			[]string{"916438275"},
			404431096944,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
