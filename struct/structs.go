package main

import "fmt"

// Defining a struct type
type Company struct {
	EmployeeName string
	MobNo        int
	Salary       int
}

func main() {

	// Declaring a variable of a struct type
	var employee Company
	fmt.Println(employee)

	// Declaring and initializing a struct using a struct literal
	employee1 := Company{"Ashweeni", 9890843929, 25000}

	fmt.Println("Employee1: ", employee1)

	// Naming fields while initializing a struct
	employee2 := Company{
		EmployeeName: "Pratiksha",
		MobNo:        8983561702,
		Salary:       20000,
	}
	fmt.Println("Employee2: ", employee2)

	// Uninitialized fields are set to their corresponding zero-value
	employee3 := Company{
		EmployeeName: "Supriya",
	}
	fmt.Println("Employee3: ", employee3)
}
