package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_playMemoryGame(t *testing.T) {
	type args struct {
		startNumbers []int
		numTurns     int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				[]int{0, 3, 6},
				10,
			},
			0,
		},
		{
			args{
				[]int{0, 3, 6},
				2020,
			},
			436,
		},
		{
			args{
				[]int{1, 3, 2},
				2020,
			},
			1,
		},
		{
			args{
				[]int{2, 1, 3},
				2020,
			},
			10,
		},
		{
			args{
				[]int{1, 2, 3},
				2020,
			},
			27,
		},
		{
			args{
				[]int{2, 3, 1},
				2020,
			},
			78,
		},
		{
			args{
				[]int{3, 2, 1},
				2020,
			},
			438,
		},
		{
			args{
				[]int{3, 1, 2},
				2020,
			},
			1836,
		},
		{
			args{
				[]int{0, 3, 6},
				30000000,
			},
			175594,
		},
		{
			args{
				[]int{1, 3, 2},
				30000000,
			},
			2578,
		},
		{
			args{
				[]int{2, 1, 3},
				30000000,
			},
			3544142,
		},
		{
			args{
				[]int{1, 2, 3},
				30000000,
			},
			261214,
		},
		{
			args{
				[]int{2, 3, 1},
				30000000,
			},
			6895259,
		},
		{
			args{
				[]int{3, 2, 1},
				30000000,
			},
			18,
		},
		{
			args{
				[]int{3, 1, 2},
				30000000,
			},
			362,
		},
		{
			args{
				[]int{19, 0, 5, 1, 10, 13},
				2020,
			},
			1015,
		},
		{
			args{
				[]int{19, 0, 5, 1, 10, 13},
				30000000,
			},
			201,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := playMemoryGame(tt.args.startNumbers, tt.args.numTurns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playMemoryGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
