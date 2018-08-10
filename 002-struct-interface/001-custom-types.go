package main

import "fmt"

type Minute int
type Hour int

func main() {

	var m1 Minute = 10
	m2 := Minute(20) // Notice the type conversion

	// Supports all operations of underlying types
	fmt.Println("Total Minutes: ", m1+m2)

	// Can be compared with same named type or same underlying type
	fmt.Println(m1 > m2)
	fmt.Println(m1 > 12)

	// Can not be compared with different named type even if the underlying type is same.
	// Following commented line will not work, if uncommented.
	//fmt.Println(h1 > m1)
}
