package main

import "fmt"

func main() {
	// range operator can be applied on three types: arrays, slices,maps
	arr := [5]int{11, 22, 33, 44, 55}
	slice := []string{"A", "B", "C"}
	m := map[string]string{"You": "Me", "He": "She"}

	fmt.Println("Printing Array: ")
	for k, v := range arr { // k is the index and v is a copy of that element at index 'k'.
		fmt.Println(k, v)
	}

	fmt.Println("Printing Slice: ")
	for _, v := range slice { // We can ignore the values we are not interested in using _
		fmt.Println(v)
	}

	for k := range slice { // if v needs to be ignored, _ is not needed. it can be omitted
		fmt.Println(k)
	}

	fmt.Println("Printing Map: ")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

}
