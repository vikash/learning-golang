package main

import "fmt"

func Numbers(c chan int) {
	for i := 1; ; i++ {
		c <- i
	}
	return
}

func EvenFilter(in chan int, out chan int) {
	for i := range in {
		if i%2 != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Numbers(ch)

	evenCh := make(chan int)
	go EvenFilter(ch, evenCh)

	for i := range evenCh {
		fmt.Println("Number: ", i)
		if i > 10 {
			break
		}
	}
}
