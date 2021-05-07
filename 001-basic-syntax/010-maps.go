package main

import "fmt"

// Built in associative data type. (Hash, Dictionary)

func main() {

	a := make(map[string]string) // zero value of a map s nil. Without make writing to a map will lead to panics

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

	// you can check if a key exists in a map or not
	val, ok := a["Canada"] // a["Canada"] can either return one value or two. Second value is a bool
	fmt.Println(val, ok)
}
