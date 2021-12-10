package main

import (
	"fmt"
	"testing"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		input        string
		wantGot      string
		wantExpected string
		wantIndex    int
	}{
		{"(", EOF, ")", 1},
		{"(]", "]", ")", 1},
		{"()", EOF, EOF, 2},
		{"(())", EOF, EOF, 4},
		{"(<>)", EOF, EOF, 4},
		{"(<>]", "]", ")", 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got, expected, index := parse(tt.input)
			if got != tt.wantGot {
				t.Errorf("parse() got = %v, wantGot %v", got, tt.wantGot)
			}
			if expected != tt.wantExpected {
				t.Errorf("parse() expected = %v, wantExpected %v", expected, tt.wantExpected)
			}
			if index != tt.wantIndex {
				t.Errorf("parse() index = %v, wantIndex %v", index, tt.wantIndex)
			}
		})
	}
}
