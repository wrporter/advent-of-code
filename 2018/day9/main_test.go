package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_playMarbles(t *testing.T) {
	tests := []struct {
		numPlayers int
		lastMarble int
		want       int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := playMarbles(tt.numPlayers, tt.lastMarble); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
