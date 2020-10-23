package main

import (
	"fmt"
	"math"
)

// Shape is an interface
type Shape interface {
	Area() float64
}

// Circle is a struct
type Circle struct {
	x, y, radius float64
}

// Rectangle is a struct
type Rectangle struct {
	width, height float64
}

// Area is a value receiver function
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Area is a value receiver function
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// GetArea is a value receiver function
func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	c := Circle{0, 0, 10}
	r := Rectangle{10, 10}

	fmt.Printf("Circle Area: %f\n", GetArea(c))
	fmt.Printf("Rectangle Area: %f\n", GetArea(r))
}
