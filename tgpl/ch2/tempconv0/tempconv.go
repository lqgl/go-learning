// Package tempconv performs Celsius and Farenheit temperature computations.
// package tempconv
package main

import (
	"fmt"
)

type Celsius float64
type Farenheit float64

const (
	AbsoulteZeroC Celsius = -273.15
	FrezzingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

func FToC(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {

	fmt.Printf("%g\n", BoilingC-FrezzingC) // 100
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FrezzingC)) // 180
	// fmt.Printf("%g\n", boilingF-FrezzingC) // compile error: mismatched types

	var c Celsius
	var f Farenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // compile error: mismatched types
	fmt.Println(c == Celsius(f)) // true

	c = FToC(212.0)
	fmt.Println(c.String()) // 100°C
	fmt.Printf("%v\n", c)   // 100°C; no need to call String explicitly
	fmt.Printf("%s\n", c)   // 100°C
	fmt.Println(c)          // 100°C
	fmt.Printf("%g\n", c)   // 100; does not call String
	fmt.Println(float64(c)) // 100; does not call String
}
