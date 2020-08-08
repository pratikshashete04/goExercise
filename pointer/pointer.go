package main

import "fmt"

func main() {

	var hp_laptop = 45000

	// taking a pointer variable without specifying the type
	var dell_laptop = &hp_laptop

	fmt.Println("The value stored in hp_laptop before changing is:", hp_laptop)
	fmt.Println("The address of hp_laptop =", &hp_laptop)
	fmt.Println("The value stored in pointer variable dell_laptop is:", dell_laptop)

	// using * operator before the pointer variable to access the value stored at the variable at which it is pointing
	fmt.Println("The value stored in hp_laptop(*dell_laptop) before changing is:", *dell_laptop)

	// changing the value of hp_laptop
	// by assigning the new value to the pointer variable
	*dell_laptop = 50000

	// printing the changed value
	fmt.Println("The value stored in hp_laptop(*dell_laptop) after changing is:", hp_laptop)
}
