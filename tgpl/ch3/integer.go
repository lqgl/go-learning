// Integer: some operations of integers.
package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the uinon {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	/*
		var apples int32 = 1
		var oranges int16 = 2
		var compote int = apples + oranges // compile error
	*/

	f := 3.141
	z := int(f)
	fmt.Println(f, z) // "3.141 3"

	f = 1.9
	fmt.Println(int(f)) // "1"

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 438 666 0666
	e := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", e) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 22269 国 '国'
	fmt.Printf("%d %[1]q\n", newline)       // 10 '\n'

	var fl float32 = 16777216 // 1 << 24
	fmt.Println(fl == fl+1)   // true
}
