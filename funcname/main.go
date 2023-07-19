package main

import (
	"fmt"
	"os"

	// since this is imported before init, this package init() will be called firs

	"funcname/arithmetic"
	"funcname/otherpackage"
)

var i = 10

// called before main()
func init() {
	fmt.Println("Initialization function for main")
	i = 20
}

func main() {
	// call the exported function
	otherpackage.Greet(os.Args[1:])

	// this value will be from init()
	fmt.Println("Value of I", i)

	fmt.Println("-- Arithmetic Demo --")
	fmt.Println(arithmetic.Add(10, 20))
	fmt.Println(arithmetic.Sub(10, 20))
	fmt.Println(arithmetic.Prod(10, 20))
	fmt.Println(arithmetic.Div(10, 20))
	fmt.Println(arithmetic.Div(10, 0))
	fmt.Println(arithmetic.Do(3, 4))
	fmt.Println(arithmetic.Do(3, 0))
}
