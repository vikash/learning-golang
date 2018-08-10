package main

import "fmt"

func main() {

	arr := [5]int{11, 22, 33, 44, 55}
	slice := []string{"A", "B", "C"}
	m := map[string]string{"You": "Me", "He": "She"}

	fmt.Println("Printing Array: ")
	for k, v := range arr {
		fmt.Println(k, v)
	}

	fmt.Println("Printing Slice: ")
	for _, v := range slice { // We can ignore the values we are not interested in using _
		fmt.Println(v)
	}

	fmt.Println("Printing Map: ")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

}
