package main

import (
	"fmt"
)

func main() {
	// bool
	a := true
	fmt.Printf("a: %v %T\n", a, a)

	// string
	b := "Learning GoLang"
	fmt.Printf("b: %v %T\n", b, b)

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// float32 float64
	var (
		c int     = -20
		d uint    = 10
		e float32 = 1.23
		f float64 = 1.012e12
		g float64 = 123e-12
		h int     = 1e10
	)
	fmt.Printf("c: %v %T\n", c, c)
	fmt.Printf("d: %v %T\n", d, d)
	fmt.Printf("e: %v %T\n", e, e)
	fmt.Printf("f: %v %T\n", f, f)
	fmt.Printf("g: %v %T\n", g, g)
	fmt.Printf("h: %v %T\n", h, h)

	// complex64 complex128
	var i complex128 = -5 + 12i
	fmt.Printf("i: %v %T\n", i, i)

	// Byte - alias for uint8
	var x byte = 'A'
	fmt.Printf("x: %v %T\n", x, x)

	// Rune - alias for int32
	var y rune = 'क' // This will overflow, if byte.
	fmt.Printf("y: %v %T\n", y, y)

}
