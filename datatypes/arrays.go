package main

import "fmt"

func arrays() {
	fmt.Println("--- Demonstrating Arrays ---")

	// declare array of 5 elements of type int which are initilized to 0
	var numbers [5]int

	// print first element of the
	fmt.Println(numbers[0])

	// underscore variable are ignore from unused variable
	// proof all the elements are zero
	sum := 0
	for _, el := range numbers { // this can also be for i := 0; i < len(numbers); i++
		sum += el
	}
	fmt.Println(sum)

	// length of array, here it will be 5
	fmt.Println(len(numbers))

	// this will give compile time error, going beyond allocated area
	// numbers[5] = 10; // uncomment this to get error

	// set value at particular index, only int values are allowed
	numbers[2] = 100

	// print whole array at once, this is allowed
	fmt.Println(numbers)
}
