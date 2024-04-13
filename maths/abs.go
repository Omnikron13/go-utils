package maths

// Abs returns the absolute value of x.
func Abs[T RationalNumber](x T) T {
	if x >= 0 {
		return x
	}
	return -x
}

