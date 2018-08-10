package main

import "fmt"

// Built in associative data type. (Hash, Dictionary)

func main() {

	a := make(map[string]string)

	a["India"] = "New Delhi"
	a["Nepal"] = "Kathmandu"
	a["Pakistan"] = "Islamabad"

	fmt.Println("a: ", a)

	b := a
	delete(b, "Pakistan")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	// Declare and initialize
	n := map[string]int{
		"A": 12,
		"B": 16,
	}
	fmt.Println("n: ", n)
}
