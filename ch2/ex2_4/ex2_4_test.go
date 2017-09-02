package ex2_4

import (
	"testing"
	"ch2/popcount"
	"ch2/ex2_3"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(256)
	}
}

func BenchmarkPopCountEx2_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex2_3.PopCount(256)
	}
}

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(256)
	}
}