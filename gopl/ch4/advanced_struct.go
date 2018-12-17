package main

import "fmt"

type Point struct {
    X, Y int
}

// Anoymous fields: Point
type Circle struct {
    Point
    Radius int
}

// Anoymous fields: Circle
type Wheel struct {
    Circle
    Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 10}, 15}
    /* Or
    w = WHeel{
        Circle{
            Point(X: 8, Y:8),
            Radius: 10,
        },
        Spokes: 15,
    }
    */
    fmt.Printf("%#v\n", w)
}

/*
Output:
main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:10}, Spokes:15}
 */
