package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// Can take a pre-statement like  if/else
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	// break is not needed. fallthrough needs to be specified.

	// Cases can be expressions and switch condition is optional.
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Bye!")
	}

	// t is not defined here. So, we have to reinitialize.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Bye!")
	}
}
