package ex2_3

import (
	"ch2/popcount"
	"testing"
)

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(256)
	}
}

func BenchmarkPopCountNoLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(256)
	}
}

/*

$ go test -bench=.
testing: warning: no tests to run
BenchmarkPopCountLoop-8     	50000000	        26.8 ns/op
BenchmarkPopCountNoLoop-8   	2000000000	         0.81 ns/op
PASS

Not looping is much faster.

*/