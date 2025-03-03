package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

func IsAboveFunc(p Point, y float64) bool {
	return p.Y > y
}

/*
Compare the above function with the corresponding method
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}
*/

func IsEqualFunc(p1 Point, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

/*
Compare the above function with the corresponding method
func (p1 Point) IsEqual(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
*/

func main() {
	p := Point{2.5, -3.0}

	fmt.Println("Point : ", p)

	fmt.Println("Is Point p located above the line y = 1.0 ? : ", IsAboveFunc(p, 1))
	fmt.Println("Is Point p Equal to (2,4) ? : ", IsEqualFunc(p, Point{2, 4}))
}
