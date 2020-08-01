package main

import "fmt"

// add function will add the parameters
// it will return the integer value
func add(a, b int) int {
	return a + b
}

// remove function will substract the values of params
// it will return the integer value
func remove(a, b, c int) int {
	return a - b - c
}

// main is the main function which will be execute first
// We can call other functions from program in this main function
func main() {
	res := add(10, 50)
	fmt.Println("The addition is:", res)

	res = remove(100, 65, 5)
	fmt.Println("The substraction is:", res)
}
