package main

import "fmt"

// Notice the channel directions. Function and only send to c channel
func Numbers(c chan<- int) {
	for i := 1; ; i++ {
		c <- i
	}
	return
}

// Here function can only send to out channel and receive from in channel
func MultipleFilter(num int, in <-chan int, out chan<- int) {
	for i := range in {
		if i%num != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Numbers(ch)

	out := make(chan int)
	go MultipleFilter(4, ch, out)

	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}

}
