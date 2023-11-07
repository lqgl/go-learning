// Pset2.4 PopCount counts bits by shfiting its argument through 64 bit positions.
package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount1 returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>1*8)] +
		pc[byte(x>>2*8)] +
		pc[byte(x>>3*8)] +
		pc[byte(x>>4*8)] +
		pc[byte(x>>5*8)] +
		pc[byte(x>>6*8)] +
		pc[byte(x>>7*8)])
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		if x&uint64(1) > 0 {
			sum++
		}
		x >>= 1
	}
	return sum
}
