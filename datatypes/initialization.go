package main

import "fmt"

func variableInitialization() {
	fmt.Println("--- Demonstrating Variable Initialization ---")

	// Golang forbids unused variables
	// uncomment following line to give compile-time error
	// var unused_var string

	// Standard declaration: var variableName dataType
	var age int = 20 // reassigning values to the variable

	// Type inference: variableName := value
	name := "John"

	// Multiple variable declaration: var ( variable1 dataType1 variable2 dataType2 )
	var height, weight float64

	// Short declaration for multiple variables: variable1, variable2 := value1, value2
	country, population := "INDIA", 142.03

	// Printing the values of the variables
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Height:", height)
	fmt.Println("Weight:", weight)
	fmt.Println("Country:", country)
	fmt.Println("Population:", population)

}
