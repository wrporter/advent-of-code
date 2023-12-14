package runegrid

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate1(t *testing.T) {
	tests := []struct {
		grid [][]rune
		want [][]rune
	}{
		{
			[][]rune{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},

			[][]rune{
				{6, 3, 0},
				{7, 4, 1},
				{8, 5, 2},
			},
		},
		{
			[][]rune{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},

			[][]rune{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Rotate(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRotateFull(t *testing.T) {
	tests := []struct {
		grid [][]rune
	}{
		{
			[][]rune{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got := Rotate(tt.grid)
			got = Rotate(got)
			got = Rotate(got)
			got = Rotate(got)
			if !reflect.DeepEqual(got, tt.grid) {
				t.Errorf("Rotate() = %v, want %v", got, tt.grid)
			}
		})
	}
}
