package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		input []string
		want  interface{}
	}{
		{
			[]string{"0 2 7 4"},
			5,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
