package arithmetic

// Do method takes two operands and perform all 4 arithmetic operators
// you can also have types associated separately in parameters like this
func Do(a int, b int) (int, int, int, float32, error) {
	// since div function return two values
	// it needs to be deconstructed before return
	div, err := Div(a, b)

	// Add, Sub, Prod can be directly used as
	return Add(a, b), Sub(a, b), Prod(a, b), div, err
}
