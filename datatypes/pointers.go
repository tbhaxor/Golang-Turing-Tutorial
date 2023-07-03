package main

import (
	"fmt"
)

func pointers() {
	fmt.Println("--- Demonstrating Pointer Variables ---")
	// Create an array of integers
	array := [3]int{1, 2, 3}
	fmt.Println(array)

	// Create a pointer to the array
	pointer := &array

	// Modify the second element of the array using the pointer
	// * is used to dereference the pointer
	(*pointer)[1] = 10

	// Print the modified array
	fmt.Println(array)

	// allocate new int type pointer
	num := new(int)
	*num = 10
	fmt.Println(num, *num)
}
