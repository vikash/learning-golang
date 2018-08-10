package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func main() {

	var r1 Rectangle
	r1.width = 11
	r1.height = 12
	fmt.Println("r1:", r1)

	// Named struct fields. You can omit a field if you want - nil value will be taken as default
	r2 := Rectangle{width: 10, height: 20}
	fmt.Println("r2:", r2)

	// Implied names by sequence. You can not omit any field here.
	r3 := Rectangle{1, 2}
	fmt.Println("r3:", r3)

	// You can take a pointer as well.
	r4 := &Rectangle{width: 1}
	fmt.Println("r4:", r4, r4.width, r4.height)
}
