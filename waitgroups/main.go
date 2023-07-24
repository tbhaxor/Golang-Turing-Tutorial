package main

import (
	"fmt"
	"sync"
	"time"
)

func helloWorld(wg *sync.WaitGroup, sleepTime time.Duration) {
	// inform goroutine is completed
	// when this function is over
	defer wg.Done()

	fmt.Println("Hello World!!")
	time.Sleep(sleepTime)
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Executing")
	// it will wait for 2 goroutines to be completed
	wg.Add(2)
	go helloWorld(wg, 2*time.Second)
	go helloWorld(wg, 1*time.Second)

	// // second way to do the same thing
	// wg.Add(1)
	// go helloWorld(wg, 2*time.Second)
	// wg.Add(1)
	// go helloWorld(wg, 1*time.Second)

	fmt.Println("Submitted 2 goroutines")

	// waits until all routines are done
	wg.Wait()

	fmt.Println("All Done")
}
