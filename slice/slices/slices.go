package main

import "fmt"

func main() {

	myslice := []int{2, 3, 5, 7, 11, 13}
	printSlice(myslice)

	// Slice the slice to give it zero length.
	myslice = myslice[:0]
	printSlice(myslice)

	// Extend the length of slice.
	myslice = myslice[:5]
	printSlice(myslice)

	// Drop slice's first three values.
	myslice = myslice[3:]
	printSlice(myslice)
}

func printSlice(myslice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(myslice), cap(myslice), myslice)
}
