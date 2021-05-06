package main

import (
	"fmt"
)

// Slice is a triple of pointer to first element, length and capacity (length ≤ capacity).
// The append built-in function appends elements to the end of a slice. If it has sufficient capacity,
// the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated.
// https://blog.allegro.tech/2017/07/golang-slices-gotcha.html
// https://blog.golang.org/slices-intro

func a() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)
	fmt.Println(y, z)
}

func b() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}

// Be careful while returning slice from a function. As the original array will not get garbage collected till the slice is in scope.

func main() {
	a()
	b()
}
