package main

import "fmt"

func main() {

	a := 10
	b := &a
	fmt.Println(a, b, *b)
	fmt.Printf("%T %T %T\n", a, b, *b)

	// Function can take pointer as parameter or can return a pointer
	x := 10
	doubleIt(&x)
	fmt.Println(x)  // NOTE:Go does not support pointer arithmetic like C

}

func doubleIt(a *int) {
	*a = *a * 2
}
