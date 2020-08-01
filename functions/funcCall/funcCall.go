package main

import "fmt"

func main() {

	// local variable definition
	var a int = 150
	var b int = 250
	var ret int

	// calling a function to get max value
	ret = max(a, b)
	fmt.Printf("The maximum value is: %d\n", ret)
}

// max function will return the maximum value between two numbers
func max(num1, num2 int) int {
	// local variable declaration
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
