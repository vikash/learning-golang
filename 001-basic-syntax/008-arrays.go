package main

import "fmt"

func main() {

	var a [3]string
	fmt.Println("a:", a)

	// Set using index
	a[1] = "Hello"
	fmt.Println("a:", a, "a[1]: ", a[1]) // Get using index

	a[0] = "Hi!"
	fmt.Println("a:", a, " Length: ", len(a)) // Get Length

	// Declare and initialize together
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// Two dimensional Arrays
	var c [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("c:", c)
}
