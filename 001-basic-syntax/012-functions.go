package main

import "fmt"

// Typical function
func add(a int, b int) int {
	return a + b
}

// Function with named return
func sub(a int, b int) (r int) {
	r = a - b
	return
}

// Function with multiple return values
func squares(a int, b int) (int, int) {
	return a * a, b * b
}

// Function with recursion
func factorial(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}

// Variadic function
func addAll(nums ...int) (total int) {

	for _, num := range nums {
		total += num
	}

	return
}

func main() {
	fmt.Println("2 + 6 = ", add(2, 6))
	fmt.Println("2 - 6 = ", sub(2, 6))

	sq2, sq6 := squares(2, 6)
	fmt.Println("Square of 2 and 6 are: ", sq2, sq6)

	fmt.Println("Factorial 10 is: ", factorial(10))

	fmt.Println(addAll(1, 2), addAll(1, 2, 3), addAll(10, 11, 12))

	nums := []int{1, 2, 3, 4}
	fmt.Println(addAll(nums...)) // Note the slice...
}
