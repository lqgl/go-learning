// Pset2.5 PopCount by useing x&(x-1).
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

// PopCount2 returns the population count(number of set bits) of x.
func PopCount2(x uint64) int {
	sum := 0
	for x != 0 {
		x &= x - 1
		sum++
	}
	return sum
}
