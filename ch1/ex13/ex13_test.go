package ex13

import (
	"testing"
)

func BenchmarkEchoNoJoin(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		EchoNoJoin([]string{"this", "is", "a", "test", "string", "with", "more", "words", "in", "it"})
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		EchoJoin([]string{"this", "is", "a", "test", "string", "with", "more", "words", "in", "it"})
	}
}

//$ go test -bench=.
//testing: warning: no tests to run
//BenchmarkEchoNoJoin-8   	 1000000	      1251 ns/op
//BenchmarkEchoJoin-8     	 5000000	       381 ns/op
