package ex2_5

import (
	"testing"
	"ch2/popcount"
	"ch2/ex2_3"
	"ch2/ex2_4"
)

func TestPopCountEx2_5 (t *testing.T) {
	if PopCount(15) != 4 {
		t.Error(`PopCount(15) != 4`)
	}
	if PopCount(389) != 4 {
		t.Error(`PopCount(389) != 4`)
	}
}

func TestPopCountEx2_4 (t *testing.T) {
	if ex2_4.PopCount(15) != 4 {
		t.Error(`PopCount(15) != 4`)
	}
	if ex2_4.PopCount(389) != 4 {
		t.Error(`PopCount(389) != 4`)
	}
}

func TestPopCountEx2_3 (t *testing.T) {
	if ex2_3.PopCount(15) != 4 {
		t.Error(`PopCount(15) != 4`)
	}
	if ex2_3.PopCount(389) != 4 {
		t.Error(`PopCount(389) != 4`)
	}
}

func TestPopCountOriginal (t *testing.T) {
	if popcount.PopCount(15) != 4 {
		t.Error(`PopCount(15) != 4`)
	}
	if popcount.PopCount(389) != 4 {
		t.Error(`PopCount(389) != 4`)
	}
}

func BenchmarkPopCountEx2_5Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountEx2_4Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex2_4.PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountEx2_3Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex2_3.PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountOriginalMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountEx2_5Rand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(7613792394324794754)
	}
}

func BenchmarkPopCountEx2_4Rand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex2_4.PopCount(7613792394324794754)
	}
}

func BenchmarkPopCountEx2_3Rand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex2_3.PopCount(7613792394324794754)
	}
}

func BenchmarkPopCountOriginalRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(7613792394324794754)
	}
}

/*
	For max uint64 this alorithm falls inbetween the loop over the prebuilt table and bitshifting each bit off the integer. At roughly 2x the loop over the 
	prebuilt table. However, for a random number that was fairly sparse with 1's
		2#110100110101001100111111101011111010001010011001101110110000010
	this method was fairly equivalent to looping over the prebuilt table. However, the original algorithm is still pretty superior.

	$ go test -bench=.
	BenchmarkPopCountEx2_5Max-8       	30000000	        53.7 ns/op
	BenchmarkPopCountEx2_4Max-8       	20000000	        84.1 ns/op
	BenchmarkPopCountEx2_3Max-8       	50000000	        26.6 ns/op
	BenchmarkPopCountOriginalMax-8    	2000000000	         0.79 ns/op
	BenchmarkPopCountEx2_5Rand-8      	50000000	        30.6 ns/op
	BenchmarkPopCountEx2_4Rand-8      	20000000	        81.1 ns/op
	BenchmarkPopCountEx2_3Rand-8      	50000000	        26.4 ns/op
	BenchmarkPopCountOriginalRand-8   	2000000000	         0.79 ns/op


*/
