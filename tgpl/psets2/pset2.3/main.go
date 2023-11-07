// Pset2.3 Rewirte PopCount to use a loop instead of a single expression.
package main

// pc[i] is the population count of i.
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


// PopCount2 returns the population count (number of set bits)  of x.
func PopCount2(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>i*8)])
	}
	return sum
}
