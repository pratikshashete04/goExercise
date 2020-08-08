package main

import (
	"fmt"
	"strings"
)

// variadic function
func joinstr(element ...string) string {
	return strings.Join(element, " ")
}

func main() {

	// calling variadic function in main function
	// zero argument
	fmt.Println(joinstr())

	// multiple arguments
	fmt.Println(joinstr("Floratechno", "solutions", "Pvt.", "Ltd.", "Pune."))
	fmt.Println(joinstr("F", "L", "O", "R", "A"))

	// pass a slice in variadic function
	element := []string{"Flora", "techno", "solutions"}
	fmt.Println(joinstr(element...))
}
