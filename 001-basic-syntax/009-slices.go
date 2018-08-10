package main

import "fmt"

// Slices are similar to arrays but no fixed size and can support mutation and related operations

func main() {

	// Declare First
	var a []string
	fmt.Println(a, len(a))

	// Use make to initialize
	var b = make([]string, 3)
	fmt.Println(b, len(b))

	// Shorthand Notation to declare
	c := []string{"Hi", "Hello", "Hola", "Haha", "Hehe"}
	fmt.Println(c, len(c))

	// Slice operator to get a slice of slice
	d := c[1:]
	fmt.Println("d: ", d, len(d))

	e := c[2:4]
	fmt.Println("e:", e, len(e))

	// Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
	e[0] = "Test"
	fmt.Println("c", c, len(c))

	// Copy function to copy it.
	x := make([]string, 2)
	copy(x, c)
	fmt.Println("x", x, len(x))
	x[0] = "Test"
	fmt.Println("x", x, len(x))
	fmt.Println("c", c, len(c))

	// Append
	i := []int{1, 2, 3}
	i = append(i, 4)
	fmt.Println("i: ", i)

}
