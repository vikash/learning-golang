package main

import "fmt"

func Numbers(c chan int) {
	for i := 1; ; i++ {
		c <- i
	}
	return
}

func main() {
	ch := make(chan int)
	go Numbers(ch)

	for i := range ch {
		fmt.Println("Number: ", i)
		if i == 10 {
			break
		}
	}
}
