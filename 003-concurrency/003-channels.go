package main

import (
	"fmt"
	"time"
)

func main() {

	// Anonymous function as goroutine
	c := make(chan string)
	go func() {
		c <- "Hello Channel"
	}()

	fmt.Println(<-c)

	// Blocking for a goroutine to finish - communication using channel
	d := make(chan bool)
	fmt.Println("Starting  the work")
	go work(d)
	fmt.Println("Doing other things")
	// Blocking till work gets finished - a message to receive on channel d.
	<-d
}

func work(done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Work done. ")
	done <- true
}
