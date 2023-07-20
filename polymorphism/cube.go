package main

import "math"

type Cube struct {
	side float32
}

func (c *Cube) Volumne() float32 {
	return float32(math.Pow(float64(c.side), 3))
}
