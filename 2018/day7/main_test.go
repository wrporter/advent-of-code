package main

import (
	"fmt"
	"testing"
)

func Test_instruction_Less(t *testing.T) {
	tests := []struct {
		a    *instruction
		b    *instruction
		want bool
	}{
		{
			a: &instruction{
				step:   "A",
				before: []string{"B"},
			},
			b: &instruction{
				step:   "B",
				before: nil,
			},
			want: true,
		},
		{
			a: &instruction{
				step:   "E",
				before: nil,
			},
			b: &instruction{
				step:   "F",
				before: []string{"E"},
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := tt.a.Less(tt.b); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBefore(t *testing.T) {
	tests := []struct {
		m     map[string][]string
		step1 string
		step2 string
		want  bool
	}{
		{
			m: map[string][]string{
				"A": {"B", "D"},
				"B": {"E"},
				"C": {"A", "F"},
				"D": {"E"},
				"E": {},
				"F": {"E"},
			},
			step1: "A",
			step2: "B",
			want:  true,
		},
		{
			m: map[string][]string{
				"A": {"B", "D"},
				"B": {"E"},
				"C": {"A", "F"},
				"D": {"E"},
				"E": {},
				"F": {"E"},
			},
			step1: "A",
			step2: "C",
			want:  false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := less(tt.m, tt.step1, tt.step2); got != tt.want {
				t.Errorf("less() = %v, want %v", got, tt.want)
			}
		})
	}
}
