package main

import (
	"fmt"
)

func hashMaps() {
	fmt.Println("--- Demonstrating Hashmaps ---")

	// both key-value pair are string type
	// you can also use make, make(map[string]int)
	var myMap map[string]int = map[string]int{}

	// get value from the map
	value, exists := myMap["Alice"]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found") // this will be printed
	}

	// set or update value for the key Alice
	myMap["Alice"] = 20

	// print entier map
	fmt.Println(myMap)

	// when you are certain the key exists
	// feel free to ditch if-else
	value = myMap["Alice"]
	fmt.Println(value)

	// gives the number of pairs
	fmt.Println(len(myMap))

	// range based for loop gives key on left and value on right side
	for key, value := range myMap {
		fmt.Printf("Key=%s\tValue=%d\n", key, value)
	}

	// delete the if exists
	delete(myMap, "Alice")
	fmt.Println(len(myMap))

}
