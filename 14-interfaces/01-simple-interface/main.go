package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

// -------------------------------------
type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// --------------------------------------
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// --------------------------------------
func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("area  =", g.area())
	fmt.Println("perim =", g.perim())
}

func detectCircle(g Geometry) {
	if c, ok := g.(Circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func detectRect(g Geometry) {
	if c, ok := g.(Rect); ok {
		fmt.Printf("rectangle %f * %f\n", c.width, c.height)
	}
}

func main() {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

	detectRect(r)
	detectRect(c)
}
