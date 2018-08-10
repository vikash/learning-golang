package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		// () optional but {} required
		if i%2 == 0 {
			continue
		} else {
			fmt.Println("Odd: ", i)
		}
	}

	// If with short statement
	if x := 42; x < 0 {
		fmt.Println("x is negative")
	} else if x > 0 {
		fmt.Println("x is positive")
	} else {
		fmt.Println("x is zero")
	}
	// x is not defined here.
}
