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
		{[]string{"1"}, 0},
		{[]string{"12"}, 3},
		{[]string{"23"}, 2},
		{[]string{"1024"}, 31},
		{[]string{"265149"}, 438},
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
		{[]string{"1"}, 2},
		{[]string{"2"}, 4},
		{[]string{"3"}, 4},
		{[]string{"23"}, 25},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
