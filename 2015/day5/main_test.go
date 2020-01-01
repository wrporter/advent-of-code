package main

import (
	"fmt"
	"testing"
)

func Test_countNiceStrings(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{[]string{"ugknbfddgicrmopn"}, 1},
		{[]string{"aaa"}, 1},
		{[]string{"jchzalrnumimnmhp"}, 0},
		{[]string{"haegwjzuvuyypxyu"}, 0},
		{[]string{"dvszwmarrgswjxmb"}, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := countNiceStrings(tt.words); got != tt.want {
				t.Errorf("countNiceStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
