package main

import "fmt"

func main() {

	// Declare first, initialize later
	var a int
	fmt.Println("a = ", a)

	a = 10
	fmt.Println("a = ", a)

	// Declare and initialize together
	var b int = 10
	fmt.Println("b = ", b)

	// Short Symbols
	c := 10
	fmt.Println("c = ", c)

	// Multiple Variables
	var x, y string = "Hello", "World!"
	fmt.Println(x, y)

	// Multiple Shorthand declaration and initialization
	p, q := "Hello", 10
	fmt.Println(p, q)

}
