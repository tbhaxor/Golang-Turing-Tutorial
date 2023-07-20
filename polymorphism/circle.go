package main

import "math"

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}
