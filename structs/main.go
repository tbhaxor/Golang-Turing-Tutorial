package main

import (
	"fmt"
)

// define struct with two field
type Person struct {
	name string
	age  uint
}

// this a method attached to each instance of struct
func (p *Person) whoami() {
	fmt.Printf("Hi! I am %s and %d years old\n", p.name, p.age)
}

// define struct with embedding Person struct
type Employee struct {
	Person  // embedding happens here
	email   string
	website string
}

func main() {
	// initialize object of Person struct with values
	person := &Person{name: "Gurkirat Singh", age: 21}

	// print entire struct
	fmt.Printf("%+v\n", person)

	// print selective fields of struct
	fmt.Printf("Name=%s\t\tAge=%d\n", person.name, person.age)

	// change value
	person.name = "Amit"
	fmt.Printf("%+v\n", person)

	// call method of person instance
	person.whoami()

	// instantiate Employee struct and pass person information
	// since person was point, it needs to be dereferenced
	employee := &Employee{Person: *person,
		email:   "info@tbhaxor.com",
		website: "https://tbhaxor.com",
	}
	fmt.Printf("%+v\n", employee)

	// access value of embedded struct directly
	// from employee object
	employee.name = "Amit Singh"
	fmt.Printf("%+v\n", employee)

	// methods also gets embedded
	employee.whoami()
}
