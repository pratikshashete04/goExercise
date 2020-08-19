package main

import (
	"fmt"
	"sync"
)

func first(wg *sync.WaitGroup) {
	defer wg.Done() // This increases counter by 1
	fmt.Println("*First step here*")
}

func second(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("*Second step here*")
}

func third(wg *sync.WaitGroup) {
	defer wg.Done()

	// checking condition by using if
	if true {
		fmt.Println("*This is third one step*")
	} else {
		fmt.Println("error occurred while execution!!")
	}
}

func steps() {
	wg := new(sync.WaitGroup)

	// The integer value in wg.Add() is the waitgroup counter
	// Here we are increasing the counter by 3, because we have 3 goroutines
	wg.Add(3)

	// calling the goroutines
	go first(wg)
	go second(wg)
	go third(wg)

	// This blocks the execution until it's counter becomes 0
	wg.Wait()
}

func main() {
	steps()
}
