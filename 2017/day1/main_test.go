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
			input: []string{"1122"},
			want:  3,
		},
		{
			input: []string{"1111"},
			want:  4,
		},
		{
			input: []string{"1234"},
			want:  0,
		},
		{
			input: []string{"91212129"},
			want:  9,
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
			input: []string{"1212"},
			want:  6,
		},
		{
			input: []string{"1221"},
			want:  0,
		},
		{
			input: []string{"123425"},
			want:  4,
		},
		{
			input: []string{"123123"},
			want:  12,
		},
		{
			input: []string{"12131415"},
			want:  4,
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
