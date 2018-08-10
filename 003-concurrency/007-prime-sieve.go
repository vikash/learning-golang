package main

import "fmt"

func Numbers(c chan<- int) {
	for i := 2; ; i++ {
		c <- i
	}
	return
}

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

	for i := 0; i < 1000; i++ {
		prime := <-ch
		fmt.Println(prime)
		out := make(chan int)
		go MultipleFilter(prime, ch, out)
		ch = out
	}

}
