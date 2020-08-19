package main

import "fmt"

type tank interface {
	// Methods
	TankArea() float64
	TankVolume() float64
}

type value struct {
	radius float64
	height float64
}

// implementing the methods of interface
func (v value) TankArea() float64 {

	return 2*v.radius*v.height + 2*3.14*v.radius*v.radius
}

func (v value) TankVolume() float64 {

	return 3.14 * v.radius * v.radius * v.height
}

func main() {

	// accessing elements of the tank interface
	var t tank
	t = value{12, 20}
	fmt.Println("The area of the tank:", t.TankArea())
	fmt.Println("The volume of the tank:", t.TankVolume())
}
