package main

import "fmt"

// return a function which returns int
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// int_seq function is assigned to a variable
	next_int := int_seq()

	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())

	// declaring other variable
	next_int1 := int_seq()

	fmt.Println(next_int1())
	fmt.Println(next_int1())
	fmt.Println(next_int1())
	fmt.Println(next_int1())
}
