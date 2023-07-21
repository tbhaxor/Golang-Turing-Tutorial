package main

import "fmt"

type Shape2D interface {
	Area() float32
	Perimeter() float32
}

// print2DShapeDetails is a function that takes a variadic parameter 'shapes' of type Shape2D.
// It iterates over each shape in the 'shapes' slice and prints its area and perimeter.
func print2DShapeDetails(shapes ...Shape2D) {
	// The function uses a variadic parameter, allowing it to accept any number of Shape2D instances.
	for _, shape := range shapes {
		// The function iterates over each shape in the 'shapes' slice using a range loop.

		// Using the 'Area()' and 'Perimeter()' methods of the Shape2D interface to get the area and perimeter of the shape.
		// The Shape2D interface guarantees that all implementing types have these two methods.

		// 'shape.Area()' calculates and returns the area of the current shape.
		// 'shape.Perimeter()' calculates and returns the perimeter of the current shape.

		// The function then uses 'fmt.Printf' to print the area and perimeter with three decimal places.
		fmt.Printf("Area=%.3f\tPerimeter=%.3f\n", shape.Area(), shape.Perimeter())
	}
}

func getTypeDetails(shapes ...Shape2D) {
	for _, shape := range shapes {
		// 'shape.(type)' syntax is used within the type switch
		// to get the actual type of 'shape'.
		switch shape := shape.(type) {
		case *Circle:
			fmt.Printf("[Circle]\tr=%.2f\n", shape.radius)
		case *Triangle:
			fmt.Printf("[Triangle]\ta=%.2f\tb=%.2f\tc=%.2f\n", shape.a, shape.b, shape.c)
		case *Rectangle:
			fmt.Printf("[Rectangle]\tl=%.2f\tw=%.2f\n", shape.length, shape.width)
		case *Square:
			fmt.Printf("[Square]\ts=%.2f\n", shape.size)
		}
	}
}

func main() {
	// instantiate all the structs in array
	shapes := []Shape2D{
		&Square{size: 4},
		&Rectangle{length: 3, width: 4},
		&Circle{radius: 10},
		&Triangle{a: 3, b: 4, c: 5},
	}

	print2DShapeDetails(shapes...)

	for idx, shape := range shapes {
		// check if the underlying type of 'shape' is a Circle.
		// cast the interface to Circle type and set ok to true if type assertion was successful.
		circle, ok := shape.(*Circle)

		if ok {
			fmt.Printf("Found circle shape at %d index\n", idx)
			fmt.Println("Radius of circle:", circle.radius)
			print2DShapeDetails(circle)

			// shape is found, now exit the loop
			break
		}
	}

	fmt.Println()
	getTypeDetails(shapes...)

	// you will get compiler error if interface is not implemented
	// cannot use &Cube{…} (variable of type *Cube) as Shape2D value in argument to
	// print2DShapeDetails: *Cube does not implement Shape2D (missing method Area)
	// print2DShapeDetails(&Cube{side: 10}) // uncomment this go get the error
}
