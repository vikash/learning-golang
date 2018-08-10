package main

import (
	"errors"
	"fmt"
	"math"
)

type NegativeNumber struct {
	num int
}

func (n NegativeNumber) Error() string {
	return fmt.Sprintf("The number '%v' is negative", n.num)
}

func sqrt(n int) (float64, error) {
	if n < 0 {
		return 0.0, NegativeNumber{n}
	}

	return math.Sqrt(float64(n)), nil
}

func main() {

	a, err := sqrt(2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("a:", a)
	}

	b, err := sqrt(-22)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("b:", b)
	}

	// we can also create error like this
	e := errors.New("This is an error")
	fmt.Printf("%T %v\n", e, e)

}
