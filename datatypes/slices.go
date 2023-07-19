package main

import "fmt"

func slices() {
	fmt.Println("--- Demonstrating Slices ---")

	var numbers [5]int // just for demo

	// initialize slice from the length
	var slice1 []int = numbers[:3]

	// print length and capacity of slice1
	fmt.Println(len(slice1), cap(slice1))

	// convert whole array to slice
	var slice2 []int = numbers[:]

	// print length and capacity of slice2
	fmt.Println(len(slice2), cap(slice2))

	// you can have slice without any array like this,
	// if you remove the right hand side, that will be length 0 capacity 0 slice
	var slice3 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// print length and capacity of slice3
	fmt.Println(len(slice3), cap(slice3))

	// print slice before append
	fmt.Printf("%v\t%p\n", slice3, &slice3)

	// append element to the last of slice
	slice3 = append(slice3, 10)

	// print slice after append
	fmt.Printf("%v\t%p\n", slice3, &slice3)

	// print length and capacity of slice3
	fmt.Println(len(slice3), cap(slice3))

	deleteIndex := 3

	// deleting third index from the slice
	slice3 = append(slice3[:3], slice3[deleteIndex+1:]...)

	// print slice after delete
	fmt.Printf("%v\t%p\n", slice3, &slice3)

	// print length and capacity of slice3
	fmt.Println(len(slice3), cap(slice3))

	// create slice with different length and capacity
	var slice4 = make([]int, 10, 20)

	// print length and capacity of slice2
	fmt.Println(len(slice4), cap(slice4))

	// this will iterate until length is reached

	for idx := range slice4 {
		fmt.Printf("%d ", idx)
	}
	fmt.Println()
}
