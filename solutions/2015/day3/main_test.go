package main

import (
	"fmt"
	"testing"
)

func Test_deliver(t *testing.T) {
	tests := []struct {
		directions string
		want       int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := deliver(tt.directions); got != tt.want {
				t.Errorf("deliver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deliverWithRobo(t *testing.T) {
	tests := []struct {
		directions string
		want       int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := deliverWithRobo(tt.directions); got != tt.want {
				t.Errorf("deliverWithRobo() = %v, want %v", got, tt.want)
			}
		})
	}
}
