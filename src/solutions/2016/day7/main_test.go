package main

import (
	"fmt"
	"testing"
)

func Test_hasRepeatBackwardsPair(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"", false},
		{"a", false},
		{"aa", false},
		{"asdfgh", false},
		{"abba", true},
		{"fabbak", true},
		{"ioxxoj", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasRepeatBackwardsPair(tt.str); got != tt.want {
				t.Errorf("hasRepeatBackwardsPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
