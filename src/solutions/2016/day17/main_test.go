package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getShortestPath(t *testing.T) {
	tests := []struct {
		passcode string
		want     string
	}{
		{"hijkl", ""},
		{"ihgpwlah", "DDRRRD"},
		{"kglvqrro", "DDUDRLRRUDRD"},
		{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getShortestPath(tt.passcode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLongestPathLength(t *testing.T) {
	tests := []struct {
		passcode string
		want     int
	}{
		{"hijkl", 0},
		{"ihgpwlah", 370},
		{"kglvqrro", 492},
		{"ulqzkmiv", 830},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getLongestPathLength(tt.passcode); got != tt.want {
				t.Errorf("getLongestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
