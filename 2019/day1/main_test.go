package main

import (
	"fmt"
	"testing"
)

func Test_calculateModuleFuel(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := calculateModuleFuel(tt.mass); got != tt.want {
				t.Errorf("calculateModuleFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateRequiredFuel(t *testing.T) {
	tests := []struct {
		moduleMasses []int
		want         int
	}{
		{
			[]int{92349, 57040, 64079, 121555, 143735, 64642, 104858, 144446, 88871, 62338, 113424, 59960, 53999, 86867, 67224, 124130, 108921, 130492, 120361, 74426, 70397, 88106, 125442, 74237, 137818, 66633, 71756, 143276, 143456, 135698, 121124, 67739, 112861, 78572, 73565, 111899, 57543, 130314, 121605, 121426, 117143, 129957, 98042, 104760, 144846, 131238, 101076, 53328, 83592, 104077, 101952, 54137, 115363, 60556, 133086, 113361, 117829, 75003, 93729, 140022, 126219, 59907, 140589, 91812, 50485, 56232, 92858, 106820, 123423, 98553, 135315, 95583, 72278, 98702, 55709, 146773, 89719, 134752, 79562, 70455, 88468, 139824, 138646, 117516, 123267, 113754, 120353, 139145, 53219, 63053, 131434, 91705, 53650, 145234, 78461, 119587, 108976, 113613, 121790, 120366},
			5055835,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := calculateRequiredFuel(tt.moduleMasses); got != tt.want {
				t.Errorf("calculateRequiredFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}