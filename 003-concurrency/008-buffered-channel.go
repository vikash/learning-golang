package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1 // This does not wait for someone to receive as the channel has a buffer size of 2.

	fmt.Printf("Length: %v Capacity: %v \n", len(ch), cap(ch))

	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Printf("Length: %v Capacity: %v \n", len(ch), cap(ch))

	// Sending to buffered channel blocks after length equals capacity

}
