package main

import (
	"fmt"
	"testing"
)

func Test_hasIncreasingStraight(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"", false},
		{"a", false},
		{"abc", true},
		{"ahijmxy", true},
		{"ahikmxy", false},
		{"ahikmxyz", true},
		{"abcdffaa", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasIncreasingStraight(tt.password); got != tt.want {
				t.Errorf("hasIncreasingStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"abcdffaa", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := isValid(tt.password); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasBadLetters(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"abcdffaa", false},
		{"i", true},
		{"o", true},
		{"l", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasBadLetters(tt.password); got != tt.want {
				t.Errorf("hasBadLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasNonOverlappingPairs(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"aabb", true},
		{"abcdffaa", true},
		{"aaa", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasNonOverlappingPairs(tt.password); got != tt.want {
				t.Errorf("hasNonOverlappingPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextValid(t *testing.T) {
	tests := []struct {
		password string
		want     string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
		{"vzbxkghb", "vzbxxyzz"},
		{"vzbxxyzz", "vzcaabcc"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got := nextValid(tt.password)
			if got != tt.want {
				t.Errorf("nextValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
