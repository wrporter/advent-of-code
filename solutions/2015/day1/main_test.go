package main

import (
	"fmt"
	"testing"
)

func Test_elevator(t *testing.T) {
	tests := []struct {
		stops string
		want  int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := elevator(tt.stops); got != tt.want {
				t.Errorf("elevator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstTimeAt(t *testing.T) {
	tests := []struct {
		instructions string
		stop         int
		want         int
	}{
		{")", -1, 1},
		{"()())", -1, 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := firstTimeAt(tt.instructions, tt.stop); got != tt.want {
				t.Errorf("firstTimeAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
