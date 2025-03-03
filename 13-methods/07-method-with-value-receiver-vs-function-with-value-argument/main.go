package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func IsAboveFunc(p Point, y float64) bool {
	return p.Y > y
}

func main() {
	p := Point{0, 4}
	ptr := &p

	fmt.Println("Point p = ", p)

	// Calling a Method with Value receiver
	p.IsAbove(1)
	ptr.IsAbove(1)

	// Calling a Function with a Value argument
	IsAboveFunc(p, 1)
	IsAboveFunc(*ptr, 1)
}
