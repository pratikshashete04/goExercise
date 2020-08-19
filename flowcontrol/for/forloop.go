package main

import (
	"fmt"
)

func main() {

	// ========================== single condition =================== //
	s := 1
	for s <= 8 {
		fmt.Println(s)
		s = s + 1
	}

	// ============= simple loop with initial/condition/post ========== //
	addition := 0

	for i := 0; i < 500; i++ {
		addition += i
	}
	fmt.Println("Sum is:", addition)

	// ====================== "continue" in for loop ======================== //

	for a := 0; a <= 15; a++ {
		if a%3 == 0 {
			continue
		}
		fmt.Println(a)
	}

	// ====================== for with "break" ============================= //

	for {
		fmt.Println("For loop")
		break
	}
}
