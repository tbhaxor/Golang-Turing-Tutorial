package main

import (
	"os"

	"funcname/otherpackage"
)

func main() {
	// call the exported function
	otherpackage.Greet(os.Args[1:])
}
