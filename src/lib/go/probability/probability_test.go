package probability

import (
	"reflect"
	"testing"
)

func TestComboSize(t *testing.T) {
	type args struct {
		values    []int
		startSize int
		endSize   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "returns single values",
			args: args{
				values:    []int{1, 2},
				startSize: 1,
				endSize:   1,
			},
			want: [][]int{{1}, {2}},
		},
		{
			name: "returns 1-2 values",
			args: args{
				values:    []int{1, 2},
				startSize: 1,
				endSize:   2,
			},
			want: [][]int{{1}, {2}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [][]int
			ComboSize(tt.args.values, tt.args.startSize, tt.args.endSize, func(ints []int) {
				got = append(got, ints)
			})
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
