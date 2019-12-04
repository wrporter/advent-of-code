package main

import (
	"fmt"
	"testing"
)

func Test_digitsNeverDecrease(t *testing.T) {
	tests := []struct {
		candidate string
		want      bool
	}{
		{"1", true},
		{"111", true},
		{"121", false},
		{"1221", false},
		{"12311", false},
		{"12345", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := digitsNeverDecrease(tt.candidate); got != tt.want {
				t.Errorf("Test_digitsNeverDecrease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasTwoConsecutiveDigits(t *testing.T) {
	tests := []struct {
		candidate string
		want      bool
	}{
		{"1", false},
		{"11", true},
		{"121", false},
		{"1221", true},
		{"12311", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasTwoConsecutiveDigits(tt.candidate); got != tt.want {
				t.Errorf("hasTwoConsecutiveDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasOnlyTwoConsecutiveDigits(t *testing.T) {
	tests := []struct {
		candidate string
		want      bool
	}{
		{"1", false},
		{"11", true},
		{"121", false},
		{"1221", true},
		{"12311", true},
		{"111111", false},
		{"111122", true},
		{"221111", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasOnlyTwoConsecutiveDigits(tt.candidate); got != tt.want {
				t.Errorf("hasOnlyTwoConsecutiveDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
