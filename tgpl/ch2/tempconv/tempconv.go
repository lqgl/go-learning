// Package tempconv performs Celsuis and Farenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoulteZeroC Celsius = -273.15
	FrezzingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF converts Celsius temperature to Farenheit.
func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

// FToC converts Farenheit temperature to Celsius.
func FToC(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
