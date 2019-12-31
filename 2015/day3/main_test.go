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
