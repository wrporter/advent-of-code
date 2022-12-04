package main

import (
	"fmt"
	"testing"
)

func Test_lookAndSay(t *testing.T) {
	tests := []struct {
		sequence string
		want     string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := lookAndSay(tt.sequence); got != tt.want {
				t.Errorf("lookAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
