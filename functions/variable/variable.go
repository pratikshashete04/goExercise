package main

import "fmt"

// declaration of global variable
var x int = 50

func main() {

	// declaration of local variables
	var x int = 11
	var y = 22
	z := 0

	fmt.Printf("Value of x in main(): %d\n", x)
	z = sum(x, y)
	fmt.Printf("Value of z in main(): %d\n", z)
}

func sum(x, y int) int {
	fmt.Printf("Value of x in sum(): %d\n", x)
	fmt.Printf("Value of y in sum(): %d\n", y)

	return x + y
}
