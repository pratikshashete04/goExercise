package main

import "fmt"

func main() {

	// ================ simple if condition =========== //
	if india := 500; india == 500 {
		fmt.Println("Maharashtra")
	}

	// ================ if statement ===================== //
	var india = "Maharashtra"
	country := true
	if country {
		fmt.Println(india)
	}

	// ==================== if... else statement ========== //
	value := 1000
	if value == 5000 {
		fmt.Println("Maharashtra")
	} else {
		fmt.Println("Gujrat")
	}

	// =================== if...else if...else statement ====== //
	number := 800
	if number == 700 {
		fmt.Println("Maharashtra")
	} else if number == 800 {
		fmt.Println("Goa")
	} else {
		fmt.Println("Rajasthan")
	}
}
