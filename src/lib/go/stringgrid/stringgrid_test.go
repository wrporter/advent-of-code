package stringgrid

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate90Clockwise(t *testing.T) {
	tests := []struct {
		array []string
		want  []string
	}{
		{
			[]string{
				"123",
				"456",
				"789",
			},
			[]string{
				"741",
				"852",
				"963",
			},
		},
		{
			[]string{
				"123",
				"456",
			},
			[]string{
				"41",
				"52",
				"63",
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Rotate90Clockwise(tt.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate90Clockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}
