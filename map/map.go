package main

import "fmt"

func main() {

	// creating and initializing a map which is empty using var keyword
	var map_first map[int]int

	// checking the map is nil or not
	if map_first == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// creating and initializing a map
	// using map literals
	map_second := map[int]string{

		10: "English",
		20: "Hindi",
		30: "Marathi",
		40: "Spanish",
		50: "German",
	}

	// printing the map
	fmt.Println("The map is:", map_second)
}
