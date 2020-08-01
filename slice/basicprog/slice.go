package main

import "fmt"

func main() {

	// Creating an array
	name := [7]string{"Floratechno", "solutions", "Pvt.", "Ltd.", "Pune,", "Maharashtra,", "India."}

	fmt.Println("Name of the company:", name)

	// Creating a slice
	var myslice []string = name[0:4]

	// Display the slice
	fmt.Println("Slice :", myslice)

	// Display the length of the slice
	fmt.Println("The length of the slice is:", len(myslice))

	// Display the capacity of the slice
	fmt.Println("The capacity of the slice is:", cap(myslice))
}
