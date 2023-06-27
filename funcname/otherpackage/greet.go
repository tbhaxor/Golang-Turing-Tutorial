package otherpackage

import "fmt"

// greetname will print hello to name
func greetname(name string) {
	fmt.Printf("Hello, %s\n", name)
}

// greetguest will print a simple string
func greetguest() {
	fmt.Println("Hello, Guest")
}

// Greet will print greetguest or greetname based on length of names
func Greet(names []string) {
	if len(names) == 0 {
		greetguest()
	} else {
		for _, name := range names {
			greetname(name)
		}
	}
}
