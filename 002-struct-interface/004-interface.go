package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func PrintArea(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Printf("Area of the Shape (%T: %v) is: %v \n", s, s, s.Area())
	}
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	Radius float64
}

// This method has a receiver type of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {

	var a, b Shape

	a = Circle{1}
	b = Rectangle{1, 2}

	PrintArea(a, b)
}
