package main

import (
	"fmt"
	"time"
)

func create() {
	fmt.Println("Creating goroutine!")
}

func main() {

	fmt.Println("Before calling goroutine")

	// creating goroutine
	go create()

	// schedule another goroutine
	time.Sleep(10 * time.Millisecond)

	fmt.Println("After calling goroutine")

}
