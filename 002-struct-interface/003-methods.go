package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

// This method has a receiver type of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// This method has a receiver type of *Circle
// Use pointer receiver - to avoid copying large data on method call OR to mutate the struct
func (c *Circle) inflate(n float64) {
	c.Radius *= n
}

func main() {
	c := Circle{1}
	fmt.Println("Circle:", c, "Area: ", c.Area())

	c.inflate(2)
	fmt.Println("Area: ", c.Area())
}
