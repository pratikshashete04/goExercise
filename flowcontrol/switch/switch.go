package main

import "fmt"

func main() {
	// =============== expression switch======================== //
	var month int = 10

	// switch statement without an optional statement and expression
	switch {
	case month == 1:
		fmt.Println("January")
	case month == 2:
		fmt.Println("February")
	case month == 3:
		fmt.Println("March")
	case month == 4:
		fmt.Println("April")
	case month == 5:
		fmt.Println("May")
	case month == 6:
		fmt.Println("June")
	case month == 7:
		fmt.Println("July")
	case month == 8:
		fmt.Println("August")
	case month == 9:
		fmt.Println("September")
	case month == 10:
		fmt.Println("October")
	case month == 11:
		fmt.Println("November")
	case month == 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid month")
	}
	testType(100)
}

// =================== type switch ================== //

func testType(value interface{}) {

	switch value_new := value.(type) {
	case int:
		fmt.Println("The value is of integer type")
	case float64:
		fmt.Println("The value is of float64 type")
	case bool:
		fmt.Println("The value is of boolean type")
	default:
		fmt.Printf("The value is of type: %T!\n", value_new)
	}
}
