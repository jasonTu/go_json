package main

import (
	"fmt"
	"image/color"
	"go_json/gopl/ch6/geometry"
	"go_json/gopl/ch6/coloredpoint"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))    // "5", function call
	fmt.Println(p.Distance(q))    // "5", method call

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())    // "12"

	var cp coloredpoint.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)    // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y)    // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var cp2 = coloredpoint.ColoredPoint{geometry.Point{1, 1}, red}
	var cq2 = coloredpoint.ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(cp2.Distance(cq2.Point))    // "5"

	// Method Values and Expressions

	p3 := geometry.Point{1, 2}
	q3 := geometry.Point{4, 6}

	distanceFromP := p3.Distance    // method value
	fmt.Println(distanceFromP(q3))

	var origin geometry.Point    // {0, 0}
	fmt.Println(distanceFromP(origin))    // 2.23606797749979

	distance := geometry.Point.Distance    // method expression
	fmt.Println(distance(p3, q3))    // "5"

	fmt.Printf("%T\n", distance)    // func(geometry.Point, geometry.Point) float64
	fmt.Printf("%T\n", distanceFromP)    // func(geometry.Point) float64
}
