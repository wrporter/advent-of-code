package fft

import (
	"fmt"
	"testing"
)

func TestApply(t *testing.T) {
	tests := []struct {
		signal string
		phases int
		want   string
	}{
		{
			"12345678",
			0,
			"12345678",
		},
		{
			"12345678",
			1,
			"48226158",
		},
		{
			"12345678",
			2,
			"34040438",
		},
		{
			"12345678",
			3,
			"03415518",
		},
		{
			"12345678",
			4,
			"01029498",
		},
		{
			"80871224585914546619083218645595",
			100,
			"24176176480919046114038763195595",
		},
		{
			"19617804207202209144916044189917",
			100,
			"73745418557257259149466599639917",
		},
		{
			"69317163492948606335995924319873",
			100,
			"52432133292998606880495974869873",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Apply(tt.signal, tt.phases); got != tt.want {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		signal string
		want   string
	}{
		{"03036732577212944063491565474664", "84462026"},
		{"02935109699940807407585447034323", "78725270"},
		{"03081770884921959731165446850517", "53553731"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := Decode(tt.signal); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
