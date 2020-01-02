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
			if got := countNiceStringsPart1(tt.words); got != tt.want {
				t.Errorf("countNiceStringsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasRepeatPairWithoutOverlap(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"aaa", false},
		{"aaaa", true},
		{"abab", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasRepeatPairWithoutOverlap(tt.str); got != tt.want {
				t.Errorf("hasRepeatPairWithoutOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasRepeatLetterWithOneInBetween(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"aaa", true},
		{"aaaa", true},
		{"abab", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := hasRepeatLetterWithOneInBetween(tt.str); got != tt.want {
				t.Errorf("hasRepeatLetterWithOneInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
