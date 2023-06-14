package main

import (
	"fmt"
	"os"
)

// main is like any other programming language,
// is the application's entry point.
// Exactly 1 definition of this function is required in main package
func main() {
	// check user has provided any
	// https://pkg.go.dev/os@go1.20.5#Args
	// https://pkg.go.dev/builtin#len
	if len(os.Args) > 1 {
		// take the second element of the array and print hello to user
		// https://pkg.go.dev/fmt@go1.20.5#Printf
		fmt.Printf("Hello, %s\n", os.Args[1])
	} else {
		// if no argument is provided, print message with new line
		// https://pkg.go.dev/fmt@go1.20.5#Println
		fmt.Println("Hello Guest")
	}
}
