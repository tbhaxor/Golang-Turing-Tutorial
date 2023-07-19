package arithmetic

import "errors"

func Div(a, b int) (float32, error) {
	if b == 0 {
		// divide by zero is undefined
		// if user provided this value, then send 0 and error
		return 0, errors.New("divide by zero error")
	}

	// since division can be fractional, convert operands to float32()
	// if b != 0 then we can safely divide and return value
	// error interface is allowed store and return nil values
	// only pointers and error can be assigned to nil
	return float32(a) / float32(b), nil
}
