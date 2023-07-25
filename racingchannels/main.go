package main

import (
	"fmt"
	"time"
)

func waitForMe(d time.Duration, done chan bool) {
	defer func() {
		// send data to channel to end blocking in main function
		done <- true
	}()

	time.Sleep(d)
}

func main() {
	// create channel for two different ask
	task1Chan := make(chan bool)
	task2Chan := make(chan bool)

	// submit two goroutines
	go waitForMe(2*time.Second, task1Chan)
	go waitForMe(1*time.Second, task2Chan)

	// wait for multiple tasks
	// execute whichever sends data early
	select {
	case <-task1Chan:
		fmt.Println("Task 1 completed")
	case <-task2Chan:
		fmt.Println("Task 2 completed")
	}
}
