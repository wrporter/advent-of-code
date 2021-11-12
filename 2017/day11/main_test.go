package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getDistance(t *testing.T) {
	tests := []struct {
		stepsStr string
		want     int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"se,sw,se,sw,sw", 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getDistance(tt.stepsStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
