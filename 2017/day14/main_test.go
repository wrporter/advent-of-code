package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_countRegions(t *testing.T) {
	tests := []struct {
		grid []string
		want interface{}
	}{
		{
			grid: []string{
				"10",
			},
			want: 1,
		},
		{
			grid: []string{
				"10",
				"10",
			},
			want: 1,
		},
		{
			grid: []string{
				"10000111",
				"11110001",
				"00000001",
				"10110000",
				"00000001",
				"10110101",
			},
			want: 8,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := countRegions(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRegions() = %v, want %v", got, tt.want)
			}
		})
	}
}
