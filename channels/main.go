package main

import "fmt"

func sum(a, b int, result chan int) {
	result <- a + b
}

func main() {
	// create a channel of type int.
	c := make(chan int)

	// call goroutine using anonymous function.
	go func() {
		// send a value to the channel.
		c <- 10
	}()

	// receive the value from the channel.
	v := <-c

	// print the value.
	fmt.Println(v)

	// pass the channel to function
	go sum(10, 20, c)

	v = <-c
	fmt.Println(v)

	// buffered channel with a capacity of 3
	ch := make(chan int, 3)

	go func() {
		// send until the buffer is full
		// (non-blocking until it is full)
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	// receive as long as there's data in the buffer
	result1 := <-ch // No block
	result2 := <-ch // No block
	result3 := <-ch // No block

	fmt.Printf("%d\n%d\n%d\n", result1, result2, result3) // Output: 1 2 3
}
