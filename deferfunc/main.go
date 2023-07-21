package main

import (
	"fmt"
	"os"
)

func main() {
	// defer function 1
	// order of execution 3
	defer func() {
		fmt.Println("world")
	}()

	// defer function 2
	// order of execution 2
	defer fmt.Println("hello")

	// open file
	file, _ := os.Open("somefile.txt")

	/**
	Do some changes in the file
	*/

	// close the file as function ends
	// defer function 3
	// order of execution 1
	defer file.Close()
}
