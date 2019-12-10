package image

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		data   string
		width  int
		height int
	}
	tests := []struct {
		args args
		want *Image
	}{
		{
			args{"123456789012", 3, 2},
			&Image{[][][]int{
				{
					{1, 2, 3},
					{4, 5, 6},
				},
				{
					{7, 8, 9},
					{0, 1, 2},
				},
			},
				3,
				2,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := New(tt.args.data, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
