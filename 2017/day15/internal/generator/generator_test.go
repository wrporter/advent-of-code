package generator

import "testing"

func BenchmarkCountSynchronously(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountSynchronously()
	}
}

func BenchmarkCountWithChannels(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountWithChannels()
	}
}
