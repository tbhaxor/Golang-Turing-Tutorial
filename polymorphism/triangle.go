package main

import "math"

type Triangle struct {
	a, b, c float32
}

// Area uses heron's formula
func (t *Triangle) Area() float32 {
	// calculate semiperimeter
	s := t.Perimeter() / 2

	return float32(math.Sqrt(float64(s * (s - t.a) * (s - t.b) * (s - t.c))))
}

func (t *Triangle) Perimeter() float32 {
	return t.a + t.b + t.c
}
