package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Point.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	p1 := Point{1, 2}
	q1 := Point{4, 6}
	distanceFromP := p1.Distance   // method value
	fmt.Println(distanceFromP(q1)) // "5"
	scaleP := p1.ScaleBy           // method value
	scaleP(2)

	distance := Point.Distance    // method expression
	fmt.Println(distance(p1, q1)) // "5"
	scale := (*Point).ScaleBy     // method expression
	scale(&p1, 2)
	fmt.Println(p1)           // "{4 8}"
	fmt.Printf("%T\n", scale) // func(*main.Point, float64)
}
