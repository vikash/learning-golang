package main

import "fmt"

func main() {

	// Typical For loop - initial/condition/after
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Like a while loop
	j := 2
	for j <= 10 {
		fmt.Println(j)
		j += 2
	}

	// Infinite loop
	k := 1
	for {
		fmt.Println(k)
		if k == 10 {
			break
		}
		k++
	}
}
