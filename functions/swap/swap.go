package main

import "fmt"

// swap function will swap the two strings
func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("Solutions", "Floratechno")
	fmt.Println(a, b)
}
