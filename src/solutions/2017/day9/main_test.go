package main

import (
	"fmt"
	"testing"
)

func Test_getScore(t *testing.T) {
	tests := []struct {
		stream      string
		wantScore   int
		wantGarbage int
	}{
		{"{}", 1, 0},
		{"{{{}}}", 6, 0},
		{"{{},{}}", 5, 0},
		{"{{{},{},{{}}}}", 16, 0},
		{"{<a>,<a>,<a>,<a>}", 1, 4},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
		{"<>", 0, 0},
		{"<random characters>", 0, 17},
		{"<<<<>", 0, 3},
		{"<{!>}>", 0, 2},
		{"<!!>", 0, 0},
		{"<!!!>>", 0, 0},
		{"<{o\"i!a,<{i<a>", 0, 10},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			score, garbage := countScoreAndGarbage(tt.stream)
			if score != tt.wantScore {
				t.Errorf("countScoreAndGarbage() = %v, want %v", score, tt.wantScore)
			}
			if garbage != tt.wantGarbage {
				t.Errorf("countScoreAndGarbage() = %v, want %v", garbage, tt.wantGarbage)
			}
		})
	}
}
