package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
	i := 0

	defer fmt.Println(i)
	i++

   // Example 2:

	// multiple defer functions are put into a stack. They are popped during execution
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
