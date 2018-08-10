package main

import "fmt"

func counter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

func main() {

	aCounter := counter(10)
	bCounter := counter(100)

	var p = fmt.Println

	p(aCounter(), aCounter())
	p(bCounter(), bCounter())

	p(counter(50)())

}
