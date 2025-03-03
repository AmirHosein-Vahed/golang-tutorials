package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

// Methods with receiver `Point`
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func (p1 Point) IsEqual(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func main() {
	p := Point{2.0, 4.0}

	fmt.Println("Point : ", p)

	fmt.Println("Is Point p located above the line y = 1.0 ? : ", p.IsAbove(1))
	fmt.Println("Is Point p Equal to (2,4) ? : ", p.IsEqual(Point{2, 4}))
}
