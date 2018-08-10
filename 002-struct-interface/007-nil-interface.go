package main

import (
	"fmt"
)

func printDouble(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println(v * v)
	case float64:
		fmt.Println(v * v)
	case string:
		fmt.Println(v, v)
	case []int:
		for _, i := range v {
			printDouble(i)
		}
	default:
		fmt.Printf("Can't Double %T\n", v)
	}
}

func main() {
	printDouble(2)
	printDouble(2.3)
	printDouble("Hi!")
	printDouble(map[string]int{})
	printDouble([]int{2, 3, 4, 5, 6})
}
