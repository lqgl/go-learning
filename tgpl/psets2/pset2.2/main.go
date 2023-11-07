// Pset2.2 converts numeric command-line arguments in feet&meters and pounds&kilograms.
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Meters float64

type Pounds float64
type Kilograms float64

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pset2.2: %v\n", err)
			os.Exit(1)
		}
		f := Feet(t)
		m := Meters(t)
		p := Pounds(t)
		k := Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FToM(f), m, MtoF(m))
		fmt.Printf("%s = %s, %s = %s\n", p, PToK(p), k, KToP(k))
	}
}

func FToM(f Feet) Meters {
	return Meters(f * 0.3048)
}

func MtoF(m Meters) Feet {
	return Feet(m * 3.2808)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%gm", m)
}

func PToK(p Pounds) Kilograms {
	return Kilograms(p * 0.4536)
}

func KToP(k Kilograms) Pounds {
	return Pounds(k * 2.205)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%gkg", k)
}
