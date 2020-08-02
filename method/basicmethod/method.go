package main

import "fmt"

// declaration of book struct
type book struct {
	author      string
	publication string
	pages       int
	prize       int
}

// method with a receiver of book type
func (b book) details() {
	fmt.Println("The name of the author is:", b.author)
	fmt.Println("The publication of the book is:", b.publication)
	fmt.Println("The number of pages in this book are:", b.pages)
	fmt.Println("The prize of this book is:", b.prize)
}

// main function
func main() {
	res := book{
		author:      "R.S.Agarwal",
		publication: "S.Chand Publishing",
		pages:       1500,
		prize:       250,
	}
	// calling the method in main
	res.details()
}
