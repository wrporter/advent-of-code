package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_hash(t *testing.T) {
	tests := []struct {
		size    int
		lengths []int
		want    interface{}
	}{
		{5, []int{3, 4, 1, 5}, 12},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := singleRoundHash(tt.size, tt.lengths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleRoundHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		input []string
		want  interface{}
	}{
		{[]string{""}, "a2582a3a0e66e6e86e3812dcb672a272"},
		{[]string{"AoC 2017"}, "33efeb34ea91902bb2f59c9920caa6cd"},
		{[]string{"1,2,3"}, "3efbe78a8d82f29979031a4aa0b16a9d"},
		{[]string{"1,2,4"}, "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
