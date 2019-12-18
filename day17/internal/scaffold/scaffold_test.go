package scaffold

import (
	"fmt"
	"testing"
)

func TestDirection_GetTurnDirection(t *testing.T) {
	tests := []struct {
		d    Direction
		dest Direction
		want TurnDirection
	}{
		{Up, Right, TurnRight},
		{Up, Left, TurnLeft},
		{Right, Down, TurnRight},
		{Right, Up, TurnLeft},
		{Down, Left, TurnRight},
		{Down, Right, TurnLeft},
		{Left, Down, TurnLeft},
		{Left, Up, TurnRight},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := tt.d.GetTurnDirection(tt.dest); got != tt.want {
				t.Errorf("GetTurnDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
