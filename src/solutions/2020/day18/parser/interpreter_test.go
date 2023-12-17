package parser

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expression string
		want       int
	}{
		{
			"2 + 3",
			5,
		},
		{
			"2 * 3 + (4 * 5)",
			26,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			437,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			12240,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			13632,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Evaluate(tt.expression); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

//1 + (2 * 3) + (4 * (5 + 6)) still becomes 51.
//2 * 3 + (4 * 5) becomes 46.
//5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 1445.
//5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 669060.
//((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 23340.

func TestEvaluate_precedence(t *testing.T) {
	tests := []struct {
		expression string
		want       int
	}{
		{
			"2 + 3",
			5,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			46,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			1445,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			669060,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			23340,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Evaluate(tt.expression); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
