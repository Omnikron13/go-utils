package maths

// Abs returns the absolute value of x.
func Abs[T RationalNumber](x T) T {
	if x >= 0 {
		return x
	}
   // TODO: benchmatk this against the bit-twiddling that the float64 only version uses.
	return -x
}

