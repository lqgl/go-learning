// Pset4.1 popCount of sum256.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	n := 0
	for _, v := range c1 {
		n += popCount(v)
	}
	fmt.Printf("popCount: %d\n", n)
}

func popCount(x byte) int {
	cnt := 0
	for x != 0 {
		x &= x - 1
		cnt++
	}
	return cnt
}
