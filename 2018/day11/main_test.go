package main

import (
	"fmt"
	"testing"
)

func Test_getDigit(t *testing.T) {
	tests := []struct {
		number   int
		position int
		want     int
	}{
		{12345, 1, 5},
		{5, 2, 0},
		{25, 2, 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getDigit(tt.number, tt.position); got != tt.want {
				t.Errorf("getDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPowerLevel(t *testing.T) {
	tests := []struct {
		x            int
		y            int
		serialNumber int
		want         int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
		{33, 45, 18, 4},
		{34, 45, 18, 4},
		{35, 45, 18, 4},
		{33, 46, 18, 3},
		{34, 46, 18, 3},
		{35, 46, 18, 4},
		{33, 47, 18, 1},
		{34, 47, 18, 2},
		{35, 47, 18, 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getPowerLevel(tt.x, tt.y, tt.serialNumber); got != tt.want {
				t.Errorf("getPowerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
