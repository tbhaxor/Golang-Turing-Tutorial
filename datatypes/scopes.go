package main

import "fmt"

// you cannot use short hand := initialization on global scope
var globalVariable = 10

func printVariables() {
	fmt.Println("--- Demonstrating Variable Scopes ---")
	localVariable := 5

	// Print global and local variables within main function
	fmt.Println("Global:", globalVariable)
	fmt.Println("Local (main):", localVariable)

	{
		// Declare local variables within compound block
		localVariable := 7
		anotherLocalVariable := 15
		globalVariable := 20

		// Print local and global variables within compound block
		fmt.Println("Local (block):", localVariable)
		fmt.Println("Another local (block):", anotherLocalVariable)
		fmt.Println("Global (block):", globalVariable)
	}

	// Print local and global variables after the compound block
	fmt.Println("Local (main, after block):", localVariable)
	fmt.Println("Global (main, after block):", globalVariable)
}
