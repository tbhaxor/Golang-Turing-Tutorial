package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type Employee struct {
	Person
	Company string `json:"company"`
}

// parseJson can accept body of type string but can return pointer to "any" type data
// because of T used as generics here
func parseJson[T any](body string) *T {
	s := new(T)

	decoder := json.NewDecoder(bytes.NewBufferString(body))

	if err := decoder.Decode(s); err != nil {
		panic(err)
	} else {
		return s
	}
}

const (
	jsonString = `{"name": "Gurkirat Singh", "age": 21, "company": "Turing"}`
)

func main() {
	// type specified in the [ and ] will be replaced as the type in the function
	// on hover you will see it returns *Person
	person := parseJson[Person](jsonString)

	// here, on hover you will see it returns *Employee
	employee := parseJson[Employee](jsonString)

	fmt.Printf("%+v\n", person)
	fmt.Printf("%+v\n", employee)

}
