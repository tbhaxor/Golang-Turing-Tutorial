package main

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.length + r.width)
}
