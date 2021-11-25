package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		values []rune
		start  int
		end    int
		want   []rune
	}{
		{[]rune("0123456789"), 3, 8, []rune("0128765439")},
		{[]rune("abcdefgh"), 0, 1, []rune("bacdefgh")},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := reverse(tt.values, tt.start, tt.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_move(t *testing.T) {
	tests := []struct {
		values []rune
		from   int
		to     int
		want   []rune
	}{
		{[]rune("abcdefgh"), 5, 7, []rune("abcdeghf")},
		{[]rune("abcdefgh"), 5, 2, []rune("abfcdegh")},
		{[]rune("abcdefgh"), 7, 1, []rune("ahbcdefg")},
		{[]rune("abcdefgh"), 0, 5, []rune("bcdefagh")},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := move(tt.values, tt.from, tt.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
