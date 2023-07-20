package main

type Square struct {
	size float32
}

func (s *Square) Area() float32 {
	return s.size * s.size
}

func (s *Square) Perimeter() float32 {
	return 4 * s.size
}
