// version of popcount that shifts its argument through 64 bit positions testing the right most bit each time
package ex2_4

func PopCount(x uint64) int {
	var sum int
	for i := 0; i < 64; i++ {
		sum += int(x&1)
		x = x>>1
	}
	return sum
}