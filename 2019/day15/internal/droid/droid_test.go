package droid

import (
	"fmt"
	"testing"
)

func TestDirection_Opposite(t *testing.T) {
	tests := []struct {
		d    Direction
		want Direction
	}{
		{North, South},
		{South, North},
		{West, East},
		{East, West},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := tt.d.Opposite(); got != tt.want {
				t.Errorf("Opposite() = %v, want %v", got, tt.want)
			}
		})
	}
}
