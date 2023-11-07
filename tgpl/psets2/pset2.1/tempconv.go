// Pset2.1 processing temperatures in the Kelvin scale.
package main

import "fmt"

type Celsius float64
type Farenheit float64
type Kalvin float64

const (
	AbsoulteZeroC = -273.15
	BoilingC      = 100
	FrezzingC     = 0
)

// CToF converts Celsius temperature to Farenheit.
func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

// FToC converts Farenheit temperature to Celsius.
func FToC(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converts Celsius temperature to Kalvin.
func CToK(c Celsius) Kalvin {
	return Kalvin(c - AbsoulteZeroC)
}

// KToC converts Kalvin temperature to Celsius.
func KToC(k Kalvin) Celsius {
	return Celsius(k + AbsoulteZeroC)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
