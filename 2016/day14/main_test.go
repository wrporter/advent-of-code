package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getFirstTriplet(t *testing.T) {
	tests := []struct {
		stream string
		want   string
	}{
		{"aaa", "aaa"},
		{"123aaa", "aaa"},
		{"123aaaccc", "aaa"},
		{"2df6e9378c3c53abed6d3508b6285fff", "fff"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getFirstTriple(tt.stream); got != tt.want {
				t.Errorf("getFirstTriple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getQuints(t *testing.T) {
	tests := []struct {
		stream string
		want   map[string]bool
	}{
		{"aaaaa", map[string]bool{"aaaaa": true}},
		{"aaaaa2bbbbb", map[string]bool{"aaaaa": true, "bbbbb": true}},
		{"aaaaa2bbbbb3ccccc", map[string]bool{"aaaaa": true, "bbbbb": true, "ccccc": true}},
		{"aaaaaaaaaa", map[string]bool{"aaaaa": true}},
		{"aaaa", map[string]bool{}},
		{"2e559978fffff9ac9c9012eb764c6391", map[string]bool{"fffff": true}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getQuints(tt.stream); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getQuints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHash2017(t *testing.T) {
	tests := []struct {
		salt  string
		index int
		want  string
	}{
		{"abc", 0, "a107ff634856bb300138cac6568c0f24"},
		{"abc", 22551, "2df6e9378c3c53abed6d3508b6285fff"},
		{"abc", 22859, "2e559978fffff9ac9c9012eb764c6391"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := getStretchedHash(tt.salt, tt.index); got != tt.want {
				t.Errorf("getStretchedHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
