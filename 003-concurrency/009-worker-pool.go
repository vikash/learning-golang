package main

import "fmt"

func main() {

	work := make(chan int, 100)
	result := make(chan int, 100)

	// Let's start some workers
	for w := 1; w <= 3; w++ {
		go SquareWorker(w, work, result)
	}

	// Now we can keep sending them work.
	// Since work and result channels are buffered, we can queue upto 100 work and result

	// Let's send some work
	for j := 1; j <= 100; j++ {
		work <- j
	}
	close(work) // We can close the work channel, as we are not sending more jobs.

	// Let's get the results
	for a := 1; a <= 100; a++ {
		<-result
	}

	// Please note that if we don't get the results then our program will finish
	// before all work have been completed.
}

func SquareWorker(id int, work <-chan int, result chan<- int) {
	for w := range work {
		fmt.Printf("Worker %v working on %v \n", id, w)
		result <- w * w
		fmt.Printf("Worker %v finished %v \n", id, w)
	}
}
