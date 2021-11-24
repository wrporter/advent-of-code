package main

import (
	"fmt"
	"testing"
)

func Test_getFirstTriplet(t *testing.T) {
	tests := []struct {
		stream string
		want   string
	}{
		{"aaa", "aaa"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getFirstTriplet(tt.stream); got != tt.want {
				t.Errorf("getFirstTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
