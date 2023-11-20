package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 8
	w.Spokes = 20

	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	fmt.Printf("%#v\n", w)
	// Output:
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:42, Y:8}, Radius:5}, Spokes:20
}
