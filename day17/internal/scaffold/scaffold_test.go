package scaffold

import (
	"fmt"
	"reflect"
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

func Test_compressCommands(t *testing.T) {
	tests := []struct {
		commands string
		want     string
		want1    map[string]string
	}{
		{
			"R,6,L,10,R,8,R,8,R,12,L,8,L,10,R,6,L,10,R,8,R,8,R,12,L,10,R,6,L,10,R,12,L,8,L,10,R,12,L,10,R,6,L,10,R,6,L,10,R,8,R,8,R,12,L,8,L,10,R,6,L,10,R,8,R,8,R,12,L,10,R,6,L,10",
			"A,B,A,C,B,C,A,B,A,C",
			map[string]string{
				"A": "R,6,L,10,R,8,R,8",
				"B": "R,12,L,8,L,10",
				"C": "R,12,L,10,R,6,L,10",
			},
		},
		{
			"L,12,R,8,L,6,R,8,L,6,R,8,L,12,L,12,R,8,L,12,R,8,L,6,R,8,L,6,L,12,R,8,L,6,R,8,L,6,R,8,L,12,L,12,R,8,L,6,R,6,L,12,R,8,L,12,L,12,R,8,L,6,R,6,L,12,L,6,R,6,L,12,R,8,L,12,L,12,R,8",
			"A,B,A,A,B,C,B,C,C,B",
			map[string]string{
				"A": "L,12,R,8,L,6,R,8,L,6",
				"B": "R,8,L,12,L,12,R,8",
				"C": "L,6,R,6,L,12",
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got, got1 := compressCommands(tt.commands)
			if got != tt.want {
				t.Errorf("compressCommands() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("compressCommands() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
