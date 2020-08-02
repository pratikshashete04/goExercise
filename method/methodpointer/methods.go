// method can accept both pointer and value.
package main

import "fmt"

// college structure
type college struct {
	name        string
	field       string
	departments int
}

// method with value receiver of college type
func (c college) info_new() {
	c.field = "Engineering & Technology"
	c.departments = 10
	fmt.Println("The field of the new college is:", c.field)
	fmt.Println("The number of departments in the new college are:", c.departments)
}

// method with a pointer receiver of college type
func (c *college) info_new1(cname string) {
	(*c).name = cname
}

func main() {

	// initialzing the values of college structure
	res := college{
		name:        "S.T.B.College",
		field:       "Engineering",
		departments: 8,
	}
	fmt.Println("The name of the old college was:", res.name)

	// calling the info_new1 method (pointer method) with value
	res.info_new1("R.P.College")
	fmt.Println("The name of the new college is:", res.name)
	fmt.Println("The field of the old college was:", res.field)

	// calling the info_new method (value method) with a pointer
	(&res).info_new()
	fmt.Println("The number of departments in the old college were:", res.departments)
}
