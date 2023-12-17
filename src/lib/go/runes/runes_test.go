package runes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		values []rune
		amount int
		want   []rune
	}{
		{[]rune("12345"), 1, []rune("51234")},
		{[]rune("12345"), -1, []rune("23451")},
		{[]rune("12345"), 3, []rune("34512")},
		{[]rune("12345"), -3, []rune("45123")},
		{[]rune("edfgbach"), 3, []rune("achedfgb")},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Rotate(tt.values, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}
