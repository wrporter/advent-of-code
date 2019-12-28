package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeck_Shuffle(t *testing.T) {
	tests := []struct {
		size       int
		techniques []string
		want       []int
	}{
		{
			10,
			[]string{
				"deal with increment 7",
				"deal into new stack",
				"deal into new stack",
			},
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			10,
			[]string{
				"cut 6",
				"deal with increment 7",
				"deal into new stack",
			},
			[]int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			10,
			[]string{
				"deal with increment 7",
				"deal with increment 9",
				"cut -2",
			},
			[]int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			10,
			[]string{
				"deal into new stack",
				"cut -2",
				"deal with increment 7",
				"cut 8",
				"cut -4",
				"deal with increment 7",
				"cut 3",
				"deal with increment 9",
				"deal with increment 3",
				"cut -1",
			},
			[]int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			deck := New(tt.size)
			deck.Shuffle(tt.techniques)
			if got := deck.Cards; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_ShuffleStack(t *testing.T) {
	tests := []struct {
		size int
		want []int
	}{
		{
			10,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			deck := New(tt.size)
			deck.ShuffleStack()
			if got := deck.Cards; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_ShuffleIncrement(t *testing.T) {
	tests := []struct {
		size      int
		increment int
		want      []int
	}{
		{
			10,
			3,
			[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
		{
			10,
			6,
			[]int{5, 0, 2, 0, 4, 0, 1, 0, 3, 0},
		},
		{
			10,
			7,
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			deck := New(tt.size)
			deck.ShuffleIncrement(tt.increment)
			if got := deck.Cards; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_ShuffleCut(t *testing.T) {
	tests := []struct {
		size   int
		amount int
		want   []int
	}{
		{
			10,
			10,
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			10,
			3,
			[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			10,
			6,
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
		{
			10,
			-10,
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			10,
			-4,
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
		{
			10,
			-7,
			[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			deck := New(tt.size)
			deck.ShuffleCut(tt.amount)
			if got := deck.Cards; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
