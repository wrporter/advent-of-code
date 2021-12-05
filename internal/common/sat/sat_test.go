package sat

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateSummedAreaTable(t *testing.T) {
	tests := []struct {
		grid [][]int
		want [][]int
	}{
		{
			grid: [][]int{
				{4, 5, 2, 1},
				{0, 9, 3, 2},
				{5, 6, 8, 1},
				{2, 3, 0, 0},
			},
			want: [][]int{
				{4, 9, 11, 12},
				{4, 18, 23, 26},
				{9, 29, 42, 46},
				{11, 34, 47, 51},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := CreateSummedAreaTable(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSummedAreaTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
