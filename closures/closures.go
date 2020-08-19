// Que :Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

// makeFib function to create fibonacci series
func makeFib() func() int {

	fib_first := 0
	fib_second := 1

	// anonymous function
	return func() int {
		fib_second, fib_first = (fib_first + fib_second), fib_second
		return fib_first
	}
}

func main() {

	// makeFib function assigned to a variable
	genFib := makeFib()
	for i := 0; i < 10; i++ {
		fmt.Println(genFib())
	}
}
