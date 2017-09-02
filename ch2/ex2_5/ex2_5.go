package ex2_5

func PopCount(x uint64) int {
	var c int
	for ; x > 0; x &= x-1 {
		c++
	}
	return c
}