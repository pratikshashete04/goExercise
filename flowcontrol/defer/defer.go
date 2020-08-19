package main

import "fmt"

func multiply(x, y int) int {
	result := x * y
	fmt.Println("The result of multiplication is:", result)
	return 0
}

func do() {
	fmt.Println("defer in Go language")
}

func main() {

	// calling multiple function without defer
	defer multiply(20, 20)

	// calling do function
	do()

	// calling multiply function with defer keyword
	multiply(30, 30)
}
