// Exercise 1.3 - Benchmark the difference in running time between our potentially inefficient versions and one that uses strings.Join
package ex13

import (
	"strings"
)

func EchoNoJoin(str []string) {
	s, sep := "", ""
	for _, arg := range str {
		s += sep + arg
		sep = " "
	}
}

func EchoJoin(s []string) {
	strings.Join(s, " ")
}